pragma solidity 0.8.15;

import "../interfaces/automation/AutomationCompatibleInterface.sol";
import "../dev/interfaces/automation/MercuryLookupCompatibleInterface.sol";
import {ArbSys} from "../dev/vendor/@arbitrum/nitro-contracts/src/precompiles/ArbSys.sol";

contract MercuryUpkeep is AutomationCompatibleInterface, MercuryLookupCompatibleInterface {
  event MercuryPerformEvent(address indexed origin, address indexed sender, uint256 blockNumber, bytes data);

  ArbSys internal constant ARB_SYS = ArbSys(0x0000000000000000000000000000000000000064);

  uint256 public testRange;
  uint256 public interval;
  uint256 public lastBlock;
  uint256 public previousPerformBlock;
  uint256 public initialBlock;
  uint256 public counter;
  string[] public feeds;
  string public feedLabel;
  string public queryLabel;

  constructor(uint256 _testRange, uint256 _interval) {
    testRange = _testRange;
    interval = _interval;
    previousPerformBlock = 0;
    lastBlock = ARB_SYS.arbBlockNumber();
    initialBlock = 0;
    counter = 0;
    feedLabel = "feedIDStr";
    feeds = ["ETH-USD-ARBITRUM-TESTNET", "BTC-USD-ARBITRUM-TESTNET"];
    queryLabel = "blockNumber";
  }

  function mercuryCallback(bytes[] memory values, bytes memory extraData) external pure returns (bool, bytes memory) {
    bytes memory performData = new bytes(0);
    for (uint256 i = 0; i < values.length; i++) {
      performData = bytes.concat(performData, values[i]);
    }
    performData = bytes.concat(performData, extraData);
    return (true, performData);
  }

  function checkUpkeep(bytes calldata data) external view returns (bool, bytes memory) {
    if (!eligible()) {
      return (false, data);
    }
    revert MercuryLookup(feedLabel, feeds, queryLabel, ARB_SYS.arbBlockNumber(), data);
  }

  function performUpkeep(bytes calldata performData) external {
    uint256 blockNumber = ARB_SYS.arbBlockNumber();
    if (initialBlock == 0) {
      initialBlock = blockNumber;
    }
    lastBlock = blockNumber;
    counter = counter + 1;
    emit MercuryPerformEvent(tx.origin, msg.sender, blockNumber, performData);
    previousPerformBlock = lastBlock;
  }

  function eligible() public view returns (bool) {
    if (initialBlock == 0) {
      return true;
    }

    uint256 blockNumber = ARB_SYS.arbBlockNumber();
    return (blockNumber - initialBlock) < testRange && (blockNumber - lastBlock) >= interval;
  }

  function setConfig(uint256 _testRange, uint256 _interval) external {
    testRange = _testRange;
    interval = _interval;
    initialBlock = 0;
    counter = 0;
  }

  function setFeeds(string[] memory input) external {
    feeds = input;
  }
}
