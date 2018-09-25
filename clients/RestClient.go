package clients

import (
	"context"
	"google.golang.org/appengine/urlfetch"
	"io/ioutil"
	"net/http"
)

type Client struct {
}

func (cl *Client) RestClientGet(context context.Context, url string) ([]byte, error) {

	var resp *http.Response
	var err error
	client := urlfetch.Client(context)
	if resp, err = client.Get(url); err != nil {
		return []byte{}, err
	}

	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
