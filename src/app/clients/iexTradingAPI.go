package clients

import (
	"io/ioutil"
	"net/http"
)

//  https://api.iextrading.com/1.0
func IexTradingAPI(url string) ([]byte, error) {
	var resp *http.Response
	var err error

	if resp, err = http.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
