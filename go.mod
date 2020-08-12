module github.com/aragon/aragon-chain

go 1.14

require (
	github.com/cosmos/cosmos-sdk v0.39.1
	github.com/cosmos/ethermint v0.0.0-20190802135314-3f32f9ba8a1f
	github.com/ethereum/go-ethereum v1.9.18
	github.com/spf13/cobra v1.0.0
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.33.7
	github.com/tendermint/tm-db v0.5.1
)

// use ChainSafe fork
replace github.com/cosmos/ethermint => github.com/ChainSafe/ethermint v0.0.0-20200812220222-d732e74bd4c1
