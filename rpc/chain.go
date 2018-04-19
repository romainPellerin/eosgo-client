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
	"eosgo-client/common"
	"eosgo-client/network"
	"eosgo-client/errors"
	"encoding/json"
	"strings"
	"strconv"
	"time"
	"fmt"
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

func ChainGetInfo() (*model.ChainInfo, *errors.AppError){

	data, err := network.Get(common.EOS_URL+"/v1/chain/get_info", map[string]string{})

	if err != nil {
		return nil, err
	}


	chainInfo := model.ChainInfo{}
	errM := json.Unmarshal(data, &chainInfo)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return &chainInfo, nil
}

func ChainGetBlock(blockNumOrId string) (*model.Block, *errors.AppError){

	data, err := network.Post(common.EOS_URL+"/v1/chain/get_block", map[string]interface{}{"block_num_or_id":blockNumOrId}, nil)

	if err != nil {
		return nil, err
	}

	block := model.Block{}
	errM := json.Unmarshal(data, &block)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return &block, nil
}

func ChainGetAccount(accountName string) (*model.Account, *errors.AppError){

	data, err := network.Post(common.EOS_URL+"/v1/chain/get_account", map[string]interface{}{"account_name":accountName}, nil)

	if err != nil {
		return nil, err
	}

	account := model.Account{}
	errM := json.Unmarshal(data, &account)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return &account, nil
}


func ChainGetCode(accountName string) (*model.Code, *errors.AppError){

	data, err := network.Post(common.EOS_URL+"/v1/chain/get_code", map[string]interface{}{"account_name":accountName}, nil)

	if err != nil {
		return nil, err
	}

	code := model.Code{}
	errM := json.Unmarshal(data, &code)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return &code, nil
}

func ChainGetTableRows(scope string, code string, table string, toJSON bool, lowerBound int, upperBound int, limit int ) (*model.TableRows, *errors.AppError){

	_toJSON := "false"

	if toJSON {
		_toJSON = "true"
	}

	params := map[string]interface{}{
		"table": table,
		"scope":scope,
		"code":code,
		"json": _toJSON,
		"lower_bound": lowerBound,
		"upper_bound": upperBound,
		"limit": limit,
	}

	data, err := network.Post(common.EOS_URL+"/v1/chain/get_table_rows", params, nil)

	if err != nil {
		return nil, err
	}

	tableRows := model.TableRows{}
	errM := json.Unmarshal(data, &tableRows)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return &tableRows, nil
}

func ChainAbiJSONToBin(abiJSON *model.AbiJSON) (*model.Abi, *errors.AppError){

	bin, err := model.AbiJSONToBytes(abiJSON)

	data, err := network.Post(common.EOS_URL+"/v1/chain/abi_json_to_bin", nil, bin)

	if err != nil {
		return nil, err
	}

	abiBin := model.AbiBin{}
	errM := json.Unmarshal(data, &abiBin)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	abi := model.Abi {
		*abiJSON,
		abiBin,
	}

	return &abi, nil
}

func ChainAbiBinToJSON(abi *model.Abi) (*model.Abi, *errors.AppError){

	bin, err := model.AbiToBytes(abi)

	data, err := network.Post(common.EOS_URL+"/v1/chain/abi_bin_to_json", nil, bin)

	if err != nil {
		return nil, err
	}

	abiJSON := model.AbiJSON{}

	errM := json.Unmarshal(data, &abiJSON)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	abi.AbiJSON = abiJSON

	return abi, nil
}


func ChainPushTransaction(trx model.Transaction, pubKeys []string , chainId string) (*model.Transaction, *errors.AppError) {

	chainInfo, err := ChainGetInfo()
	block, err := ChainGetBlock(strconv.Itoa(chainInfo.LastIrreversibleBlockNum))

	trx.RefBlockNum = chainInfo.LastIrreversibleBlockNum
	trx.RefBlockPrefix = int64(block.RefBlockPrefix)

	actions:= trx.Actions

	// calculate expiration date
	time := time.Now().UTC().Add(time.Duration(common.Config.TRANSACTION_EXPIRATION_DELAY*1000*1000*1000))
	trx.Expiration = time.Format("2006-01-02T15:04:05")

	fmt.Println("time:", trx.Expiration)

	// calculate HEX data for each action
	for i := 0;  i<len(trx.Actions); i++ {

		abiJSON := model.AbiJSON{
			trx.Actions[i].Account,
			trx.Actions[i].Name,
			trx.Actions[i].Args,
		}

		data, err := ChainAbiJSONToBin(&abiJSON)

		if err != nil {
			return nil, err
		}

		trx.Actions[i].Data = string(data.Binargs)
	}

	// sign transaction
	trxSigned, err := WalletSignTransaction(trx, pubKeys, chainId)

	if err != nil {
		return nil, err
	}

	trx.Signatures = trxSigned.Signatures
	trx.Actions = actions

	// encode trx

	raw, err := model.TransactionToJSON(&trx)

	if err != nil {
		return nil, err
	}

	// encode signatures
	signatures := "["

	for i:=0; i<len(trx.Signatures); i++ {

		signatures += "\""+trx.Signatures[i]+"\""

		if i+1 < len(trx.Signatures) {
			signatures += ","
		}
	}

	signatures += "]"

	trxData, err := network.PostRawData(common.EOS_URL+"/v1/chain/push_transaction", "{\"signatures\":"+signatures+",\"transaction\":"+raw+",\"compression\":\"none\"}")

	if err != nil {
		return nil, err
	}

	var trxDataMap map[string]interface{}
	dec := json.NewDecoder(strings.NewReader(string(trxData)))
	errD := dec.Decode(&trxDataMap)

	if errD != nil {
		return  nil, errors.NewAppError(nil, "cannot parse transaction data from: "+string(trxData), -1, nil)
	}

	trx.ID = trxDataMap["transaction_id"].(string)

	return &trx, nil
}


func ChainGetRequiredKeys(trx *model.Transaction) ([]string, *errors.AppError){

	myPubKeys, err := WalletGetPublicKeys()

	if err != nil {
		return nil, err
	}

	// encode signatures
	pubKeys := "["

	for i:=0; i<len(myPubKeys); i++ {
		pubKeys += "\"" + myPubKeys[i] + "\""

		if i+1 < len(myPubKeys) {
			pubKeys += ","
		}
	}

	pubKeys += "]"

	toJson, err := model.TransactionToJSON(trx)

	if err != nil {
		return nil, err
	}

	data, err := network.PostRawData(common.EOS_URL+"/v1/chain/get_required_keys", "{\"transaction\":"+toJson+",\"available_keys\":"+pubKeys+"}")

	if err != nil {
		return nil, err
	}

	keys := map[string][]string{}
	errM := json.Unmarshal(data, &keys)

	if errM != nil {
		return nil, errors.NewAppError(nil, "cannot parse result", -1, nil)
	}

	return keys["required_keys"], nil
}