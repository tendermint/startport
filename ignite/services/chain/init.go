package chain

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/imdario/mergo"

	"github.com/ignite/cli/ignite/chainconfig"
	chaincmdrunner "github.com/ignite/cli/ignite/pkg/chaincmd/runner"
	"github.com/ignite/cli/ignite/pkg/cliui/view/accountview"
	"github.com/ignite/cli/ignite/pkg/confile"
)

const (
	moniker = "mynode"
)

// Init initializes the chain and applies all optional configurations.
func (c *Chain) Init(ctx context.Context, initAccounts bool) error {
	conf, err := c.Config()
	if err != nil {
		return &CannotBuildAppError{err}
	}

	if err := c.InitChain(ctx, conf, true); err != nil {
		return err
	}

	if initAccounts {
		return c.InitAccounts(ctx, conf)
	}
	return nil
}

// InitChain initializes the chain.
func (c *Chain) InitChain(ctx context.Context, conf *chainconfig.Config, updateNodeConfig bool) error {
	chainID, err := c.ID()
	if err != nil {
		return err
	}

	// cleanup persistent data from previous `serve`.
	home, err := c.Home()
	if err != nil {
		return err
	}
	if err := os.RemoveAll(home); err != nil {
		return err
	}

	commands, err := c.Commands(ctx)
	if err != nil {
		return err
	}

	// init node.
	if err := commands.Init(ctx, moniker); err != nil {
		return err
	}

	// overwrite app config files with the values defined in Ignite's config file
	if updateNodeConfig {
		if err := c.plugin.Configure(home, conf); err != nil {
			return err
		}
	}

	// make sure that chain id given during chain.New() has the most priority.
	if conf.Genesis != nil {
		conf.Genesis["chain_id"] = chainID
	}

	// update genesis file with the genesis values defined in the config
	if err := c.updateGenesisFile(conf.Genesis); err != nil {
		return err
	}

	return nil
}

// InitAccounts initializes the chain accounts and creates validator gentxs
func (c *Chain) InitAccounts(ctx context.Context, conf *chainconfig.Config) error {
	commands, err := c.Commands(ctx)
	if err != nil {
		return err
	}

	var accounts accountview.Accounts

	// add accounts from config into genesis
	for _, account := range conf.Accounts {
		var generatedAccount chaincmdrunner.Account
		accountAddress := account.Address

		// If the account doesn't provide an address, we create one
		if accountAddress == "" {
			generatedAccount, err = commands.AddAccount(ctx, account.Name, account.Mnemonic, account.CoinType)
			if err != nil {
				return err
			}
			accountAddress = generatedAccount.Address
		}

		coins := strings.Join(account.Coins, ",")
		if err := commands.AddGenesisAccount(ctx, accountAddress, coins); err != nil {
			return err
		}

		if account.Address == "" {
			accounts = append(accounts, accountview.NewAccount(
				generatedAccount.Name,
				accountAddress,
				accountview.WithMnemonic(generatedAccount.Mnemonic),
			))
		} else {
			accounts = append(accounts, accountview.NewAccount(account.Name, accountAddress))
		}
	}

	c.ev.Send("🗂  Initialize accounts...")
	c.ev.SendView(accounts)

	// In the case where there are no validators, assume we are just building a genesis with no gentx.
	// Gentx can be added later in the case of using ignite network.
	if conf.Validators != nil {
		_, err = c.IssueGentx(ctx, createValidatorFromConfig(conf))
	}

	return err
}

// CheckConfigAccounts verifies that the config is set up in the proper manner to be used for an initial genesis
// and adds to genesis
func (c *Chain) CheckConfigAccounts(ctx context.Context, conf *chainconfig.Config) error {
	for _, account := range conf.Accounts {
		if account.Mnemonic != "" {
			return fmt.Errorf("account %s must not have mnemonic displayed.  Do not distribute private account "+
				"mnemonics", account.Name)
		}

		if account.Address == "" {
			return fmt.Errorf("account %s must have address field populated", account.Name)
		}
	}

	return nil
}

