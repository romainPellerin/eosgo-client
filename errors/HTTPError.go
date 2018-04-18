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

package errors

import (
	"encoding/json"
)

type HTTPError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error struct {
		Code int    `json:"code"`
		Name string `json:"name"`
		What string `json:"what"`
		Details []struct {
			Message    string `json:"message"`
			File       string `json:"file"`
			LineNumber int    `json:"line_number"`
			Method     string `json:"method"`
		} `json:"details"`
	} `json:"error"`
}

func HTTPErrorTOJSON(error HTTPError) (string, *AppError) {

	json, err := json.Marshal(error)

	if err != nil {
		return "", NewAppError(err, "error trying to marshal HTTPError", -1, nil)
	}

	return string(json), nil
}
