package starportcmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/tendermint/starport/starport/pkg/clispinner"
	"github.com/tendermint/starport/starport/services/network"
	"github.com/tendermint/starport/starport/services/network/networkchain"
)

// NewNetworkChainInstall returns a new command to install a chain's binary by the launch id.
func NewNetworkChainInstall() *cobra.Command {
	c := &cobra.Command{
		Use:   "install [launch-id]",
		Short: "Install chain binary for a launch",
		Args:  cobra.ExactArgs(1),
		RunE:  networkChainInstallHandler,
	}
	c.Flags().AddFlagSet(flagNetworkFrom())
	return c
}

func networkChainInstallHandler(cmd *cobra.Command, args []string) error {
	nb, err := newNetworkBuilder(cmd)
	if err != nil {
		return err
	}
	defer nb.Cleanup()

	// parse launch ID
	launchID, err := network.ParseID(args[0])
	if err != nil {
		return err
	}

	n, err := nb.Network()
	if err != nil {
		return err
	}

	chainLaunch, err := n.ChainLaunch(cmd.Context(), launchID)
	if err != nil {
		return err
	}

	c, err := nb.Chain(networkchain.SourceLaunch(chainLaunch))
	if err != nil {
		return err
	}

	binaryName, err := c.Build(cmd.Context())
	if err != nil {
		return err
	}

	fmt.Printf("%s Binary installed: %s\n", clispinner.OK, binaryName)

	return nil
}
