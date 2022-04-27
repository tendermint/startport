//go:build !relayer
// +build !relayer

package app_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ignite-hq/cli/ignite/pkg/cmdrunner/step"
	envtest "github.com/ignite-hq/cli/integration"
)

func TestGenerateAnApp(t *testing.T) {
	var (
		env  = envtest.New(t)
		path = env.Scaffold("github.com/test/blog")
	)

	_, statErr := os.Stat(filepath.Join(path, "x", "blog"))
	require.False(t, os.IsNotExist(statErr), "the default module should be scaffolded")

	env.EnsureAppIsSteady(path)
}

// TestGenerateAnAppWithName tests scaffolding a new chain using a local name instead of a GitHub URI.
func TestGenerateAnAppWithName(t *testing.T) {
	var (
		env     = envtest.New(t)
		appName = "blog"
		root    = env.TmpDir()
	)

	// TODO: Change Env.Scaffold to avoid prefixing app names with GitHub URLs to avoid explicit scaffolding ?
	env.Exec("scaffold an app",
		step.NewSteps(step.New(
			step.Exec(
				envtest.IgniteApp,
				"scaffold",
				"chain",
				appName,
			),
			step.Workdir(root),
		)),
	)

	// Remove the files that were generated during the test when the integration test ends
	env.SetCleanup(func() {
		os.RemoveAll(filepath.Join(env.Home(), fmt.Sprintf(".%s", appName)))
	})

	path := filepath.Join(root, appName)

	_, statErr := os.Stat(filepath.Join(path, "x", "blog"))
	require.False(t, os.IsNotExist(statErr), "the default module should be scaffolded")

	env.EnsureAppIsSteady(path)
}

func TestGenerateAnAppWithNoDefaultModule(t *testing.T) {
	var (
		env     = envtest.New(t)
		appName = "blog"
	)

	root := env.TmpDir()
	env.Exec("scaffold an app",
		step.NewSteps(step.New(
			step.Exec(
				envtest.IgniteApp,
				"scaffold",
				"chain",
				fmt.Sprintf("github.com/test/%s", appName),
				"--no-module",
			),
			step.Workdir(root),
		)),
	)

	// Cleanup the home directory of the app
	env.SetCleanup(func() {
		os.RemoveAll(filepath.Join(env.Home(), fmt.Sprintf(".%s", appName)))
	})

	path := filepath.Join(root, appName)

	_, statErr := os.Stat(filepath.Join(path, "x", "blog"))
	require.True(t, os.IsNotExist(statErr), "the default module should not be scaffolded")

	env.EnsureAppIsSteady(path)
}

func TestGenerateAnAppWithNoDefaultModuleAndCreateAModule(t *testing.T) {
	var (
		env  = envtest.New(t)
		path = env.Scaffold("github.com/test/blog", "--no-module")
	)

	defer env.EnsureAppIsSteady(path)

	env.Must(env.Exec("should scaffold a new module into a chain that never had modules before",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "first_module"),
			step.Workdir(path),
		)),
	))
}

func TestGenerateAnAppWithWasm(t *testing.T) {
	t.Skip()

	var (
		env  = envtest.New(t)
		path = env.Scaffold("github.com/test/blog")
	)

	env.Must(env.Exec("add Wasm module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "wasm"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should not add Wasm module second time",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "wasm"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.EnsureAppIsSteady(path)
}

func TestGenerateAStargateAppWithEmptyModule(t *testing.T) {
	var (
		env  = envtest.New(t)
		path = env.Scaffold("github.com/test/blog")
	)

	env.Must(env.Exec("create a module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "example", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating an existing module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "example", "--require-registration"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a module with an invalid name",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "example1", "--require-registration"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a module with a reserved name",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "tx", "--require-registration"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a module with a forbidden prefix",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "ibcfoo", "--require-registration"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a module prefixed with an existing module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "s", "module", "examplefoo", "--require-registration"),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("create a module with dependencies",
		step.NewSteps(step.New(
			step.Exec(
				envtest.IgniteApp,
				"s",
				"module",
				"with_dep",
				"--dep",
				"account,bank,staking,slashing,example",
				"--require-registration",
			),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a module with invalid dependencies",
		step.NewSteps(step.New(
			step.Exec(
				envtest.IgniteApp,
				"s",
				"module",
				"with_wrong_dep",
				"--dep",
				"dup,dup",
				"--require-registration",
			),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.Must(env.Exec("should prevent creating a module with a non registered dependency",
		step.NewSteps(step.New(
			step.Exec(
				envtest.IgniteApp,
				"s",
				"module",
				"with_no_dep",
				"--dep",
				"inexistent",
				"--require-registration",
			),
			step.Workdir(path),
		)),
		envtest.ExecShouldError(),
	))

	env.EnsureAppIsSteady(path)
}
