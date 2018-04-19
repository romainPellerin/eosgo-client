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

package common

import (
	"fmt"
	"encoding/json"
	"os"
	"strconv"
)

type Configuration struct {
	NODE_PRODUCER_NAME			 string
	NODE_PUB_KEY				 string
	ENV_EOS_SRC_PATH			 string
	ENV_EOSGO_PATH				 string
	API_PORT                     int
	API_URL                      string
	API_METHOD                   string
	LOGGING_MODE                 string
	WALLET_NAME                  string
	WALLET_PRIV_KEY              string
	TRANSACTION_EXPIRATION_DELAY int
}

var (
	Config  Configuration
	EOS_URL string
)

func ConfigInit(file *os.File) {

	decoder := json.NewDecoder(file)
	Config = Configuration{}
	err := decoder.Decode(&Config)

	if err != nil {
		fmt.Println("error:", err)
		panic("unable to find application config file")
	}

	EOS_URL = Config.API_METHOD + "://" + Config.API_URL + ":" + strconv.Itoa(Config.API_PORT)
}
