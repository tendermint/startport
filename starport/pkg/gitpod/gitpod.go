package gitpod

import (
	"bytes"
	"context"
	"os"
	"strings"

	"github.com/tendermint/starport/starport/pkg/cmdrunner"
	"github.com/tendermint/starport/starport/pkg/cmdrunner/step"
)

// IsOnGitpod reports whether if running on Gitpod or not.
func IsOnGitpod() bool {
	return os.Getenv("GITPOD_WORKSPACE_ID") != ""
}

func URLForPort(ctx context.Context, port string) (string, error) {
	buf := bytes.Buffer{}
	if err := cmdrunner.New(cmdrunner.DefaultStdout(&buf)).Run(ctx, step.New(step.Exec("gp", "url", port))); err != nil {
		return "", err
	}

	return strings.TrimSpace(buf.String()), nil
}
