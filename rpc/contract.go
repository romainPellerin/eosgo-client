/**
 *  @copyright defined in eosgo-client/LICENSE.txt
 *  @author Romain Pellerin - romain@slyseed.com
 *
 *  Donation appreciated :)
 *  EOS wallet: 0x8d39307d9a0687c894058115365f6ad0f03b9ff9
 *	ETH wallet: 0x317b60152f0a90c10cea52d751ccb4dfad2fe9e7
 *  BTC wallet: 3HdFx8C3WNA6RyyGYKygECGbLD1BxyeqTN
 *  BCH wallet: 14e9fnmcHSZxxd3fnrkaG6EYph9JuWcF9T
 */

package rpc

import (
	"eosgo-client/model"
	"eosgo-client/errors"
)

/**
	See if you have EOS source {$EOS_SOURCE}/docs/group__eosiorpc.html#chainrpc for detailed specs of:
	or download from https://eosio.github.io/eos/group__eosiorpc.html
	- GetInfo
	- GetBlock
	- GetAccount
	- GetCode
	- GetTableRows
	- AbiJSONToBin
	- AbiBinToJSON
	- PushTransaction
	- //TODO: PushTransactions
	- GetRequiredKeys

 */

func ContractNewAccount(creator string, accountName string, ownerKey string, activeKey string, recoveryKey string) (*model.Transaction, *errors.AppError) {

	if activeKey == "" {
		activeKey = ownerKey
	}

	if recoveryKey == "" {
		recoveryKey = ownerKey
	}

	auth := model.Authorization{
		"eosio",
		"active",
	}

	ownerAuthority := model.NewAuthority(ownerKey, 1)
	activeAuthority := model.NewAuthority(activeKey, 1)
	recoveryAuthority := model.NewAuthority(recoveryKey, 1)

	action := model.Action{
		"eosio",
		"eosio",
		"newaccount",
		[]string{"eosio", creator},
		[]model.Authorization{auth},
		"",
		map[string]interface{}{
			"creator":  creator,
			"name":     accountName,
			"owner":    ownerAuthority,
			"active":   activeAuthority,
			"recovery": recoveryAuthority,
		},
	}

	trx := model.Transaction{
		49344,
		4171690928,
		0,
		"",
		[]string{},
		[]string{},
		[]model.Action{action},
		[]string{},
		[]model.Authorization{auth},
		"",
		0,
		0,
		0,
		"",
		"none",
		[]map[string]interface{}{},
		0,
		0,
		[]model.Action{},
	}

	//right call
	trxPushed, err := ChainPushTransaction(trx, []string{"EOS6MRyAjQq8ud7hVNYcfnVPJqcVpscN5So8BhtHuGYqET5GDW5CV"}, "")

	if err != nil {
		return nil, err
	}

	return trxPushed, nil
}
