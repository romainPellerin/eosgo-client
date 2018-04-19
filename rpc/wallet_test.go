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
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"os"
	"eosgo-client/common"
	"time"
	"eosgo-client/model"
	"eosgo-client/errors"
)

var (
	wallet  string
	privKey string
)

func TWalletConfig() {
	path := os.Getenv("EOSGO_PATH")
	confFile := os.Getenv("EOSGO_CONF")
	uri := path + "/" + confFile + ".conf"

	if path == "" {
		panic("env var EOSGO_PATH not found")
	}

	if confFile == "" {
		panic("env var EOSGO_CONF not found")
	}

	fmt.Println("loading config file: " + uri)

	// load config file
	file, _ := os.Open(uri)

	common.ConfigInit(file)

	// init common logger
	common.LoggerInit("debug")
}

func TestWalletInit(t *testing.T) {
	TWalletConfig()
}

func TestWalletCreate(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	wallet = common.ToolsWalletGenerateName("eosgowallet")

	priv_key, err := WalletCreate(wallet)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	assert.NotEqual(t, "", priv_key, "test priv_key, should not be empty")

	privKey = priv_key

	fmt.Println("priv_key: ", priv_key)
}

func TestWalletOpen(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	fmt.Println("opening wallet: ", wallet)

	err := WalletOpen(wallet)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")
}

func TestWalletUnlock(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	fmt.Println("wallet: ", wallet)

	err := WalletUnlock(wallet, privKey)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")
}

// TODO: the following test requires to manually create accounts and keys
// TODO: some keys are hardcoded
// TODO: it will soon be updated to be standalone
func TestWalletImportKey(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	err := WalletImportKey(wallet, "5J5hbKx5SoBRXokkDzNY3Zdu1PnG4dNsFRZhLE3fQPhxPATX5jG")

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")
}

func TestWalletList(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	opened, closed, err := WalletList()

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	fmt.Println("opened wallets ", opened)
	fmt.Println("closed wallets ", closed)

}

func TestWalletListKeys(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	keys, err := WalletListKeys()

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	fmt.Println("keys: ", keys)
}

func TestWalletGetPublicKeys(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	keys, err := WalletGetPublicKeys()

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	fmt.Println("keys: ", keys)
}

func TestWalletLock(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	err := WalletLock(wallet)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")
}

func TestWalletLockAll(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	err := WalletLockAll()

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")
}

func TestWalletSetTimeout(t *testing.T) {

	walletFromConf := false
	var err *errors.AppError

	if common.Config.API_URL == "" {
		TWalletConfig()
		walletFromConf = true
	}

	fmt.Println("walletFromConf: ", walletFromConf)

	if walletFromConf {
		err = WalletUnlock(common.Config.WALLET_NAME, common.Config.WALLET_PRIV_KEY)
	} else {
		err = WalletUnlock(wallet, privKey)
	}

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	err = WalletSetTimeout(60 * 60)

	assert.Nil(t, err, "test get error")
}

func TestWalletSignTransaction(t *testing.T) {

	if common.Config.API_URL == "" {
		TWalletConfig()
	}

	auth := model.Authorization{
		common.Config.NODE_PRODUCER_NAME,
		"active",
	}

	action := model.Action{
		common.Config.NODE_PRODUCER_NAME,
		common.Config.NODE_PRODUCER_NAME,
		"issue",
		[]string{common.Config.NODE_PRODUCER_NAME, "eosgo"},
		[]model.Authorization{auth},
		"",
		map[string]interface{}{"to": "eosgo", "quantity": "2.0001 EOS", "memo": ""},
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

	time := time.Now().UTC().Add(time.Duration(common.Config.TRANSACTION_EXPIRATION_DELAY * 1000 * 1000 * 1000))
	trx.Expiration = time.Format("2006-01-02T15:04:05")
	fmt.Println("time:", trx.Expiration)

	// sign transaction
	trxUpdated, err := WalletSignTransaction(trx, []string{common.Config.NODE_PUB_KEY}, "")

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	assert.NotEqual(t, 0, trxUpdated, "transaction signatures should not be empty")
	fmt.Println("transaction: ", trxUpdated)
}
