package gokreta

import (
	"fmt"
	"testing"
)

func TestGetAllExams(t *testing.T) {
	t.Run("low-level", TGetAllExams)
	t.Run("high-level", TUserGetAllExams)
}

func TestGetAllExamsByDate(t *testing.T) {
	t.Run("low-level", TGetAllExamsByDate)
	t.Run("high-level", TUserGetAllExamsByDate)
}

func TGetAllExams(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams!")
		return
	}
	exams, err := GetAllExams(authDetails)
	if err != nil || exams == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams!")
		return
	}
	fmt.Println("TGetAllExams passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}

func TGetAllExamsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date!")
		return
	}
	exams, err := GetAllExamsByDate(authDetails, "null", "null")
	if err != nil || exams == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date!")
		return
	}
	fmt.Println("TGetAllExamsByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}

func TUserGetAllExams(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams (from User)!")
		return
	}
	exams, err := user.GetAllExams()
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams (from User)!")
		return
	}
	fmt.Println("TUserGetAllExams passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}

func TUserGetAllExamsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date (from User)!")
		return
	}
	exams, err := user.GetAllExams()
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get all the exams by date(from User)!")
		return
	}
	fmt.Println("TUserGetAllExamsByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(exams)
	fmt.Println()
}
