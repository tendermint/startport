// +build !relayer
package integration_test

import (
	"testing"

	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
)

func TestCreateMapWithStargate(t *testing.T) {
	var (
		env  = newEnv(t)
		path = env.Scaffold("blog")
	)

	env.Must(env.Exec("create a map",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "map", "user", "email"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a map with no message",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "map", "nomessage", "email", "--no-message"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a module",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "module", "example", "--require-registration"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("create a list",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "list", "user", "email", "--module", "example"),
			step.Workdir(path),
		)),
	))

	env.Must(env.Exec("should prevent creating a map with a typename that already exist",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "map", "user", "email", "--module", "example"),
			step.Workdir(path),
		)),
		ExecShouldError(),
	))

	env.Must(env.Exec("create a map in a custom module",
		step.NewSteps(step.New(
			step.Exec("starport", "s", "map", "mapuser", "email", "--module", "example"),
			step.Workdir(path),
		)),
	))

	env.EnsureAppIsSteady(path)
}
