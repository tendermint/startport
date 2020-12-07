package spn

import (
	"context"
	"errors"
	"github.com/cosmos/cosmos-sdk/types"
	genesistypes "github.com/tendermint/spn/x/genesis/types"
	"github.com/tendermint/starport/starport/pkg/jsondoc"
)

// ProposalType represents the type of the proposal
type ProposalType string

const (
	ProposalTypeAll          = ""
	ProposalTypeAddAccount   = "add-account"
	ProposalTypeAddValidator = "add-validator"
)

// ProposalStatus represents the status of the proposal
type ProposalStatus string

const (
	ProposalStatusAll      = ""
	ProposalStatusPending  = "pending"
	ProposalStatusApproved = "approved"
	ProposalStatusRejected = "rejected"
)

// Proposal represents a proposal.
type Proposal struct {
	ID        int                   `yaml:",omitempty"`
	Status    ProposalStatus        `yaml:",omitempty"`
	Account   *ProposalAddAccount   `yaml:",omitempty"`
	Validator *ProposalAddValidator `yaml:",omitempty"`
}

// ProposalAddAccount used to propose adding an account.
type ProposalAddAccount struct {
	Address string
	Coins   types.Coins
}

// ProposalAddValidator used to propose adding a validator.
type ProposalAddValidator struct {
	Gentx            jsondoc.Doc
	ValidatorAddress string
	SelfDelegation   types.Coin
	P2PAddress       string
}

// ProposalList lists proposals on a chain by status.
func (c *Client) ProposalList(ctx context.Context, acccountName, chainID string, status ProposalStatus, proposalType ProposalType) ([]Proposal, error) {
	var proposals []Proposal
	var spnProposals []*genesistypes.Proposal

	queryClient := genesistypes.NewQueryClient(c.clientCtx)

	// Dispatch spn proposal status
	var spnStatus genesistypes.ProposalStatus
	switch status {
	case ProposalStatusAll:
		spnStatus = genesistypes.ProposalStatus_ANY_STATUS
	case ProposalStatusPending:
		spnStatus = genesistypes.ProposalStatus_PENDING
	case ProposalStatusApproved:
		spnStatus = genesistypes.ProposalStatus_APPROVED
	case ProposalStatusRejected:
		spnStatus = genesistypes.ProposalStatus_REJECTED
	default:
		return nil, errors.New("unrecognized status")
	}

	// Dispatch spn proposal type
	var spnType genesistypes.ProposalType
	switch proposalType {
	case ProposalTypeAll:
		spnType = genesistypes.ProposalType_ANY_TYPE
	case ProposalTypeAddAccount:
		spnType = genesistypes.ProposalType_ADD_ACCOUNT
	case ProposalTypeAddValidator:
		spnType = genesistypes.ProposalType_ADD_VALIDATOR
	default:
		return nil, errors.New("unrecognized type")
	}

	// Send query
	res, err := queryClient.ListProposals(ctx, &genesistypes.QueryListProposalsRequest{
		ChainID: chainID,
		Status:  spnStatus,
		Type:    spnType,
	})
	if err != nil {
		return nil, err
	}
	spnProposals = res.Proposals

	// Format proposals
	for _, gp := range spnProposals {
		proposal, err := c.toProposal(*gp)
		if err != nil {
			return nil, err
		}

		proposals = append(proposals, proposal)
	}

	return proposals, nil
}

var toStatus = map[genesistypes.ProposalStatus]ProposalStatus{
	genesistypes.ProposalStatus_PENDING:  ProposalStatusPending,
	genesistypes.ProposalStatus_APPROVED: ProposalStatusApproved,
	genesistypes.ProposalStatus_REJECTED: ProposalStatusRejected,
}

func (c *Client) toProposal(proposal genesistypes.Proposal) (Proposal, error) {
	p := Proposal{
		ID:     int(proposal.ProposalInformation.ProposalID),
		Status: toStatus[proposal.ProposalState.GetStatus()],
	}
	switch payload := proposal.Payload.(type) {
	case *genesistypes.Proposal_AddAccountPayload:
		p.Account = &ProposalAddAccount{
			Address: payload.AddAccountPayload.Address.String(),
			Coins:   payload.AddAccountPayload.Coins,
		}

	case *genesistypes.Proposal_AddValidatorPayload:
		p.Validator = &ProposalAddValidator{
			P2PAddress: payload.AddValidatorPayload.Peer,
			Gentx:      payload.AddValidatorPayload.GenTx,
		}
	}

	return p, nil
}

func (c *Client) ProposalGet(ctx context.Context, accountName, chainID string, id int) (Proposal, error) {
	queryClient := genesistypes.NewQueryClient(c.clientCtx)

	// Query the proposal
	param := &genesistypes.QueryShowProposalRequest{
		ChainID:    chainID,
		ProposalID: int32(id),
	}
	res, err := queryClient.ShowProposal(ctx, param)
	if err != nil {
		return Proposal{}, err
	}

	return c.toProposal(*res.Proposal)
}

// ProposalOption configures Proposal to set a spesific type of proposal.
type ProposalOption func(*Proposal)

// AddAccountProposal creates an add account proposal option.
func AddAccountProposal(address string, coins types.Coins) ProposalOption {
	return func(p *Proposal) {
		p.Account = &ProposalAddAccount{address, coins}
	}
}

// AddValidatorProposal creates an add validator proposal option.
func AddValidatorProposal(gentx jsondoc.Doc, validatorAddress string, selfDelegation types.Coin, p2pAddress string) ProposalOption {
	return func(p *Proposal) {
		p.Validator = &ProposalAddValidator{gentx, validatorAddress, selfDelegation, p2pAddress}
	}
}

// Propose proposes given proposals in batch for chainID by using SPN accountName.
func (c *Client) Propose(ctx context.Context, accountName, chainID string, proposals ...ProposalOption) error {
	if len(proposals) == 0 {
		return errors.New("at least one proposal required")
	}

	clientCtx, err := c.buildClientCtx(accountName)
	if err != nil {
		return err
	}

	var msgs []types.Msg

	for _, p := range proposals {
		var proposal Proposal
		p(&proposal)

		switch {
		case proposal.Account != nil:
			addr, err := types.AccAddressFromBech32(proposal.Account.Address)
			if err != nil {
				return err
			}

			// Create the proposal payload
			payload := genesistypes.NewProposalAddAccountPayload(
				addr,
				proposal.Account.Coins,
			)

			msgs = append(msgs, genesistypes.NewMsgProposalAddAccount(
				chainID,
				clientCtx.GetFromAddress(),
				payload,
			))

		case proposal.Validator != nil:
			// Get the validator address
			addr, err := types.AccAddressFromBech32(proposal.Validator.ValidatorAddress)
			if err != nil {
				return err
			}
			validatorAddress := types.ValAddress(addr)

			// Create the proposal payload
			payload := genesistypes.NewProposalAddValidatorPayload(
				proposal.Validator.Gentx,
				validatorAddress,
				proposal.Validator.SelfDelegation,
				proposal.Validator.P2PAddress,
			)

			msgs = append(msgs, genesistypes.NewMsgProposalAddValidator(
				chainID,
				clientCtx.GetFromAddress(),
				payload,
			))
		}
	}

	return c.broadcast(ctx, clientCtx, msgs...)
}
