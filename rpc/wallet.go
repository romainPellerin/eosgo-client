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
	"encoding/json"
	"github.com/romainPellerin/eosgo-client/common"
	"github.com/romainPellerin/eosgo-client/errors"
	"github.com/romainPellerin/eosgo-client/model"
	"github.com/romainPellerin/eosgo-client/network"
	"strconv"
	"strings"
)

/**
	See if you have eos source {$EOS_SOURCE}/docs/group__eosiorpc.html#walletrpc
	or download from https://eosio.github.io/eos/group__eosiorpc.html
	for detailed specs of:
	- Create
	- Open
	- Lock
	- Unlock
	- LockAll
	- ImportKey
    - ListWallets
	- ListKeys
	- GetPublicKeys
	- SetTimeout
	- SignTrx
 */

func WalletCreate(name string) (string, *errors.AppError) {


	data, err := network.PostRawData(common.EOS_URL+"/v1/wallet/create", name)

	if err != nil {
		return "", err
	}

	if data == nil {
		return "", errors.NewAppError(nil, "empty response, no key returned by nodeos", -1, nil)
	}

	return string(data), nil
}

func WalletOpen(name string) (*errors.AppError) {

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/open", name)

	if err != nil {
		return err
	}

	return nil
}

func WalletUnlock(name string, privKey string) (*errors.AppError) {

	if !strings.HasPrefix(privKey,"\"") {
		privKey = "\""+privKey+"\""
	}

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/unlock", "[\""+name+"\","+privKey+"]")

	if err != nil {
		return err
	}

	return nil
}

func WalletLock(name string) (*errors.AppError) {

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/lock", name)

	if err != nil {
		return err
	}

	return nil
}

func WalletLockAll() (*errors.AppError) {

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/lock_all", "")

	if err != nil {
		return err
	}

	return nil
}

func WalletImportKey(name string, privKey string) (*errors.AppError) {

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/import_key", "[\""+name+"\",\""+privKey+"\"]")

	if err != nil {
		return err
	}

	return nil
}

func WalletCreateKey(name string, privKey string) (*errors.AppError) {

	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/import_key", "["+name+","+privKey+"]")

	if err != nil {
		return err
	}

	return nil
}

// return array of opened wallets, array of closed wallets
func WalletList() ([]string, []string, *errors.AppError) {

	wallets, err := network.Get(common.EOS_URL+"/v1/wallet/list_wallets", map[string]string{})

	if err != nil {
		return nil, nil, err
	}

	var los []string
	dec := json.NewDecoder(strings.NewReader(string(wallets)))
	errD := dec.Decode(&los)

	if errD != nil {
		return nil, nil, errors.NewAppError(nil, "cannot parse wallets from: "+string(wallets), -1, nil)
	}

	var openedWallets []string
	var closedWallets []string

	suffix := " *" // nodeos used this suffix to distinguish opened wallets

	for i := 0; i < len(los); i++ {
		opened := strings.HasSuffix(los[i], suffix)
		if opened {
			openedWallets = append(openedWallets, los[i][:len(los[i])-len(suffix)])
		} else {
			closedWallets = append(closedWallets, los[i])
		}
	}

	return openedWallets, closedWallets, nil
}

func WalletListKeys() ([][]string, *errors.AppError) {

	wallets, err := network.Get(common.EOS_URL+"/v1/wallet/list_keys", map[string]string{})

	if err != nil {
		return nil, err
	}

	var los [][]string
	dec := json.NewDecoder(strings.NewReader(string(wallets)))
	errD := dec.Decode(&los)

	if errD != nil {
		return nil, errors.NewAppError(nil, "cannot parse keys from: "+string(wallets), -1, nil)
	}

	return los, nil
}

func WalletGetPublicKeys() ([]string, *errors.AppError) {

	wallets, err := network.Get(common.EOS_URL+"/v1/wallet/get_public_keys", map[string]string{})

	if err != nil {
		return nil, err
	}

	var los []string
	dec := json.NewDecoder(strings.NewReader(string(wallets)))
	errD := dec.Decode(&los)

	if errD != nil {
		return nil, errors.NewAppError(nil, "cannot parse keys from: "+string(wallets), -1, nil)
	}

	return los, nil
}

func WalletSetTimeout(seconds int64) (*errors.AppError) {
	_, err := network.PostRawData(common.EOS_URL+"/v1/wallet/set_timeout", strconv.Itoa(int(seconds)))

	if err != nil {
		return err
	}

	return nil
}

func WalletSignTransaction(trx model.Transaction, pubKeys []string, chainId string) (*model.Transaction, *errors.AppError) {

	trxJson, err := model.TransactionToJSON(&trx)

	// encode keys
	pubKeysRaw := "["

	for i := 0; i < len(pubKeys); i++ {

		pubKeysRaw += "\"" + pubKeys[i] + "\""

		if i+1 < len(pubKeys) {
			pubKeysRaw += ","
		}
	}

	pubKeysRaw += "]"

	raw := "[" + trxJson + "," + pubKeysRaw + ",\"" + chainId + "\"]"

	if err != nil {
		return nil, err
	}

	trxData, err := network.PostRawData(common.EOS_URL+"/v1/wallet/sign_transaction", raw)

	if err != nil {
		return nil, err
	}

	trxUpdated, err := model.JSONToTransaction(string(trxData))

	if err != nil {
		return nil, err
	}

	return trxUpdated, nil
}
