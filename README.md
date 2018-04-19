# README #

### Overview ###

eosgo-client is a simple Go/Golang wrapper for EOS blockchain (https://eos.io).

It wraps the nodeos RPC API and will offer a high level set of API to simplify the development on top of EOS.

### Releases ###

- v0.0.3: implement more eosio contracts (soon)
- v0.0.2: implements standalone test cases
- v0.0.1: fully functional wrapper of nodeos RPC API, see https://eosio.github.io/eos/group__eosiorpc.html for detailed specs

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

### How to start ###

Configure your own conf file (for exemple *test.conf*) based on *default.conf* one with
```
    "NODE_PRODUCER_NAME":"eosio",               // see eosio/config.ini file, default is eosio
    "NODE_PUB_KEY":"KEY",                       // see eosio/config.ini file,
    "ENV_EOS_SRC_PATH":"YOURPATH/eos",		    // path to your EOS source folder
    "ENV_EOSGO_PATH":"YOURPATH/eosgo-client",   // path to this eosgo-client project
    "API_PORT": 8888,                           // nodeos PORT
    "API_URL": "localhost",                     // nodeos URL
    "API_METHOD": "http",                       // https or https
    "LOGGING_MODE": "STDOUT",                   // STDOUT or SYSLOG
    "LOGGING_LEVEL": "debug",                   // debug, info or error
    "WALLET_NAME":"NAME",                       // your wallet name
    "WALLET_PRIV_KEY":"KEY"                     // your wallet private key
    "TRANSACTION_EXPIRATION_DELAY":30           // default 30 seconds
```

Take a look at rpc/chain_test.go, rpc/wallet_test.go and rpc/contracts_test.go.

Also, you have to create 2 environment vars in order to load your conf file (see default.conf), in your system:
```
export EOSGO_PATH=/your_path/eosgo-client/.
export EOSGO_CONF=default
```
or for your IDE (here for a Goland project):
```
EOSGO_PATH=.;EOSGO_CONF=test
```
