package gokreta

import (
	"encoding/json"
	"net/http"
)

type Exam struct {
	Id          int    `json:"Id"`
	Date        string `json:"Datum"`
	LessonCount int    `json:"Oraszam"`
	Subject     string `json:"Tantargy"`
	Teacher     string `json:"Tanar"`
	Name        string `json:"SzamonkeresMegnevezese"`
	Type        string `json:"SzamonkeresModja"`
	PostDate    string `json:"BejelentesDatuma"`
}

func GetAllExams(instituteCode string, accessToken string) ([]Exam, error) {
	var exams []Exam
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/BejelentettSzamonkeres",
		headers,
		"",
	)
	if err != nil {
		return exams, err
	}
	err = json.Unmarshal(body, &exams)
	return exams, err
}

func GetAllExamsByDate(instituteCode string, accessToken string, fromDate string, toDate string) ([]Exam, error) {
	var exams []Exam
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/BejelentettSzamonkeres?DatumTol="+fromDate+"&DatumIg="+toDate,
		headers,
		"",
	)
	if err != nil {
		return exams, err
	}
	err = json.Unmarshal(body, &exams)
	return exams, err
}
