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

package model

import (
	"encoding/json"
	"github.com/romainPellerin/eosgo-client/errors"
)

type Abi struct {
	AbiJSON
	AbiBin
}

type AbiJSON struct {
	Code   string                 `json:"code"`
	Action string                 `json:"action"`
	Args   map[string]interface{} `json:"args"`
}

type AbiBin struct {
	Binargs       string        `json:"binargs"`
	RequiredScope []interface{} `json:"required_scope"`
	RequiredAuth  []interface{} `json:"required_auth"`
}

func AbiToBytes(obj *Abi) ([]byte, *errors.AppError) {

	bytes, err := json.Marshal(&obj)

	if err != nil {
		return nil, errors.NewAppError(err, "cannot marshal AbiBin", -1, nil)
	}

	return bytes, nil
}

func AbiJSONToBytes(obj *AbiJSON) ([]byte, *errors.AppError) {

	bytes, err := json.Marshal(&obj)

	if err != nil {
		return nil, errors.NewAppError(err, "cannot marshal AbiJSON", -1, nil)
	}

	return bytes, nil
}
