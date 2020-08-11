package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// AragonBech32Prefix defines the Bech32 prefix used for EthAccounts on Aragon Chain
	AragonBech32Prefix = "aragon"

	// Bech32PrefixAccAddr defines the Bech32 prefix of an account's address
	Bech32PrefixAccAddr = AragonBech32Prefix
	// Bech32PrefixAccPub defines the Bech32 prefix of an account's public key
	Bech32PrefixAccPub = AragonBech32Prefix + sdk.PrefixPublic
	// Bech32PrefixValAddr defines the Bech32 prefix of a validator's operator address
	Bech32PrefixValAddr = AragonBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator
	// Bech32PrefixValPub defines the Bech32 prefix of a validator's operator public key
	Bech32PrefixValPub = AragonBech32Prefix + sdk.PrefixValidator + sdk.PrefixOperator + sdk.PrefixPublic
	// Bech32PrefixConsAddr defines the Bech32 prefix of a consensus node address
	Bech32PrefixConsAddr = AragonBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus
	// Bech32PrefixConsPub defines the Bech32 prefix of a consensus node public key
	Bech32PrefixConsPub = AragonBech32Prefix + sdk.PrefixValidator + sdk.PrefixConsensus + sdk.PrefixPublic
)

// SetBech32Prefixes sets the global prefixes to be used when serializing addresses and public keys to Bech32 strings.
func SetBech32Prefixes(config *sdk.Config) {
	config.SetBech32PrefixForAccount(Bech32PrefixAccAddr, Bech32PrefixAccPub)
	config.SetBech32PrefixForValidator(Bech32PrefixValAddr, Bech32PrefixValPub)
	config.SetBech32PrefixForConsensusNode(Bech32PrefixConsAddr, Bech32PrefixConsPub)
}
