package gokreta

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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
