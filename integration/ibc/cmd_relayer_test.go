//go:build !relayer

package ibc_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v2"

	"github.com/ignite/cli/v28/ignite/config/chain"
	"github.com/ignite/cli/v28/ignite/config/chain/base"
	v1 "github.com/ignite/cli/v28/ignite/config/chain/v1"
	"github.com/ignite/cli/v28/ignite/pkg/availableport"
	"github.com/ignite/cli/v28/ignite/pkg/cmdrunner/step"
	"github.com/ignite/cli/v28/ignite/pkg/goanalysis"
	"github.com/ignite/cli/v28/ignite/pkg/goenv"
	"github.com/ignite/cli/v28/ignite/pkg/randstr"
	yamlmap "github.com/ignite/cli/v28/ignite/pkg/yaml"
	envtest "github.com/ignite/cli/v28/integration"
)

const (
	relayerMnemonic = "great immense still pill defense fetch pencil slow purchase symptom speed arm shoot fence have divorce cigar rapid hen vehicle pear evolve correct nerve"
)

var (
	bobName    = "bob"
	marsConfig = v1.Config{
		Config: base.Config{
			Version: 1,
			Build: base.Build{
				Proto: base.Proto{
					Path:            "proto",
					ThirdPartyPaths: []string{"third_party/proto", "proto_vendor"},
				},
			},
			Accounts: []base.Account{
				{
					Name:     "alice",
					Coins:    []string{"100000000000token", "10000000000000000000stake"},
					Mnemonic: "slide moment original seven milk crawl help text kick fluid boring awkward doll wonder sure fragile plate grid hard next casual expire okay body",
				},
				{
					Name:     "bob",
					Coins:    []string{"100000000000token", "10000000000000000000stake"},
					Mnemonic: "trap possible liquid elite embody host segment fantasy swim cable digital eager tiny broom burden diary earn hen grow engine pigeon fringe claim program",
				},
				{
					Name:     "relayer",
					Coins:    []string{"100000000000token", "1000000000000000000000stake"},
					Mnemonic: relayerMnemonic,
				},
			},
			Faucet: base.Faucet{
				Name:  &bobName,
				Coins: []string{"500token", "100000000stake"},
				Host:  ":4501",
			},
			Genesis: yamlmap.Map{"chain_id": "mars-1"},
		},
		Validators: []v1.Validator{
			{
				Name:   "alice",
				Bonded: "100000000stake",
				App: yamlmap.Map{
					"api":      yamlmap.Map{"address": ":1318"},
					"grpc":     yamlmap.Map{"address": ":9092"},
					"grpc-web": yamlmap.Map{"address": ":9093"},
				},
				Config: yamlmap.Map{
					"p2p": yamlmap.Map{"laddr": ":26658"},
					"rpc": yamlmap.Map{"laddr": ":26658", "pprof_laddr": ":6061"},
				},
				Home: "$HOME/.mars",
			},
		},
	}
	earthConfig = v1.Config{
		Config: base.Config{
			Version: 1,
			Build: base.Build{
				Proto: base.Proto{
					Path:            "proto",
					ThirdPartyPaths: []string{"third_party/proto", "proto_vendor"},
				},
			},
			Accounts: []base.Account{
				{
					Name:     "alice",
					Coins:    []string{"100000000000token", "10000000000000000000stake"},
					Mnemonic: "slide moment original seven milk crawl help text kick fluid boring awkward doll wonder sure fragile plate grid hard next casual expire okay body",
				},
				{
					Name:     "bob",
					Coins:    []string{"100000000000token", "10000000000000000000stake"},
					Mnemonic: "trap possible liquid elite embody host segment fantasy swim cable digital eager tiny broom burden diary earn hen grow engine pigeon fringe claim program",
				},
				{
					Name:     "relayer",
					Coins:    []string{"100000000000token", "1000000000000000000000stake"},
					Mnemonic: relayerMnemonic,
				},
			},
			Faucet: base.Faucet{
				Name:  &bobName,
				Coins: []string{"500token", "100000000stake"},
				Host:  ":4500",
			},
			Genesis: yamlmap.Map{"chain_id": "earth-1"},
		},
		Validators: []v1.Validator{
			{
				Name:   "alice",
				Bonded: "100000000stake",
				App: yamlmap.Map{
					"api":      yamlmap.Map{"address": ":1317"},
					"grpc":     yamlmap.Map{"address": ":9090"},
					"grpc-web": yamlmap.Map{"address": ":9091"},
				},
				Config: yamlmap.Map{
					"p2p": yamlmap.Map{"laddr": ":26656"},
					"rpc": yamlmap.Map{"laddr": ":26656", "pprof_laddr": ":6060"},
				},
				Home: "$HOME/.earth",
			},
		},
	}

	nameSendIbcPost = "SendIbcPost"
	funcSendIbcPost = `package keeper
func (k msgServer) SendIbcPost(goCtx context.Context, msg *types.MsgSendIbcPost) (*types.MsgSendIbcPostResponse, error) {
    ctx := sdk.UnwrapSDKContext(goCtx)
    // Construct the packet
    var packet types.IbcPostPacketData
    packet.Title = msg.Title
    packet.Content = msg.Content
    // Transmit the packet
    _, err := k.TransmitIbcPostPacket(
        ctx,
        packet,
        msg.Port,
        msg.ChannelID,
        clienttypes.ZeroHeight(),
        msg.TimeoutTimestamp,
    )
    return &types.MsgSendIbcPostResponse{}, err
}`

	nameOnRecvIbcPostPacket = "OnRecvIbcPostPacket"
	funcOnRecvIbcPostPacket = `package keeper
func (k Keeper) OnRecvIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData) (packetAck types.IbcPostPacketAck, err error) {
    // validate packet data upon receiving
    if err := data.ValidateBasic(); err != nil {
        return packetAck, err
    }
    packetAck.PostId = k.AppendPost(ctx, types.Post{Title: data.Title, Content: data.Content})
    return packetAck, nil
}`

	nameOnAcknowledgementIbcPostPacket = "OnAcknowledgementIbcPostPacket"
	funcOnAcknowledgementIbcPostPacket = `package keeper
func (k Keeper) OnAcknowledgementIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData, ack channeltypes.Acknowledgement) error {
    switch dispatchedAck := ack.Response.(type) {
    case *channeltypes.Acknowledgement_Error:
        // We will not treat acknowledgment error in this tutorial
        return nil
    case *channeltypes.Acknowledgement_Result:
        // Decode the packet acknowledgment
        var packetAck types.IbcPostPacketAck
        if err := types.ModuleCdc.UnmarshalJSON(dispatchedAck.Result, &packetAck); err != nil {
            // The counter-party module doesn't implement the correct acknowledgment format
            return errors.New("cannot unmarshal acknowledgment")
        }

        k.AppendSentPost(ctx,
            types.SentPost{
                PostId:  packetAck.PostId,
                Title:   data.Title,
                Chain:   packet.DestinationPort + "-" + packet.DestinationChannel,
            },
        )
        return nil
    default:
        return errors.New("the counter-party module does not implement the correct acknowledgment format")
    }
}`

	nameOnTimeoutIbcPostPacket = "OnTimeoutIbcPostPacket"
	funcOnTimeoutIbcPostPacket = `package keeper
func (k Keeper) OnTimeoutIbcPostPacket(ctx sdk.Context, packet channeltypes.Packet, data types.IbcPostPacketData) error {
    k.AppendTimeoutPost(ctx,
        types.TimeoutPost{
            Title:   data.Title,
            Chain:   packet.DestinationPort + "-" + packet.DestinationChannel,
        },
    )
    return nil
}`
)

