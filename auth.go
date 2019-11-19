package gokreta

import (
	"encoding/json"
)

// AuthDetails struct stores information that is
// required to access the API.
type AuthDetails struct {
	InstituteCode string
	AccessToken   string `json:"access_token"`
	TokenType     string `json:"token_type"`
	ExpiresIn     int    `json:"expires_in"`
	RefreshToken  string `json:"refresh_token"`
}

// GetAuthDetailsByCredetinals will return you an
// AuthDetails for further use.
func GetAuthDetailsByCredetinals(instituteCode string,
	userName string,
	password string,
) (AuthDetails, error) {
	var authDetails AuthDetails
	body := "institute_code=" + instituteCode + "&userName=" + userName + "&password=" + password + "&grant_type=password&client_id=" + CLIENT_ID
	headers := map[string]string{
		"Content-type": "application/x-www-form-urlencoded; charset=utf-8",
	}
	respBody, err := MakeRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", headers, body)
	if err != nil {
		return authDetails, err
	}
	err = json.Unmarshal(respBody, &authDetails)
	authDetails.InstituteCode = instituteCode
	return authDetails, err
}

// RefreshAuthDetails will request you a new AuthDetails
// if your current one becomes obsolete.
func RefreshAuthDetails(instituteCode string, refreshToken string) (AuthDetails, error) {
	var authDetails AuthDetails
	body := "institute_code=" + instituteCode + "&refresh_token=" + refreshToken + "&grant_type=refresh_token&client_id=" + CLIENT_ID
	headers := map[string]string{
		"Content-type": "application/x-www-form-urlencoded; charset=utf-8",
	}
	respBody, err := MakeRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", headers, body)
	if err != nil {
		return authDetails, err
	}
	err = json.Unmarshal(respBody, &authDetails)
	authDetails.InstituteCode = instituteCode
	return authDetails, err
}
