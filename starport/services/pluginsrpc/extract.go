package pluginsrpc

import (
	"context"
	"os/exec"
	"path"

	"github.com/hashicorp/go-plugin"
	"github.com/lukerhoads/plugintypes"
	"github.com/spf13/cobra"
	chaincfg "github.com/tendermint/starport/starport/chainconfig"
)

func (m *Manager) extractPlugins(ctx context.Context, rootCmd *cobra.Command) error {
	outputDir, err := formatPluginHome(m.ChainId, "output")
	if err != nil {
		return err
	}

	for i := 0; i < len(m.Config.Plugins); i++ {
		cmdPlugins, err := extractCommandPlugins(ctx, outputDir, rootCmd, m.Config)
		if err != nil {
			return err
		}

		hookPlugins, err := extractHookPlugins(ctx, outputDir, rootCmd, m.Config)
		if err != nil {
			return err
		}

		m.cmdPlugins = append(m.cmdPlugins, cmdPlugins...)
		m.hookPlugins = append(m.hookPlugins, hookPlugins...)
	}

	return nil
}

func extractCommandPlugins(
	ctx context.Context,
	outputDir string,
	parentCommand *cobra.Command,
	cfg chaincfg.Config,
) ([]ExtractedCommandModule, error) {
	pluginFiles, err := listDirsMatch(outputDir, `*_cmd`)
	if err != nil {
		return nil, err
	}

	if len(pluginFiles) == 0 {
		return []ExtractedCommandModule{}, nil
	}

	// Remove pluginFiles that are not specified in the config

	var extractedCommandModules []ExtractedCommandModule
	for _, pluginFile := range pluginFiles {
		pluginDir := path.Join(outputDir, pluginFile.Name())
		PluginMap := BasePluginMap

		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: HandshakeConfig,
			Plugins:         PluginMap,
			Cmd:             exec.Command(pluginDir),
		})

		rpcClient, err := client.Client()
		if err != nil {
			return []ExtractedCommandModule{}, err
		}

		raw, err := rpcClient.Dispense("command_map")
		if err != nil {
			return []ExtractedCommandModule{}, err
		}

		cmdMapper := raw.(plugintypes.CommandMapper)

		// Edit pluginMap off of that, and then make execute functions for each thing
		for _, loadedModule := range cmdMapper.Commands() {
			NewPluginMap := map[string]plugin.Plugin{
				loadedModule: &plugintypes.CommandModulePlugin{},
			}

			client2 := plugin.NewClient(&plugin.ClientConfig{
				HandshakeConfig: HandshakeConfig,
				Plugins:         NewPluginMap,
				Cmd:             exec.Command(pluginDir),
			})

			rpcClient2, err := client2.Client()
			if err != nil {
				return []ExtractedCommandModule{}, err
			}

			raw2, err := rpcClient2.Dispense(loadedModule)
			if err != nil {
				return []ExtractedCommandModule{}, err
			}

			cmdModule := raw2.(plugintypes.CommandModule)
			extractedCommandModules = append(extractedCommandModules, ExtractedCommandModule{
				ParentCommand: cmdModule.GetParentCommand(),
				Name:          cmdModule.GetName(),
				Usage:         cmdModule.GetUsage(),
				ShortDesc:     cmdModule.GetShortDesc(),
				LongDesc:      cmdModule.GetLongDesc(),
				NumArgs:       cmdModule.GetNumArgs(),
				Exec: func(cmd *cobra.Command, args []string) error {
					client := plugin.NewClient(&plugin.ClientConfig{
						HandshakeConfig: HandshakeConfig,
						Plugins:         NewPluginMap,
						Cmd:             exec.Command(pluginDir),
					})

					rpcClient, err := client.Client()
					if err != nil {
						return err
					}

					raw, err := rpcClient.Dispense(loadedModule)
					if err != nil {
						return err
					}

					cmdModuleExec := raw.(plugintypes.CommandModule)
					return cmdModuleExec.Exec(cmd, args)
				},
			})

			client2.Kill()
		}

		client.Kill()
	}

	return extractedCommandModules, nil
}

func extractHookPlugins(
	ctx context.Context,
	outputDir string,
	parentCommand *cobra.Command,
	cfg chaincfg.Config,
) ([]ExtractedHookModule, error) {
	pluginFiles, err := listDirsMatch(outputDir, `*_hook`)
	if err != nil {
		return nil, err
	}

	if len(pluginFiles) == 0 {
		return []ExtractedHookModule{}, nil
	}

	// Remove pluginFiles that are not specified in the config

	var extractedHookModules []ExtractedHookModule
	for _, pluginFile := range pluginFiles {
		pluginDir := path.Join(outputDir, pluginFile.Name())
		PluginMap := BasePluginMap

		client := plugin.NewClient(&plugin.ClientConfig{
			HandshakeConfig: HandshakeConfig,
			Plugins:         PluginMap,
			Cmd:             exec.Command(pluginDir),
		})

		rpcClient, err := client.Client()
		if err != nil {
			return []ExtractedHookModule{}, err
		}

		raw, err := rpcClient.Dispense("hook_map")
		if err != nil {
			return []ExtractedHookModule{}, err
		}

		hookMapper := raw.(plugintypes.HookMapper)

		// Edit pluginMap off of that, and then make execute functions for each thing
		for _, loadedModule := range hookMapper.Hooks() {
			NewPluginMap := map[string]plugin.Plugin{
				loadedModule: &plugintypes.HookModulePlugin{},
			}

			client2 := plugin.NewClient(&plugin.ClientConfig{
				HandshakeConfig: HandshakeConfig,
				Plugins:         NewPluginMap,
				Cmd:             exec.Command(pluginDir),
			})

			rpcClient2, err := client2.Client()
			if err != nil {
				return []ExtractedHookModule{}, err
			}

			raw2, err := rpcClient2.Dispense(loadedModule)
			if err != nil {
				return []ExtractedHookModule{}, err
			}

			hookModule := raw2.(plugintypes.HookModule)
			extractedHookModules = append(extractedHookModules, ExtractedHookModule{
				ParentCommand: hookModule.GetParentCommand(),
				Name:          hookModule.GetName(),
				HookType:      hookModule.GetType(),
				PreRun: func(cmd *cobra.Command, args []string) error {
					client := plugin.NewClient(&plugin.ClientConfig{
						HandshakeConfig: HandshakeConfig,
						Plugins:         NewPluginMap,
						Cmd:             exec.Command(pluginDir),
					})

					rpcClient, err := client.Client()
					if err != nil {
						return err
					}

					raw, err := rpcClient.Dispense(loadedModule)
					if err != nil {
						return err
					}

					cmdModuleExec := raw.(plugintypes.HookModule)
					return cmdModuleExec.PreRun(cmd, args)
				},
				PostRun: func(cmd *cobra.Command, args []string) error {
					client := plugin.NewClient(&plugin.ClientConfig{
						HandshakeConfig: HandshakeConfig,
						Plugins:         NewPluginMap,
						Cmd:             exec.Command(pluginDir),
					})

					rpcClient, err := client.Client()
					if err != nil {
						return err
					}

					raw, err := rpcClient.Dispense(loadedModule)
					if err != nil {
						return err
					}

					cmdModuleExec := raw.(plugintypes.HookModule)
					return cmdModuleExec.PostRun(cmd, args)
				},
			})

			client2.Kill()
		}

		client.Kill()
	}

	return extractedHookModules, nil
}
