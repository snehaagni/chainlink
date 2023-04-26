package main

import (
	"os"
	"testing"

	"github.com/rogpeppe/go-internal/testscript"

	"github.com/smartcontractkit/chainlink/v2/core"
	"github.com/smartcontractkit/chainlink/v2/core/static"
)

func TestMain(m *testing.M) {
	os.Exit(testscript.RunMain(m, map[string]func() int{
		"chainlink": core.Main,
	}))
}

func TestScripts(t *testing.T) {
	testscript.Run(t, testscript.Params{
		Dir: "testdata/scripts",
		Setup: func(env *testscript.Env) error {
			env.Setenv("HOME", "$WORK/home")
			env.Setenv("VERSION", static.Version)
			env.Setenv("COMMIT_SHA", static.Sha)
			return nil
		},
	})
}
