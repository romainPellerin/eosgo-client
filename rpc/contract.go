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
	"eosgo-client/common"
)

/**
  The following methods are implementing low level contracts that you can control with EOS CLI (cleos):
	- newaccount: cleos create account
	- //TODO: setcontract: cleos set contract
 */

func ContractNewAccount(creator string, accountName string, ownerKey string, activeKey string, recoveryKey string) (*model.Transaction, *errors.AppError) {

	if activeKey == "" {
		activeKey = ownerKey
	}

	if recoveryKey == "" {
		recoveryKey = ownerKey
	}

	auth := model.Authorization{
		common.Config.NODE_PRODUCER_NAME,
		"active",
	}

	ownerAuthority := model.NewAuthority(ownerKey, 1)
	activeAuthority := model.NewAuthority(activeKey, 1)
	recoveryAuthority := model.NewAuthority(recoveryKey, 1)

	action := model.Action{
		common.Config.NODE_PRODUCER_NAME,
		common.Config.NODE_PRODUCER_NAME,
		"newaccount",
		[]string{common.Config.NODE_PRODUCER_NAME, creator},
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
		0,
		0,
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
	trxPushed, err := ChainPushTransaction(trx, []string{common.Config.NODE_PUB_KEY}, "")

	if err != nil {
		return nil, err
	}

	return trxPushed, nil
}
