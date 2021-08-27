package starportcmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	"github.com/tendermint/starport/starport/pkg/clispinner"
	"github.com/tendermint/starport/starport/pkg/cosmosaccount"
	"github.com/tendermint/starport/starport/pkg/relayer"
)

// NewRelayerConnect returns a new relayer connect command to link all or some relayer paths and start
// relaying txs in between.
// if not paths are specified, all paths are linked.
func NewRelayerConnect() *cobra.Command {
	c := &cobra.Command{
		Use:   "connect [<path>,...]",
		Short: "Link chains associated with paths and start relaying tx packets in between",
		RunE:  relayerConnectHandler,
	}

	c.Flags().AddFlagSet(flagSetKeyringBackend())

	return c
}

func relayerConnectHandler(cmd *cobra.Command, args []string) error {
	s := clispinner.New()
	defer s.Stop()

	var (
		givenPathIDs = args
		pathsToUse   []relayer.Path
		privKeys     = make(map[string]string)
	)

	allPaths, err := relayer.ListPaths(cmd.Context())
	if err != nil {
		return err
	}

	if len(givenPathIDs) > 0 {
		for _, id := range givenPathIDs {
			for _, path := range allPaths {
				if id == path.ID {
					pathsToUse = append(pathsToUse, path)
					break
				}

			}
		}
	} else {
		pathsToUse = allPaths
	}

	if len(pathsToUse) == 0 {
		s.Stop()

		fmt.Println("No chains found to connect.")
		return nil
	}

	s.SetText("Linking paths between chains...")

	ca, err := cosmosaccount.New(getKeyringBackend(cmd))
	if err != nil {
		return err
	}

	_ = func(name string) error {
		if _, ok := privKeys[name]; ok {
			return nil
		}
		key, err := ca.Export(name, "")
		if err != nil {
			return err
		}
		privKeys[name] = key
		return nil
	}

	//for _, path := range pathsToUse {
	//if err := ensureKeyAppended(path.Src.Account); err != nil {
	//return err
	//}
	//if err := ensureKeyAppended(path.Dst.Account); err != nil {
	//return err
	//}
	//}

	linkedPaths, alreadyLinkedPaths, failedToLinkPaths, err := relayer.Link(cmd.Context(), pathsToUse, privKeys)
	if err != nil {
		return err
	}

	s.Stop()

	fmt.Println()
	printSection("Linking chains")

	if len(alreadyLinkedPaths) != 0 {
		fmt.Printf("✓ %d paths already created to link chains.\n", len(alreadyLinkedPaths))
		for _, id := range alreadyLinkedPaths {
			fmt.Printf("  - %s\n", id)
		}
		fmt.Println()
	}

	if len(linkedPaths) != 0 {
		fmt.Printf("✓ Linked chains with %d paths.\n", len(linkedPaths))
		for _, id := range linkedPaths {
			fmt.Printf("  - %s\n", id)
		}
		fmt.Println()
	}

	pathsToConnect := append(linkedPaths, alreadyLinkedPaths...)

	if len(failedToLinkPaths) != 0 {
		fmt.Printf("x Failed to link chains in %d paths.\n", len(failedToLinkPaths))
		for _, failed := range failedToLinkPaths {
			fmt.Printf("  - %s failed with error: %s\n", failed.ID, failed.ErrorMsg)
		}
		fmt.Println()
	}

	if len(pathsToConnect) == 0 {
		fmt.Println("No paths to connect.")
		return nil
	}

	fmt.Printf("Continuing with %d paths...\n\n", len(pathsToConnect))

	printSection("Chains by paths")

	for _, id := range pathsToConnect {
		s.SetText("Loading...").Start()

		path, err := relayer.GetPath(cmd.Context(), id)
		if err != nil {
			return err
		}

		s.Stop()

		w := tabwriter.NewWriter(os.Stdout, 0, 0, 1, ' ', tabwriter.TabIndent)
		fmt.Fprintf(w, "%s:\n", path.ID)
		fmt.Fprintf(w, "   \t%s\t>\t(port: %s)\t(channel: %s)\n", path.Src.ChainID, path.Src.PortID, path.Src.ChannelID)
		fmt.Fprintf(w, "   \t%s\t>\t(port: %s)\t(channel: %s)\n", path.Dst.ChainID, path.Dst.PortID, path.Dst.ChannelID)
		fmt.Fprintln(w)
		w.Flush()
	}

	printSection("Listening and relaying packets between chains...")

	return relayer.Start(cmd.Context(), append(linkedPaths, alreadyLinkedPaths...)...)
}
