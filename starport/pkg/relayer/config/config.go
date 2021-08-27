package relayerconf

import (
	"fmt"
	"os"
	"reflect"

	"github.com/tendermint/starport/starport/pkg/confile"
)

const supportVersion = "2"

var configPath = os.ExpandEnv("$HOME/.starport/relayer/config.yml")

type Config struct {
	Version string  `yaml:"version"`
	Chains  []Chain `yaml:"chains"`
	Paths   []Path  `yaml:"paths"`
}

func (c Config) PathByID(id string) (Path, error) {
	for _, path := range c.Paths {
		if path.ID == id {
			return path, nil
		}
	}
	return Path{}, fmt.Errorf("path %q cannot be found", id)
}

type Chain struct {
	ID            string `yaml:"id"`
	Account       string `yaml:"account"`
	AddressPrefix string `yaml:"address_prefix"`
	RPCAddress    string `yaml:"rpc_address"`
	GasPrice      string `yaml:"gas_price"`
	GasLimit      int64  `yaml:"gas_limit"`
}

type Path struct {
	ID       string  `yaml:"id"`
	Ordering string  `yaml:"ordering"`
	Src      PathEnd `yaml:"src"`
	Dst      PathEnd `yaml:"dst"`
}

type PathEnd struct {
	ChainID      string `yaml:"chain_id"`
	ConnectionID string `yaml:"connection_id"`
	ChannelID    string `yaml:"channel_id"`
	PortID       string `yaml:"port_id"`
	Version      string `yaml:"version"`
	PacketHeight int64  `yaml:"packet_height"`
	AckHeight    int64  `yaml:"ack_height"`
}

func Get() (Config, error) {
	c := Config{}
	if err := confile.New(confile.DefaultYAMLEncodingCreator, configPath).Load(&c); err != nil {
		return c, err
	}
	if !reflect.DeepEqual(c, Config{}) && c.Version != supportVersion {
		return c, fmt.Errorf("your relayer setup is outdated. run 'rm %s' and configure relayer again", configPath)
	}
	return c, nil
}

func Save(c Config) error {
	c.Version = supportVersion
	return confile.New(confile.DefaultYAMLEncodingCreator, configPath).Save(c)
}
