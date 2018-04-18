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

type WeightedKey struct {
	Key    string `json:"key"`
	Weight int    `json:"weight"`
}

type Autorithy struct {
	Threshold int           `json:"threshold"`
	Keys      []WeightedKey `json:"keys"`
	Accounts  []interface{} `json:"accounts"`
}

func NewAuthority(key string, weight int) (*Autorithy) {

	wKey := WeightedKey{
		key,
		weight,
	}

	return &Autorithy{
		1,
		[]WeightedKey{wKey},
		[]interface{}{},
	}
}
