package testutil

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/plush"
	"github.com/tendermint/starport/starport/pkg/xgenny"
)

const (
	modulePathKey = "ModulePath"
	testUtilDir   = "testutil"
)

var (
	//go:embed stargate/* stargate/**/*
	fs embed.FS

	testutilTemplate = xgenny.NewEmbedWalker(fs, "stargate/")
)

// Register testutil template using existing generator.
// Register is meant to be used by modules that depend on this module.
//nolint:interfacer
func Register(ctx *plush.Context, gen *genny.Generator) error {
	if !ctx.Has(modulePathKey) {
		return fmt.Errorf("ctx is missing value for the ket %s", modulePathKey)
	}
	path, err := filepath.Abs(testUtilDir)
	if err != nil {
		return err
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return nil
	}
	return gen.Box(testutilTemplate)
}
