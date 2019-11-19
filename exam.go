package gokreta

import (
	"encoding/json"
)

// Exam struct contains every information about
// an exam.
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

// GetAllExams will request you all exams.
// Returns an array of Exam(s).
func GetAllExams(authDetails AuthDetails) ([]Exam, error) {
	var exams []Exam
	headers := map[string]string{
		"Authorization": "Bearer " + authDetails.AccessToken,
	}
	body, err := MakeRequest("GET",
		"https://"+authDetails.InstituteCode+".e-kreta.hu/mapi/api/v1/BejelentettSzamonkeres",
		headers,
		"",
	)
	if err != nil {
		return exams, err
	}
	json.Unmarshal(body, &exams)
	return exams, err
}

// GetAllExamsByDate will request you all exams by date.
// Returns an array of Exam(s).
func GetAllExamsByDate(authDetails AuthDetails, fromDate, toDate string) ([]Exam, error) {
	var exams []Exam
	headers := map[string]string{
		"Authorization": "Bearer " + authDetails.AccessToken,
	}
	body, err := MakeRequest("GET",
		"https://"+authDetails.InstituteCode+".e-kreta.hu/mapi/api/v1/BejelentettSzamonkeres?DatumTol="+fromDate+"&DatumIg="+toDate,
		headers,
		"",
	)
	if err != nil {
		return exams, err
	}
	json.Unmarshal(body, &exams)
	return exams, err
}
