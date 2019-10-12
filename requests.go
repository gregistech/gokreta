package gokreta

import (
	"io/ioutil"
	"net/http"
)

func MakeRequest(requestType string, requestUrl string, headers http.Header, body []byte) ([]byte, error) {
	req, err := http.NewRequest(requestType, requestUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header = headers
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
