package gokreta

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func MakeRequest(requestType string, requestUrl string, headers map[string]string, body string) ([]byte, error) {
	req, err := http.NewRequest(requestType, requestUrl, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	useragent, err := GetUserAgent()
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", useragent)
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	err = CheckResponseErrors(respBody)
	return respBody, err
}

func VerboseMakeRequest(requestType string,
	requestUrl string,
	headers http.Header,
	body string,
) ([]byte, error) {
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
