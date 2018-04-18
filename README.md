# README #

### Overview ###

eosgo-client is a simple Go/Golang wrapper for EOS blockchain (https://eos.io).

It wraps the nodeos RPC API and will offer a high level set of API to simplify the development on top of EOS.

### Releases ###

- v0.0.2: implements standalone test cases and more contracts (soon)
- v0.0.1: fully functional wrapper of nodeos RPC API, see https://eosio.github.io/eos/group__eosiorpc.html for complete specs

### Current features ###

*Chain API* (see rpc/chain.go)
- get_info
- get_block
- get_account
- get_code
- get_table_rows
- abi_json_to_bin
- abi_bin_to_json
- push_transaction
- push_transactions
- get_required_keys

*Wallet API* (see rpc/wallet.go)
- wallet_create
- wallet_open
- wallet_lock
- wallet_lock_all
- wallet_import_key
- wallet_list
- wallet_list_keys
- wallet_get_public_keys
- wallet_set_timeout
- wallet_sign_trx

*EOSIO contracts* (see rpc/contracts.go)
- newaccount
