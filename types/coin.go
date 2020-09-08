package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// DefaultDenom defines the default coin denomination used in Aragon Chain in:
	//
	// - Staking parameters: denomination used as stake in the dPoS chain
	// - Mint parameters: denomination minted due to fee distribution rewards
	// - Governance parameters: denomination used for spam prevention in proposal deposits
	// - Crisis parameters: constant fee denomination used for spam prevention to check broken invariant
	// - EVM parameters: denomination used for running EVM state transitions in Ethermint.
	DefaultDenom string = "ara"
)

// NewAra is a utility function that returns an "ara" coin with the given amount.
// The function will panic if the provided amount is negative.
func NewAra(amount int64) sdk.Coin {
	return sdk.NewInt64Coin(DefaultDenom, amount)
}
