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
	"eosgo-client/errors"
)

type Authorization struct {
	Account    string `json:"actor"`
	Permission string `json:"permission"`
}

type Action struct {
	Account string `json:"account"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	//Type          	string   				`json:"type"`
	Recipients    []string               `json:"recipients"`
	Authorization []Authorization        `json:"authorization"`
	Data          string                 `json:"data"`
	Args          map[string]interface{} `json:"args"`
}

type Message struct {
	Code          string          `json:"code"`
	Type          string          `json:"type"`
	Authorization []Authorization `json:"authorization"`
	Data          string          `json:"data"`
}

type Transaction struct {
	RefBlockNum             int                      `json:"ref_block_num"`
	RefBlockPrefix          int64                    `json:"ref_block_prefix"`
	Region                  int                      `json:"region"`
	Expiration              string                   `json:"expiration"`
	Scope                   []string                 `json:"scope"`
	ReadScope               []string                 `json:"read_scope"`
	Actions                 []Action                 `json:"actions"`
	Signatures              []string                 `json:"signatures"`
	Authorizations          []Authorization          `json:"authorization"`
	Status                  string                   `json:"status"`
	MaxKcpuUsage            int                      `json:"max_kcpu_usage"`
	MaxNetUsageWords        int                      `json:"max_net_usage_words"`
	DelaySec                int                      `json:"delay_sec"`
	ID                      string                   `json:"id"`
	Compression             string                   `json:"compression"`
	ContextFreeData         []map[string]interface{} `json:"context_free_data"`
	PackedBandwidthWords    int                      `json:"packed_bandwidth_words"`
	ContextFreeCPUBandwidth int                      `json:"context_free_cpu_bandwidth"`
	ContextFreeActions      []Action                 `json:"context_free_actions"`
}

func TransactionToJSON(obj *Transaction) (string, *errors.AppError) {

	bytes, err := json.Marshal(&obj)

	if err != nil {
		return "", errors.NewAppError(err, "cannot marshal transaction", -1, nil)
	}

	return string(bytes), nil
}

func JSONToTransaction(obj string) (*Transaction, *errors.AppError) {

	o := Transaction{}
	b := []byte(obj)

	err := json.Unmarshal(b, &o);
	if err != nil {
		return nil, errors.NewAppError(err, "cannot unmarshal transaction", -1, nil)
	}

	return &o, nil
}
