package ignitecmd

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"

	"github.com/ignite/cli/ignite/pkg/cliui"
	"github.com/ignite/cli/ignite/pkg/cliui/model"
	"github.com/ignite/cli/ignite/pkg/cliui/style"
	"github.com/ignite/cli/ignite/pkg/events"
	"github.com/ignite/cli/ignite/services/chain"
)

const maxStatusEvents = 7

func initialChainServeModel(cmd *cobra.Command, session *cliui.Session) chainServeModel {
	return chainServeModel{
		starting: true,
		events:   model.NewEvents(session.EventBus(), maxStatusEvents),
		cmd:      cmd,
		session:  session,
	}
}

type chainServeModel struct {
	starting bool
	error    error
	events   model.Events
	cmd      *cobra.Command
	// TODO: Make session a value instead of a reference
	session *cliui.Session
}

func (m chainServeModel) Init() tea.Cmd {
	return tea.Batch(m.events.WaitEvent, chainServeStartCmd(m.cmd, m.session))
}

func (m chainServeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.Type == tea.KeyCtrlC {
			cmd = tea.Quit
		}
	case model.ErrorMsg:
		m.error = msg
		cmd = tea.Quit
	case model.EventMsg:
		if msg.ProgressIndication == events.IndicationFinish {
			// TODO: Listen to events in another model for the second view
			m.starting = false
		} else {
			m.events, cmd = m.events.Update(msg)
		}
	default:
		// This is required to allow event spinner updates
		m.events, cmd = m.events.Update(msg)
	}

	return m, cmd
}

func (m chainServeModel) View() string {
	if m.error != nil {
		return m.error.Error()
	}

	if m.starting {
		return m.renderStartView()
	}

	return ""
}

func (m chainServeModel) renderStartView() string {
	var view strings.Builder

	view.WriteString(m.events.View())
	fmt.Fprintf(&view, "%s\n", style.Faint.Render("Press the 'q' key to stop serve"))

	return view.String()
}

func chainServeStartCmd(cmd *cobra.Command, session *cliui.Session) tea.Cmd {
	return func() tea.Msg {
		chainOption := []chain.Option{
			chain.WithOutputer(session),
			chain.CollectEvents(session.EventBus()),
		}

		if flagGetProto3rdParty(cmd) {
			chainOption = append(chainOption, chain.EnableThirdPartyModuleCodegen())
		}

		if flagGetCheckDependencies(cmd) {
			chainOption = append(chainOption, chain.CheckDependencies())
		}

		// check if custom config is defined
		config, err := cmd.Flags().GetString(flagConfig)
		if err != nil {
			return err
		}
		if config != "" {
			chainOption = append(chainOption, chain.ConfigFile(config))
		}

		// create the chain
		c, err := newChainWithHomeFlags(cmd, chainOption...)
		if err != nil {
			return err
		}

		cacheStorage, err := newCache(cmd)
		if err != nil {
			return err
		}

		// serve the chain
		var serveOptions []chain.ServeOption

		forceUpdate, err := cmd.Flags().GetBool(flagForceReset)
		if err != nil {
			return err
		}

		if forceUpdate {
			serveOptions = append(serveOptions, chain.ServeForceReset())
		}

		resetOnce, err := cmd.Flags().GetBool(flagResetOnce)
		if err != nil {
			return err
		}

		if resetOnce {
			serveOptions = append(serveOptions, chain.ServeResetOnce())
		}

		quitOnFail, err := cmd.Flags().GetBool(flagQuitOnFail)
		if err != nil {
			return err
		}

		if quitOnFail {
			serveOptions = append(serveOptions, chain.QuitOnFail())
		}

		if flagGetSkipProto(cmd) {
			serveOptions = append(serveOptions, chain.ServeSkipProto())
		}

		return c.Serve(cmd.Context(), cacheStorage, serveOptions...)
	}
}
