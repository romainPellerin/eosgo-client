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
)

func TContractConfig() {
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

// TODO: these tests require to manually create accounts and keys
// TODO: some keys are hardcoded
// TODO: it will soon be updated to be standalone

func TestContractNewAccount(t *testing.T) {

	if common.Config.API_URL == "" {
		TContractConfig()
	}

	account := common.ToolsAccountGenerateName("eosgoeosgo")

	fmt.Println("account: "+account)

	trx, err := ContractNewAccount(common.Config.NODE_PRODUCER_NAME, account, common.Config.NODE_PUB_KEY, "", "")

	if err != nil {
		fmt.Println("err: ", err)
	}

	if err != nil {
		fmt.Println("err: ", err)
		t.FailNow()
	}

	assert.Nil(t, err, "test get error")

	if trx != nil {
		assert.NotEqual(t, "", trx.ID, "id should not be empty")
		fmt.Println("transaction id: ", trx.ID)

	}

}
