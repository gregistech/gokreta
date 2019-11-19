package gokreta

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// InstituteFeatures struct will tell you about
// every feature an institute has enabled.
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

// GetAllInstitute will return you an array of Institute(s)
// that contain every avaiable institute. (Might contain ads!)
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

// GetInstituteDetails will return an Institute.
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
