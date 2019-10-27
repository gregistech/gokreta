package gokreta

import (
	"fmt"
	"testing"
)

func TestGetTimetableByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable by date!")
		return
	}
	accessToken := authDetails.AccessToken
	lessons, err := GetTimetableByDate(instituteCode, accessToken, "null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable by date!")
		return
	}
	fmt.Println("TestGetTimetableByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}

func TestGetTimetable(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable!")
		return
	}
	accessToken := authDetails.AccessToken
	lessons, err := GetTimetable(instituteCode, accessToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable!")
		return
	}
	fmt.Println("TestGetTimetable passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}