func runChain(
	t *testing.T,
	env envtest.Env,
	app envtest.App,
	cfg v1.Config,
	ports []uint,
) (api, rpc, grpc, faucet string) {
	t.Helper()
	if len(ports) < 7 {
		t.Fatalf("invalid number of ports %d", len(ports))
	}
	var (
		ctx      = env.Ctx()
		tmpDir   = t.TempDir()
		homePath = filepath.Join(tmpDir, randstr.Runes(10))
		cfgPath  = filepath.Join(tmpDir, chain.ConfigFilenames[0])
	)
	genAddr := func(port uint) string {
		return fmt.Sprintf("0.0.0.0:%d", port)
	}

	cfg.Validators[0].Home = homePath

	cfg.Faucet.Host = genAddr(ports[0])
	cfg.Validators[0].App["api"] = yamlmap.Map{"address": genAddr(ports[1])}
	cfg.Validators[0].App["grpc"] = yamlmap.Map{"address": genAddr(ports[2])}
	cfg.Validators[0].App["grpc-web"] = yamlmap.Map{"address": genAddr(ports[3])}
	cfg.Validators[0].Config["p2p"] = yamlmap.Map{"laddr": genAddr(ports[4])}
	cfg.Validators[0].Config["rpc"] = yamlmap.Map{
		"laddr":       genAddr(ports[5]),
		"pprof_laddr": genAddr(ports[6]),
	}

	file, err := os.Create(cfgPath)
	require.NoError(t, err)
	require.NoError(t, yaml.NewEncoder(file).Encode(cfg))
	require.NoError(t, file.Close())

	ctx, cancel := context.WithCancel(ctx)
	t.Cleanup(func() {
		cancel()
		require.NoError(t, os.RemoveAll(tmpDir))
	})

	app.SetConfigPath(cfgPath)
	app.SetHomePath(homePath)
	go func() {
		env.Must(app.Serve("should serve chain", envtest.ExecCtx(ctx)))
	}()

	genHTTPAddr := func(port uint) string {
		return fmt.Sprintf("http://127.0.0.1:%d", port)
	}
	return genHTTPAddr(ports[1]), genHTTPAddr(ports[5]), genHTTPAddr(ports[2]), genHTTPAddr(ports[0])
}

