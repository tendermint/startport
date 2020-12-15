package chain

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	starporterrors "github.com/tendermint/starport/starport/errors"
	"github.com/tendermint/starport/starport/pkg/cmdrunner"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
	"github.com/tendermint/starport/starport/pkg/cosmosprotoc"
	starportconf "github.com/tendermint/starport/starport/services/chain/conf"
)

// Build builds an app.
func (c *Chain) Build(ctx context.Context) error {
	if err := c.setup(ctx); err != nil {
		return err
	}
	conf, err := c.Config()
	if err != nil {
		return &CannotBuildAppError{err}
	}

	steps, err := c.buildSteps(ctx, conf)
	if err != nil {
		return err
	}
	if err := cmdrunner.
		New(c.cmdOptions()...).
		Run(ctx, steps...); err != nil {
		return err
	}

	fmt.Fprintf(c.stdLog(logStarport).out, "🗃  Installed. Use with: %s\n", infoColor(strings.Join(c.plugin.Binaries(), ", ")))
	return nil
}

func (c *Chain) buildSteps(ctx context.Context, conf starportconf.Config) (
	steps step.Steps, err error) {
	chainID, err := c.ID()
	if err != nil {
		return nil, err
	}

	ldflags := fmt.Sprintf(`-X github.com/cosmos/cosmos-sdk/version.Name=NewApp
-X github.com/cosmos/cosmos-sdk/version.ServerName=%sd
-X github.com/cosmos/cosmos-sdk/version.ClientName=%scli
-X github.com/cosmos/cosmos-sdk/version.Version=%s
-X github.com/cosmos/cosmos-sdk/version.Commit=%s
-X %s/cmd/%s/cmd.ChainID=%s`,
		c.app.Name,
		c.app.Name,
		c.version.tag,
		c.version.hash,
		c.app.ImportPath,
		c.app.D(),
		chainID,
	)
	var (
		buildErr = &bytes.Buffer{}
	)
	captureBuildErr := func(err error) error {
		var exitErr *exec.ExitError
		if errors.As(err, &exitErr) {
			return &CannotBuildAppError{errors.New(buildErr.String())}
		}
		return err
	}

	steps.Add(step.New(step.NewOptions().
		Add(
			step.Exec(
				"go",
				"mod",
				"tidy",
			),
			step.PreExec(func() error {
				fmt.Fprintln(c.stdLog(logStarport).out, "📦 Installing dependencies...")
				return nil
			}),
			step.PostExec(captureBuildErr),
		).
		Add(c.stdSteps(logStarport)...).
		Add(step.Stderr(buildErr))...,
	))

	steps.Add(step.New(step.NewOptions().
		Add(
			step.Exec(
				"go",
				"mod",
				"verify",
			),
			step.PostExec(captureBuildErr),
		).
		Add(c.stdSteps(logBuild)...).
		Add(step.Stderr(buildErr))...,
	))

	// install the app.
	steps.Add(step.New(
		step.PreExec(func() error {
			fmt.Fprintln(c.stdLog(logStarport).out, "🛠️  Building the app...")
			return nil
		}),
	))

	for _, binary := range c.plugin.Binaries() {
		steps.Add(step.New(step.NewOptions().
			Add(
				// ldflags somehow won't work if directly execute go binary.
				// bash stays as a workaround for now.
				step.Exec(
					"bash", "-c", fmt.Sprintf("go install -mod readonly -ldflags '%s'", ldflags),
				),
				step.Workdir(filepath.Join(c.app.Path, "cmd", binary)),
				step.PostExec(captureBuildErr),
			).
			Add(c.stdSteps(logStarport)...).
			Add(step.Stderr(buildErr))...,
		))
	}
	return steps, nil
}

func (c *Chain) buildProto(ctx context.Context) error {
	// If protocgen exists, compile the proto file
	protoScriptPath := "scripts/protocgen"

	if _, err := os.Stat(protoScriptPath); os.IsNotExist(err) {
		return nil
	}

	if err := cosmosprotoc.InstallDependencies(context.Background(), c.app.Path); err != nil {
		if err == cosmosprotoc.ErrProtocNotInstalled {
			return starporterrors.ErrStarportRequiresProtoc
		}
		return err
	}

	var errb bytes.Buffer

	err := cmdrunner.
		New(c.cmdOptions()...).
		Run(ctx, step.New(
			step.Exec(
				"/bin/bash",
				protoScriptPath,
			),
			step.PreExec(func() error {
				fmt.Fprintln(c.stdLog(logStarport).out, "🛠️  Building proto...")
				return nil
			}),
			step.Stderr(&errb),
		))

	if err := errors.Wrap(err, errb.String()); err != nil {
		return &CannotBuildAppError{err}
	}

	return nil
}
