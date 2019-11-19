package gokreta

import (
	"encoding/json"
)

const API_KEY = "7856d350-1fda-45f5-822d-e1a2f3f1acf0"
const CLIENT_ID = "919e0c1c-76a2-4646-a2fb-7085bbbf3c56"

// GetApiUrls will request the current API Urls.
// Useful if you don't want to update your application
// if the API link changes.
func GetAPIUrls() (map[string]string, error) {
	urls := make(map[string]string)
	body, err := MakeRequest("GET",
		"http://kretamobile.blob.core.windows.net/configuration/ConfigurationDescriptor.json",
		nil,
		"",
	)
	if err != nil {
		return urls, err
	}
	err = json.Unmarshal(body, &urls)
	return urls, err
}
