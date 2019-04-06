package clients

import (
	"io/ioutil"
	"net/http"
)

type Client struct {
	client *http.Client
}

func (cl *Client) RestClientGet(url string) ([]byte, error) {

	var resp *http.Response
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return []byte{}, err
	} else {
		if cl.client == nil {
			cl.client = &http.Client{}
		}
		if resp, err = cl.client.Do(req); err != nil {
			return []byte{}, err
		}
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
