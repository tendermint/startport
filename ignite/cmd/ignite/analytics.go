package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/manifoldco/promptui"

	"github.com/ignite/cli/ignite/pkg/gacli"
	"github.com/ignite/cli/ignite/pkg/randstr"
	"github.com/ignite/cli/ignite/version"
)

const (
	telemetryEndpoint  = "https://telemetry-cli.ignite.com"
	envDoNotTrack      = "DO_NOT_TRACK"
	igniteDir          = ".ignite"
	igniteAnonIdentity = "anonIdentity.json"
)

var gaclient gacli.Client

type (
	// metric represents an analytics metric.
	metric struct {
		// err sets metrics type as an error metric.
		err error
		// command is the command name.
		command string
	}

	// anonIdentity represents an analytics identity file.
	anonIdentity struct {
		// name represents the username.
		Name string `json:"name" yaml:"name"`
		// doNotTrack represents the user track choice.
		DoNotTrack bool `json:"doNotTrack" yaml:"doNotTrack"`
	}
)

func sendMetric(wg *sync.WaitGroup, m metric) {
	envDoNotTrackVar := os.Getenv(envDoNotTrack)
	if envDoNotTrackVar == "1" || strings.ToLower(envDoNotTrackVar) == "true" {
		return
	}

	if m.command == "ignite version" {
		return
	}

	dntInfo, err := checkDNT()
	if err != nil || dntInfo.DoNotTrack {
		return
	}

	met := gacli.Metric{
		OS:        runtime.GOOS,
		Arch:      runtime.GOARCH,
		FullCmd:   m.command,
		SessionID: dntInfo.Name,
		Version:   version.Version,
	}

	switch {
	case m.err == nil:
		met.Status = "success"
	case m.err != nil:
		met.Status = "error"
		met.Error = m.err.Error()
	}

	cmds := strings.Split(m.command, " ")
	met.Cmd = cmds[0]
	if len(cmds) > 0 {
		met.Cmd = cmds[1]
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		_ = gaclient.SendMetric(met)
	}()
}

func checkDNT() (anonIdentity, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return anonIdentity{}, err
	}
	if err := os.Mkdir(filepath.Join(home, igniteDir), 0o700); err != nil && !os.IsExist(err) {
		return anonIdentity{}, err
	}
	identityPath := filepath.Join(home, igniteDir, igniteAnonIdentity)
	data, err := os.ReadFile(identityPath)
	if err != nil && !os.IsNotExist(err) {
		return anonIdentity{}, err
	}

	var i anonIdentity
	if err := json.Unmarshal(data, &i); err == nil {
		return i, nil
	}

	i.Name = randstr.Runes(10)
	i.DoNotTrack = false

	prompt := promptui.Select{
		Label: "Ignite collects metrics about command usage. " +
			"All data is anonymous and helps to improve Ignite. " +
			"Ignite respect the DNT rules (consoledonottrack.com). " +
			"Would you agree to share these metrics with us?",
		Items: []string{"Yes", "No"},
	}
	resultID, _, err := prompt.Run()
	if err != nil {
		return anonIdentity{}, err
	}

	if resultID != 0 {
		i.DoNotTrack = true
	}

	data, err = json.Marshal(&i)
	if err != nil {
		return i, err
	}

	return i, os.WriteFile(identityPath, data, 0o700)
}
