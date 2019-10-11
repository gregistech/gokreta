package gokreta

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
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
	req, err := http.NewRequest("GET", "https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute", nil)
	if err != nil {
		return institutes, err
	}
	headers := http.Header{
		"apiKey": []string{API_KEY},
		"Accept": []string{"text/plain", "text/html", "application/json"},
	}
	req.Header = headers
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return institutes, err
	}
	defer resp.Body.Close()
	jsonResp, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(jsonResp), &institutes)
	if err != nil {
		fmt.Println(resp.Status)
		return institutes, err
	}
	return institutes, err
}

func GetInstituteDetails(id int) (Institute, error) {
	var institute Institute
	req, err := http.NewRequest("GET",
		"https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute/"+strconv.Itoa(id),
		nil,
	)
	if err != nil {
		return institute, err
	}
	headers := http.Header{
		"apiKey": []string{API_KEY},
		"Accept": []string{"text/plain", "text/html", "application/json"},
	}
	req.Header = headers
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return institute, err
	}
	defer resp.Body.Close()
	jsonResp, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(jsonResp), &institute)
	if err != nil {
		fmt.Println(resp.Status)
		return institute, err
	}
	return institute, err
}

func GetAPIUrls() (map[string]string, error) {
	urls := make(map[string]string)
	req, err := http.NewRequest("GET",
		"http://kretamobile.blob.core.windows.net/configuration/ConfigurationDescriptor.json",
		nil,
	)
	if err != nil {
		return urls, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return urls, err
	}
	defer resp.Body.Close()
	jsonResp, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(jsonResp), &urls)
	if err != nil {
		return urls, err
	}
	return urls, err
}

func GetAuthDetailsByCredetinals(instituteCode string,
	userName string,
	password string,
) (AuthDetails, error) {
	var authDetails AuthDetails
	body := strings.NewReader("institute_code=" + instituteCode + "&userName=" + userName + "&password=" + password + "&grant_type=password&client_id=" + CLIENT_ID)
	req, err := http.NewRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", body)
	if err != nil {
		return authDetails, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return authDetails, err
	}
	defer resp.Body.Close()
	jsonResp, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonResp, &authDetails)
	if err != nil {
		return authDetails, err
	}
	return authDetails, err
}

func RefreshAuthDetails(instituteCode string, refreshToken string) (AuthDetails, error) {
	var authDetails AuthDetails
	body := strings.NewReader("institute_code=" + instituteCode + "&refresh_token=" + refreshToken + "&grant_type=refresh_token&client_id=" + CLIENT_ID)
	req, err := http.NewRequest("POST", "https://"+instituteCode+".e-kreta.hu/idp/api/v1/Token", body)
	if err != nil {
		return authDetails, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=utf-8")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return authDetails, err
	}
	defer resp.Body.Close()
	jsonResp, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(jsonResp, &authDetails)
	if err != nil {
		return authDetails, err
	}
	return authDetails, err
}
