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

type Region struct {
	Region int `json:"region"`
	CyclesSummary [][]struct {
		ReadLocks    []interface{} `json:"read_locks"`
		WriteLocks   []interface{} `json:"write_locks"`
		Transactions []Transaction `json:"transactions"`
	} `json:"cycles_summary"`
}

type Block struct {
	Previous              string        `json:"previous"`
	Timestamp             string        `json:"timestamp"`
	TransactionMerkleRoot string        `json:"transaction_mroot"`
	BlockMerkleRoot       string        `json:"block_mroot"`
	Producer              string        `json:"producer"`
	ProducerChanges       []interface{} `json:"producer_changes"`
	ProducerSignature     string        `json:"producer_signature"`
	NewProducers          []interface{} `json:"new_producers"`
	Cycles                []interface{} `json:"cycles"`
	ID                    string        `json:"id"`
	BlockNum              int           `json:"block_num"`
	RefBlockPrefix        int           `json:"ref_block_prefix"`
	ScheduleVersion       int           `json:"schedule_version"`
	Regions               []Region      `json:"regions"`
	InputTransactions     []Transaction `json:"input_transactions"`
}
