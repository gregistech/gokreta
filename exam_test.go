package gokreta

import (
	"fmt"
	"testing"
)

func TestGetAllExamsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date!")
		return
	}
	accessToken := authDetails.AccessToken
	exams, err := GetAllExamsByDate(instituteCode, accessToken, "null", "null")
	if err != nil || exams == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date!")
		return
	}
	fmt.Println("TestGetAllExamsByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}

func TestGetAllExams(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams!")
		return
	}
	accessToken := authDetails.AccessToken
	exams, err := GetAllExams(instituteCode, accessToken)
	if err != nil || exams == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams!")
		return
	}
	fmt.Println("TestGetAllExams passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}
