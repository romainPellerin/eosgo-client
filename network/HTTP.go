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
	"encoding/json"
	"net/http"
	"eosgo-client/errors"
	"strings"
	"io/ioutil"
	"net/url"
	"fmt"
)

// simple HTTP GET request taking some URL params
func Get(_url string, _params map[string]string) ([]byte, *errors.AppError) {

	baseUrl, err := url.Parse(_url)

	if err != nil {
		return nil, errors.NewAppError(err, "cannot parse url", -1, nil)
	}

	params := url.Values{}

	for k, v := range _params {

		u := &url.URL{Path: k}
		k = u.String()

		u = &url.URL{Path: v}
		v = u.String()

		params.Add(k, v)
	}

	baseUrl.RawQuery = params.Encode()

	req, _ := http.NewRequest("GET", baseUrl.String(), strings.NewReader(""))

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach nodeos API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, "nodeos API returned an error: "+errStr, int64(httpError.Code), httpError)
	}

	return body, nil
}

// simple HTTP JSON post request taking data as a map OR byte array as JSON body
// by default bytes will be overided if keyValues is passed
func Post(url string, keyValues map[string]interface{}, bytes []byte) ([]byte, *errors.AppError) {

	fmt.Println("post keyValues: ",keyValues)
	fmt.Println("post raw: "+string(bytes))

	var err error

	if keyValues != nil {
		bytes, err = json.Marshal(&keyValues)
	}

	if err != nil {
		return nil, errors.NewAppError(err, "error marshalling params", -1, nil)
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(string(bytes)))

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach nodeos API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, "nodeos API returned an error: "+errStr, int64(httpError.Code), httpError)
	}

	return body, nil
}

// simple HTTP JSON post request taking data as JSON body
func PostRawData(url string, raw string) ([]byte, *errors.AppError) {

	fmt.Println("post raw data: " + raw)

	// force quotes if not json or array object
	if !strings.HasPrefix(raw,"\"") && !strings.HasPrefix(raw,"{") && !strings.HasPrefix(raw,"[") {
		raw = "\""+raw+"\""
	}

	req, _ := http.NewRequest("POST", url, strings.NewReader(raw))

	req.Header.Add("content-type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, errors.NewAppError(err, "error trying to reach nodeos API", -1, nil)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, errors.NewAppError(err, "error reading body response", -1, nil)
	}

	httpError := errors.HTTPError{}
	json.Unmarshal(body, &httpError)

	if httpError.Code != 0 {
		errStr, _ := errors.HTTPErrorTOJSON(httpError)
		return nil, errors.NewAppError(nil, "nodeos API returned an error: "+errStr, int64(httpError.Code), httpError)
	}

	return body, nil
}
