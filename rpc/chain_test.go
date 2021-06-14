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
	"fmt"
	"github.com/romainPellerin/eosgo-client/common"
	"github.com/romainPellerin/eosgo-client/model"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
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

func TestChainGetInfo(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	//right call
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

	//right call
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

	//right call
	account, err := ChainGetAccount(common.Config.NODE_PRODUCER_NAME)

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

	//right call
	code, err := ChainGetCode(common.Config.NODE_PRODUCER_NAME)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	//fmt.Println("code: ", *code)
	assert.NotNil(t, code, "test account, should not be nil")

	fmt.Println("code.CodeHash: ", code.CodeHash)
}

// TODO: try with another contract than currency which is kind of managed in another way by nodeos
// this require to create eosio account
func TestChainGetTableRows(t *testing.T) {

	if common.Config.API_URL == "" {
		TChainConfig()
	}

	//right call
	tableRows, err := ChainGetTableRows(common.Config.NODE_PRODUCER_NAME, common.Config.NODE_PRODUCER_NAME, "accounts", true, -1, -1, -1)

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
		common.Config.NODE_PRODUCER_NAME,
		"issue",
		map[string]interface{}{"to": common.ToolsAccountGenerateName("eosgoeosgo"), "quantity": "2.0001 EOS", "memo": "just a coin"},
	}

	//right call
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

	account := common.ToolsAccountGenerateName("eosgo")

	fmt.Println("account name generated: " + account)

	_, err := ContractNewAccount(common.Config.NODE_PRODUCER_NAME, account, common.Config.NODE_PUB_KEY, "", "")

	if err != nil {
		fmt.Println("err: ", err)
		t.FailNow()
	}

	auth := model.Authorization{
		common.Config.NODE_PRODUCER_NAME,
		"active",
	}

	action := model.Action{
		common.Config.NODE_PRODUCER_NAME,
		common.Config.NODE_PRODUCER_NAME,
		"issue",
		[]string{common.Config.NODE_PRODUCER_NAME, account},
		[]model.Authorization{auth},
		"",
		map[string]interface{}{"to": account, "quantity": "2.0001 EOS", "memo": "just 2 coins"},
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
