<!--
Guiding Principles:

Changelogs are for humans, not machines.
There should be an entry for every single version.
The same types of changes should be grouped.
Versions and sections should be linkable.
The latest version comes first.
The release date of each version is displayed.
Mention whether you follow Semantic Versioning.

Usage:

Change log entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* (<tag>) \#<issue-number> message

The issue numbers will later be link-ified during the release process so you do
not have to worry about including a link manually, but you can if you wish.

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Bug Fixes" for any bug fixes.
"Client Breaking" for breaking CLI commands and REST routes used by end-users.
"API Breaking" for breaking exported APIs used by developers building on SDK.
"State Machine Breaking" for any changes that result in a different AppState given same genesisState and txList.

Ref: https://keepachangelog.com/en/1.0.0/
-->

# Changelog

## [Unreleased]

### Features

* [\#3](https://github.com/aragon/aragon-chain/pull/3) Testnet features:
  * (client) `aragonchaind testnet` command support to create a multi-node testnet with Docker compose.
  * (scripts) `init.sh` script for running a single local node.

### Improvements

* [\#3](https://github.com/aragon/aragon-chain/pull/3) Various improvements:
  * (ethermint)  Bump Ethermint version to [v0.2.0-rc1](https://github.com/ChainSafe/ethermint/releases/tag/v0.2.0-rc1)
  * (types) Define constants for `ARA` token.
  * (types) Update [BIP44](https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki) coin type to `60` to satisfy [EIP84](https://github.com/ethereum/EIPs/issues/84).
* [\#17](https://github.com/aragon/aragon-chain/pull/17) Update testnet cmd:
  * (client) `aragonchaind testnet` command takes multiple ip addresses and adds them to the genesis + configurations.
