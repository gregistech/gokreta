package gokreta

import (
	"encoding/json"
	"net/http"
)

type SubjectAverage struct {
	Subject             string  `json:"Subject"`
	SubjectCategory     string  `json:"SubjectCategory"`
	SubjectCategoryName string  `json:"SubjectCategoryName"`
	Value               float64 `json:"Value"`
	ClassValue          float64 `json:"ClassValue"`
	Difference          float64 `json:"Difference"`
}

type Student struct {
	Id              int              `json:"StudentId"`
	SchoolYearId    int              `json:"StudentYearId"`
	Name            string           `json:"Name"`
	NameOfBirth     string           `json:"NameOfBirth"`
	PlaceOfBirth    string           `json:"PlaceOfBirth"`
	MothersName     string           `json:"MothersName"`
	AddressDataList []string         `json:"AddressDataList"`
	DateOfBirthUtc  string           `json:"DateOfBirthUtc"`
	InstituteName   string           `json:"InstituteName"`
	InstituteCode   string           `json:"InstituteCode"`
	Evaluations     []Evaluation     `json:"Evaluations"`
	SubjectAverages []SubjectAverage `json:"SubjectAverages"`
	Absences        []Absence        `json:"Absences"`
	Notes           []Note           `json:"Notes"`
	Lessons         string           `json:"Lessons"`
	Events          string           `json:"Events"`
	FormTeacher     Teacher          `json:"FormTeacher"`
	Tutelaries      []Tutelary       `json:"Tutelaries"`
}

func GetStudentDetailsByDate(instituteCode string,
	accessToken string,
	fromDate string,
	toDate string,
) (Student, error) {
	var student Student
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/Student?fromDate="+fromDate+"&toDate="+toDate,
		headers,
		"",
	)
	if err != nil {
		return student, err
	}
	err = json.Unmarshal(body, &student)
	return student, err
}

func GetStudentDetails(instituteCode string, accessToken string) (Student, error) {
	var student Student
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/Student",
		headers,
		"",
	)
	if err != nil {
		return student, err
	}
	err = json.Unmarshal(body, &student)
	return student, err
}
