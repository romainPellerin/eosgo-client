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
	"eosgo-client/model"
)

func TChainConfig() {
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

// TODO: the following tests require to manually create accounts and keys
// TODO: some keys are hardcoded
// TODO: it will soon be updated to be standalone

func TestChainGetInfo(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	chainInfo, err := ChainGetInfo()

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	//fmt.Println("getInfo: ", *getInfo)
	assert.NotNil(t, chainInfo, "test getInfo, should not be nil")

	fmt.Println("getInfo.HeadBlockID: ", chainInfo)
}

func TestChainGetBlock(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	block, err := ChainGetBlock("1")

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	//fmt.Println("block: ", *block)
	assert.NotNil(t, block, "test block, should not be nil")

	fmt.Println("block: ", block)
}

// this require to create eosio account
func TestChainGetAccount(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	account, err := ChainGetAccount("token")

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	//fmt.Println("account: ", *account)
	assert.NotNil(t, account, "test account, should not be nil")

	fmt.Println("account.Name: ", account.Name)
}

// this require to create eosio account
func TestChainGetCode(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	code, err := ChainGetCode("token")

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	//fmt.Println("code: ", *code)
	assert.NotNil(t, code, "test account, should not be nil")

	fmt.Println("code.CodeHash: ", code.CodeHash)
}

func TestChainGetTableRows(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	tableRows, err := ChainGetTableRows("token", "token", "stat", true, -1, -1, -1)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	if tableRows != nil {
		//fmt.Println("tableRows: ", *tableRows)
		assert.NotNil(t, tableRows, "test tableRows, should not be nil")

		fmt.Println("nb tableRows rows: ", len(tableRows.Rows))
	}
}

func TestChainAbi(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	abiJSON := model.AbiJSON{
		"token",
		"issue",
		map[string]interface{}{"to": "accountbe", "quantity": "2.0001 EOS", "memo": "just a coin"},
	}

	abi, err := ChainAbiJSONToBin(&abiJSON)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	if abi != nil {
		//fmt.Println("abiBin: ", *abiBin)
		assert.NotNil(t, abi, "test tableRows, should not be nil")

		fmt.Println("nb abiBin.Binargs: ", len(abi.Binargs))

		// save then reinit args
		args := abiJSON.Args

		abiJSON.Args = map[string]interface{}{}
		abiBin := model.AbiBin{
			abi.Binargs,
			[]interface{}{},
			[]interface{}{},
		}

		abi := model.Abi{
			abiJSON,
			abiBin,
		}

		abiResult, err := ChainAbiBinToJSON(&abi)

		fmt.Println("err: ", err)
		assert.Nil(t, err, "test get error")
		assert.Equal(t, args["to"], abiResult.Args["to"], "test saved and retrieved to")

	}
}

func TestChainPushTransaction(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	auth := model.Authorization{
		"eow",
		"active",
	}

	action := model.Action{
		"eow",
		"token",
		"issue",
		[]string{"token", "accountbe"},
		[]model.Authorization{auth},
		"",
		map[string]interface{}{"to": "accountbe", "quantity": "2.0001 EOS", "memo": ""},
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

	trxPushed, err := ChainPushTransaction(trx, []string{"EOS7gzcNK4rbvhHRzjar4D5zM9CEw4db7gCXg9CXyVwa21wVb9JGn"}, "")

	// check required keys

	if err != nil {
		fmt.Println("err: ", err)
	}

	assert.Nil(t, err, "test get error")

	if trxPushed != nil {
		assert.NotEqual(t, "", trxPushed.ID, "id should not be empty")
		fmt.Println("transaction id: ", trxPushed.ID)

	}

	keys, err := ChainGetRequiredKeys(trxPushed)

	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("required keys:", keys)

}
