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
	body, err := VerboseMakeRequest("GET",
		"https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute",
		http.Header{"apiKey": []string{API_KEY}},
		"",
	)
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
	body, err := VerboseMakeRequest("GET",
		"https://kretaglobalmobileapi.ekreta.hu/api/v1/Institute/"+strconv.Itoa(id),
		http.Header{"apiKey": []string{API_KEY}},
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
