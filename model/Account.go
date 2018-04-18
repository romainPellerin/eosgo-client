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

type Account struct {
	Name              string `json:"account_name"`
	EosBalance        string `json:"eos_balance"`
	StakedBalance     string `json:"staked_balance"`
	UnstakingBalance  string `json:"unstaking_balance"`
	LastUnstakingTime string `json:"last_unstaking_time"`
	Permissions []struct {
		Name   string `json:"name"`
		Parent string `json:"parent"`
		RequiredAuth struct {
			Threshold int `json:"threshold"`
			Keys []struct {
				Key    string `json:"key"`
				Weight int    `json:"weight"`
			} `json:"keys"`
			Accounts []interface{} `json:"accounts"`
		} `json:"required_auth"`
	} `json:"permissions"`
}
