package gokreta

import (
	"encoding/json"
	"net/http"
)

type Lesson struct {
	Id                       int    `json:"LessondId"`
	Type                     string `json:"CalendarOraType"`
	Count                    int    `json:"Count"`
	StartTime                string `json:"StartTime"`
	Endtime                  string `json:"EndTime"`
	Subject                  string `json:"Subject"`
	SubjectCategory          string `json:"SubjectCategory"`
	SubjectCategoryName      string `json:"SubjectCategoryName"`
	ClassRoom                string `json:"ClassRoom"`
	ClassGroup               string `json:"ClassGroup"`
	Teacher                  string `json:"Teacher"`
	DeputyTeacher            string `json:"DeputyTeacher"`
	State                    string `json:"State"`
	StateName                string `json:"StateName"`
	PresenceType             string `json:"PresenceType"`
	PresenceTypeName         string `json:"PresenceTypeName"`
	TeacherHomeworkId        int    `json:"TeacherHomeworkId"`
	IsStudentHomeworkEnabled bool   `json:"IsTanuloHaziFeladatEnabled"`
	Theme                    string `json:"Theme"`
	Homework                 string `json:"Homework"`
	//TODO: Homework                 Homework `json:"Homework"`
}

func GetTimetableByDate(instituteCode string, accessToken string, fromDate string, toDate string) ([]Lesson, error) {
	var lessons []Lesson
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/Lesson?fromDate="+fromDate+"&toDate="+toDate,
		headers,
		"",
	)
	if err != nil {
		return lessons, err
	}
	err = json.Unmarshal(body, &lessons)
	return lessons, err
}

func GetTimetable(instituteCode string, accessToken string) ([]Lesson, error) {
	var lessons []Lesson
	headers := http.Header{
		"Authorization": []string{"Bearer " + accessToken},
	}
	body, err := MakeRequest("GET",
		"https://"+instituteCode+".e-kreta.hu/mapi/api/v1/Lesson",
		headers,
		"",
	)
	if err != nil {
		return lessons, err
	}
	err = json.Unmarshal(body, &lessons)
	return lessons, err
}