// IssueGentx generates a gentx from the validator information in chain config and import it in the chain genesis
func (c Chain) IssueGentx(ctx context.Context, v Validator) (string, error) {
	commands, err := c.Commands(ctx)
	if err != nil {
		return "", err
	}

	// create the gentx from the validator from the config
	gentxPath, err := c.plugin.Gentx(ctx, commands, v)
	if err != nil {
		return "", err
	}

	// import the gentx into the genesis
	return gentxPath, commands.CollectGentxs(ctx)
}

// IsInitialized checks if the chain is initialized
// the check is performed by checking if the gentx dir exist in the config
func (c *Chain) IsInitialized() (bool, error) {
	home, err := c.Home()
	if err != nil {
		return false, err
	}
	gentxDir := filepath.Join(home, "config", "gentx")

	if _, err := os.Stat(gentxDir); os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		// Return error on other error
		return false, err
	}

	return true, nil
}

func (c Chain) updateGenesisFile(data map[string]interface{}) error {
	path, err := c.GenesisPath()
	if err != nil {
		return err
	}

	genesis := make(map[string]interface{})
	cf := confile.New(confile.DefaultJSONEncodingCreator, path)
	if err := cf.Load(&genesis); err != nil {
		return err
	}

	if err := mergo.Merge(&genesis, data, mergo.WithOverride); err != nil {
		return err
	}

	if err = cf.Save(genesis); err != nil {
		return err
	}

	return nil
}

type Validator struct {
	Name                    string
	Moniker                 string
	StakingAmount           string
	CommissionRate          string
	CommissionMaxRate       string
	CommissionMaxChangeRate string
	MinSelfDelegation       string
	GasPrices               string
	Details                 string
	Identity                string
	Website                 string
	SecurityContact         string
}

// Account represents an account in the chain.
type Account struct {
	Name     string
	Address  string
	Mnemonic string `json:"mnemonic"`
	CoinType string
	Coins    string
}

func createValidatorFromConfig(conf *chainconfig.Config) (validator Validator) {
	// Currently, we support the config file with one valid validator.
	validatorFromConfig := conf.Validators[0]
	validator.Name = validatorFromConfig.Name
	validator.StakingAmount = validatorFromConfig.Bonded

	if validatorFromConfig.Gentx != nil {
		if validatorFromConfig.Gentx.Amount != "" {
			validator.StakingAmount = validatorFromConfig.Gentx.Amount
		}
		if validatorFromConfig.Gentx.Moniker != "" {
			validator.Moniker = validatorFromConfig.Gentx.Moniker
		}
		if validatorFromConfig.Gentx.CommissionRate != "" {
			validator.CommissionRate = validatorFromConfig.Gentx.CommissionRate
		}
		if validatorFromConfig.Gentx.CommissionMaxRate != "" {
			validator.CommissionMaxRate = validatorFromConfig.Gentx.CommissionMaxRate
		}
		if validatorFromConfig.Gentx.CommissionMaxChangeRate != "" {
			validator.CommissionMaxChangeRate = validatorFromConfig.Gentx.CommissionMaxChangeRate
		}
		if validatorFromConfig.Gentx.GasPrices != "" {
			validator.GasPrices = validatorFromConfig.Gentx.GasPrices
		}
		if validatorFromConfig.Gentx.Details != "" {
			validator.Details = validatorFromConfig.Gentx.Details
		}
		if validatorFromConfig.Gentx.Identity != "" {
			validator.Identity = validatorFromConfig.Gentx.Identity
		}
		if validatorFromConfig.Gentx.Website != "" {
			validator.Website = validatorFromConfig.Gentx.Website
		}
		if validatorFromConfig.Gentx.SecurityContact != "" {
			validator.SecurityContact = validatorFromConfig.Gentx.SecurityContact
		}
		if validatorFromConfig.Gentx.MinSelfDelegation != "" {
			validator.MinSelfDelegation = validatorFromConfig.Gentx.MinSelfDelegation
		}
	}
	return validator
}
