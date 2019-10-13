package gokreta

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeRequest(requestType string, requestUrl string, headers http.Header, body string) ([]byte, error) {
	req, err := http.NewRequest(requestType, requestUrl, strings.NewReader(body))
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
