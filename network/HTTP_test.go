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

package network

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"fmt"
	"eosgo-client/model"
	"encoding/json"
	"eosgo-client/common"
	"os"
	"time"
	"strconv"
)

func THTTPConfig() {

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

// test HTTP GET method
func TestHTTP_GetOK(t *testing.T) {

	if common.Config.API_URL == "" {
		THTTPConfig()
	}

	data, err := Get(common.EOS_URL+"/v1/chain/get_info", map[string]string{})

	fmt.Println("data: ", string(data))
	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	if err != nil {
		fmt.Println("httpErr: ", err.Custom)
		assert.Nil(t, err.Custom, "test http error, should be nil")
	}

	chain := model.ChainInfo{}
	errM := json.Unmarshal(data, &chain)
	assert.Nil(t, errM, "test get error")

	fmt.Println("getInfo.HeadBlockID:", chain.HeadBlockID)
	assert.NotNil(t, chain.HeadBlockID, "test result, should not be nil")

}

// test HTTP error handling via GET method
func TestHTTP_GetError(t *testing.T) {

	if common.Config.API_URL == "" {
		THTTPConfig()
	}

	data, err := Get(common.EOS_URL+"/v1/chain/get_info", map[string]string{"aé%": "b", "c": "dé@#"})

	fmt.Println("err: ", err)
	assert.NotNil(t, err, "test get error")

	if err != nil {
		assert.NotNil(t, err.Custom, "test http error, should not be nil")
	}

	fmt.Println("data: ", data)
	assert.Nil(t, data, "test result, should  be nil")

}

// test HTTP Post Method with map or bin data
func TestHTTP_Post(t *testing.T) {

	if common.Config.API_URL == "" {
		THTTPConfig()
	}

	data, err := Post(common.EOS_URL+"/v1/chain/get_block", map[string]interface{}{"block_num_or_id": "1"}, nil)

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	if err != nil {
		fmt.Println("httpErr: ", )
		assert.Nil(t, err.Custom, "test http error, should not be nil")
	}

	block := model.Block{}
	errM := json.Unmarshal(data, &block)
	assert.Nil(t, errM, "test get error")

	assert.NotNil(t, block, "test result, should  be nil")
	fmt.Println("block.Timestamp: ", block.Timestamp)
}

// test HTTP Post Method with raw data
func TestHTTP_PostRawData(t *testing.T) {

	if common.Config.API_URL == "" {
		THTTPConfig()
	}

	priv_key, err := PostRawData(common.EOS_URL+"/v1/wallet/create", "test"+strconv.Itoa(int(time.Now().Unix())))

	fmt.Println("err: ", err)
	assert.Nil(t, err, "test get error")

	if err != nil {
		fmt.Println("httpErr: ", err.Custom)
		assert.Nil(t, err.Custom, "test http error, should be nil")
	}

	assert.NotNil(t, priv_key, "test result, should not be nil")
	fmt.Println("priv_key: ", priv_key)
}
