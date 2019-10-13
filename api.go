package gokreta

import (
	"encoding/json"
	"net/http"
	"strconv"
)

const API_KEY = "7856d350-1fda-45f5-822d-e1a2f3f1acf0"
const CLIENT_ID = "919e0c1c-76a2-4646-a2fb-7085bbbf3c56"

type InstituteFeatures struct {
	JustificationFeatureEnabled string
}

type Institute struct {
	InstituteId      int
	InstituteCode    string
	Name             string
	Url              string
	City             string
	AdvertisingUrl   string
	FeatureToggleSet InstituteFeatures
}

type AuthDetails struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

func GetAllInstitutes() ([]Institute, error) {
	var institutes []Institute
	headers := http.Header{
		"apiKey": []string{API_KEY},
		"Accept": []string{"text/plain", "text/html", "application/json"},
	}
	body, err := MakeRequest("GET", "https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute", headers, "")
	if err != nil {
		return institutes, err
	}
	err = json.Unmarshal(body, &institutes)
	if err != nil {
		return institutes, err
	}
	return institutes, err
}

func GetInstituteDetails(id int) (Institute, error) {
	var institute Institute
	headers := http.Header{
		"apiKey": []string{API_KEY},
		"Accept": []string{"text/plain", "text/html", "application/json"},
	}
	body, err := MakeRequest("GET",
		"https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute/"+strconv.Itoa(id),
		headers,
		"",
	)
	if err != nil {
		return institute, err
	}
	err = json.Unmarshal(body, &institute)
	if err != nil {
		return institute, err
	}
	return institute, err
}

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

func GetAuthDetailsByCredetinals(instituteCode string,
	userName string,
	password string,
) (AuthDetails, error) {
	var authDetails AuthDetails
	body := "institute_code=" + instituteCode + "&userName=" + userName + "&password=" + password + "&grant_type=password&client_id=" + CLIENT_ID
	headers := http.Header{
		"Content-type": []string{"application/x-www-form-urlencoded; charset=utf-8"},
	}
	respBody, err := MakeRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", headers, body)
	if err != nil {
		return authDetails, err
	}
	err = json.Unmarshal(respBody, &authDetails)
	return authDetails, err
}

func RefreshAuthDetails(instituteCode string, refreshToken string) (AuthDetails, error) {
	var authDetails AuthDetails
	body := "institute_code=" + instituteCode + "&refresh_token=" + refreshToken + "&grant_type=refresh_token&client_id=" + CLIENT_ID
	headers := http.Header{
		"Content-type": []string{"application/x-www-form-urlencoded; charset=utf-8"},
	}
	respBody, err := MakeRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", headers, body)
	if err != nil {
		return authDetails, err
	}
	err = json.Unmarshal(respBody, &authDetails)
	return authDetails, err
}