func TestBlogIBC(t *testing.T) {
	var (
		env = envtest.New(t)
		app = env.Scaffold("github.com/test/planet")
		ctx = env.Ctx()
	)

	env.Must(env.Exec("create an IBC module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"s",
				"module",
				"blog",
				"--ibc",
				"--require-registration",
				"--yes",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a post type list in an IBC module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"s",
				"list",
				"post",
				"title",
				"content",
				"--no-message",
				"--module",
				"blog",
				"--yes",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a sentPost type list in an IBC module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"s",
				"list",
				"sentPost",
				"postID:uint",
				"title",
				"chain",
				"--no-message",
				"--module",
				"blog",
				"--yes",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a timeoutPost type list in an IBC module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"s",
				"list",
				"timeoutPost",
				"title",
				"chain",
				"--no-message",
				"--module",
				"blog",
				"--yes",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("create a ibcPost package in an IBC module",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"s",
				"packet",
				"ibcPost",
				"title",
				"content",
				"--ack",
				"postID:uint",
				"--module",
				"blog",
				"--yes",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	blogKeeperPath := filepath.Join(app.SourcePath(), "x/blog/keeper")
	require.NoError(t, goanalysis.ReplaceCode(
		blogKeeperPath,
		nameSendIbcPost,
		funcSendIbcPost,
	))
	require.NoError(t, goanalysis.ReplaceCode(
		blogKeeperPath,
		nameOnRecvIbcPostPacket,
		funcOnRecvIbcPostPacket,
	))
	require.NoError(t, goanalysis.ReplaceCode(
		blogKeeperPath,
		nameOnAcknowledgementIbcPostPacket,
		funcOnAcknowledgementIbcPostPacket,
	))
	require.NoError(t, goanalysis.ReplaceCode(
		blogKeeperPath,
		nameOnTimeoutIbcPostPacket,
		funcOnTimeoutIbcPostPacket,
	))

	// serve both chains.
	ports, err := availableport.Find(14)
	require.NoError(t, err)
	earthAPI, earthRPC, earthGRPC, earthFaucet := runChain(t, env, app, earthConfig, ports[:7])
	marsAPI, marsRPC, marsGRPC, marsFaucet := runChain(t, env, app, marsConfig, ports[7:])
	earthChainID := earthConfig.Genesis["chain_id"].(string)
	marsChainID := marsConfig.Genesis["chain_id"].(string)

	// check the chains is up
	stepsCheckChains := step.NewSteps(
		step.New(
			step.Exec(
				app.Binary(),
				"config",
				"output", "json",
			),
			step.PreExec(func() error {
				if err := env.IsAppServed(ctx, earthAPI); err != nil {
					return err
				}
				return env.IsAppServed(ctx, marsAPI)
			}),
		),
	)
	env.Exec("waiting the chain is up", stepsCheckChains, envtest.ExecRetry())

	// ibc relayer.
	env.Must(env.Exec("install the hermes relayer app",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"app",
				"install",
				"-g",
				filepath.Join(goenv.GoPath(), "src/github.com/ignite/apps/hermes"),
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	env.Must(env.Exec("configure the hermes relayer app",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp,
				"relayer",
				"hermes",
				"configure",
				earthChainID,
				earthRPC,
				earthGRPC,
				marsChainID,
				marsRPC,
				marsGRPC,
				"--chain-a-faucet", earthFaucet,
				"--chain-b-faucet", marsFaucet,
				"--generate-wallets",
				"--overwrite-config",
			),
			step.Workdir(app.SourcePath()),
		)),
	))

	// go func() {
	env.Must(env.Exec("run the hermes relayer",
		step.NewSteps(step.New(
			step.Exec(envtest.IgniteApp, "relayer", "hermes", "start", earthChainID, marsChainID),
			step.Workdir(app.SourcePath()),
			step.Stdout(os.Stdout),
			step.Stdin(os.Stdin),
			step.Stderr(os.Stderr),
		)),
	))
	//}()

	stepsCheckRelayer := step.NewSteps(
		step.New(
			step.Exec(
				// TODO query chain connection-id
				app.Binary(),
				"config",
				"output", "json",
			),
			step.PreExec(func() error {
				if err := env.IsAppServed(ctx, earthAPI); err != nil {
					return err
				}
				return env.IsAppServed(ctx, marsAPI)
			}),
		),
	)
	env.Exec("run the ts relayer", stepsCheckRelayer, envtest.ExecRetry())

	var (
		output     = &bytes.Buffer{}
		txResponse struct {
			Code   int
			RawLog string `json:"raw_log"`
		}
	)

	// sign tx to add an item to the list.
	stepsTx := step.NewSteps(
		step.New(
			step.Stdout(output),
			step.PreExec(func() error {
				err := env.IsAppServed(ctx, earthGRPC)
				return err
			}),
			step.Exec(
				app.Binary(),
				"tx",
				"blog",
				"send-ibc-post",
				"channel-0",
				"Hello",
				"Hello Mars, I'm Alice from Earth",
				"--chain-id", "blog",
				"--from", "alice",
				"--node", earthGRPC,
				"--output", "json",
				"--log_format", "json",
				"--yes",
			),
			step.PostExec(func(execErr error) error {
				if execErr != nil {
					return execErr
				}
				err := json.Unmarshal(output.Bytes(), &txResponse)
				if err != nil {
					return fmt.Errorf("unmarshling tx response: %w", err)
				}
				return nil
			}),
		),
	)
	if !env.Exec("sign a tx", stepsTx, envtest.ExecRetry()) {
		t.FailNow()
	}
	require.Equal(t, 0, txResponse.Code,
		"tx failed code=%d log=%s", txResponse.Code, txResponse.RawLog)
}