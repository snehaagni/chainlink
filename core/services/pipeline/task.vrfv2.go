package pipeline

import (
	"bytes"
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/smartcontractkit/chainlink/core/internal/gethwrappers/generated/vrf_coordinator_v2"
	"github.com/smartcontractkit/chainlink/core/services/eth"

	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/ethereum/go-ethereum/common"
	"github.com/pkg/errors"
	"github.com/smartcontractkit/chainlink/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/core/services/vrf/proof"
	"go.uber.org/multierr"
)

var (
	vrfCoordinatorV2ABI = eth.MustGetABI(vrf_coordinator_v2.VRFCoordinatorV2ABI)
)

type VRFTaskV2 struct {
	BaseTask           `mapstructure:",squash"`
	PublicKey          string `json:"publicKey"`
	RequestBlockHash   string `json:"requestBlockHash"`
	RequestBlockNumber string `json:"requestBlockNumber"`
	Topics             string `json:"topics"`

	keyStore VRFKeyStore
}

var _ Task = (*VRFTaskV2)(nil)

func (t *VRFTaskV2) Type() TaskType {
	return TaskTypeVRFV2
}

func (t *VRFTaskV2) Run(_ context.Context, vars Vars, inputs []Result) (result Result) {
	if len(inputs) != 1 {
		return Result{Error: ErrWrongInputCardinality}
	}
	if inputs[0].Error != nil {
		return Result{Error: ErrInputTaskErrored}
	}
	logValues, ok := inputs[0].Value.(map[string]interface{})
	if !ok {
		return Result{Error: errors.Wrap(ErrBadInput, "expected map input")}
	}
	var (
		pubKey             BytesParam
		requestBlockHash   BytesParam
		requestBlockNumber Uint64Param
		topics             HashSliceParam
	)
	err := multierr.Combine(
		errors.Wrap(ResolveParam(&pubKey, From(VarExpr(t.PublicKey, vars))), "publicKey"),
		errors.Wrap(ResolveParam(&requestBlockHash, From(VarExpr(t.RequestBlockHash, vars))), "requestBlockHash"),
		errors.Wrap(ResolveParam(&requestBlockNumber, From(VarExpr(t.RequestBlockNumber, vars))), "requestBlockNumber"),
		errors.Wrap(ResolveParam(&topics, From(VarExpr(t.Topics, vars))), "topics"),
	)
	if err != nil {
		return Result{Error: err}
	}

	requestKeyHash, ok := logValues["keyHash"].([32]byte)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid keyHash")}
	}
	requestPreSeed, ok := logValues["preSeed"].(*big.Int)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid preSeed")}
	}
	requestId, ok := logValues["requestId"].(*big.Int)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid requestId")}
	}
	subID, ok := logValues["subId"].(uint64)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid subId")}
	}
	callbackGasLimit, ok := logValues["callbackGasLimit"].(uint32)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid callbackGasLimit")}
	}
	numWords, ok := logValues["numWords"].(uint32)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid numWords")}
	}
	sender, ok := logValues["sender"].(common.Address)
	if !ok {
		return Result{Error: errors.Wrapf(ErrBadInput, "invalid sender")}
	}
	var pk secp256k1.PublicKey
	copy(pk[:], pubKey[:])
	pkh := pk.MustHash()
	// Validate the key against the spec
	if !bytes.Equal(requestKeyHash[:], pkh[:]) {
		return Result{Error: fmt.Errorf("invalid key hash %v expected %v", hex.EncodeToString(requestKeyHash[:]), hex.EncodeToString(pkh[:]))}
	}
	preSeed, err := proof.BigToSeed(requestPreSeed)
	if err != nil {
		return Result{Error: fmt.Errorf("unable to parse preseed %v", preSeed)}
	}
	preSeedData := proof.PreSeedDataV2{
		PreSeed:          preSeed,
		BlockHash:        common.BytesToHash(requestBlockHash),
		BlockNum:         uint64(requestBlockNumber),
		SubId:            subID,
		CallbackGasLimit: callbackGasLimit,
		NumWords:         numWords,
		Sender:           sender,
	}
	finalSeed := proof.FinalSeedV2(preSeedData)
	id := hexutil.Encode(pk[:])
	p, err := t.keyStore.GenerateProof(id, finalSeed)
	if err != nil {
		return Result{Error: err}
	}
	onChainProof, rc, err := proof.GenerateProofResponseFromProofV2(p, preSeedData)
	if err != nil {
		return Result{Error: err}
	}
	b, err := vrfCoordinatorV2ABI.Pack("fulfillRandomWords", onChainProof, rc)
	if err != nil {
		return Result{Error: err}
	}
	results := make(map[string]interface{})
	results["output"] = hexutil.Encode(b)
	// RequestID needs to be a [32]byte for EthTxMeta.
	results["requestID"] = hexutil.Encode(requestId.Bytes())
	return Result{Value: results}
}
