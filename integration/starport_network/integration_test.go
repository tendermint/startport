// Package starport_network_test integration test Starport Network.
package starport_network_test

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"testing"

	"github.com/tendermint/starport/starport/pkg/xexec"
)

func TestMain(m *testing.M) {
	flag.Parse()

	// check requirements
	if err := checkSystemRequirements(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// TODO: Implement localspn package
	// setup SPN for Starport Network integration test
	//ctx, cancel := context.WithCancel(context.Background())
	//cleanup, err := localspn.SetupSPN(ctx, localspn.WithBranch("develop"))
	//if err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}

	// Run tests
	errCode := m.Run()

	// TODO: Implement localspn package
	//cancel()
	//cleanup()
	os.Exit(errCode)
}

func checkSystemRequirements() error {
	if !xexec.IsCommandAvailable("starport") {
		return errors.New("starport needs to be installed")
	}
	return nil
}
