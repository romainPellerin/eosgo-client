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

type Code struct {
	Name     string `json:"name"`
	CodeHash string `json:"code_hash"`
	Wast     string `json:"wast"`
	Abi struct {
		Types []struct {
			NewTypeName string `json:"new_type_name"`
			Type        string `json:"type"`
		} `json:"types"`
		Structs []struct {
			Name string `json:"name"`
			Base string `json:"base"`
			Fields []struct {
				Name string `json:"name"`
				Type string `json:"type"`
			} `json:"fields"`
		} `json:"structs"`
		Actions []struct {
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"actions"`
		Tables []struct {
			Name      string   `json:"name"`
			Type      string   `json:"type"`
			IndexType string   `json:"index_type"`
			KeyNames  []string `json:"key_names"`
			KeyTypes  []string `json:"key_types"`
		} `json:"tables"`
	} `json:"abi"`
}
