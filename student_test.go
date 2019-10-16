package gokreta

import (
	"fmt"
	"testing"
)

func TestGetStudentDetailsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	accessToken := authDetails.AccessToken
	student, err := GetStudentDetailsByDate(instituteCode, accessToken, "null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details by date!")
		return
	}
	fmt.Println("TestGetStudentDetailsByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(student.Name)
	fmt.Println(student.AddressDataList)
	fmt.Println(student.InstituteCode)
	fmt.Println(student.InstituteName)
	fmt.Println(student.Evaluations[0])
	fmt.Println(student.SubjectAverages[0])
	fmt.Println(student.Absences[0])
	fmt.Println(student.FormTeacher)
	fmt.Println(student.Tutelaries)
	fmt.Println()
}

func TestGetStudentDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	accessToken := authDetails.AccessToken
	student, err := GetStudentDetails(instituteCode, accessToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details!")
		return
	}
	fmt.Println("TestGetStudentDetails passed!")
	fmt.Println("Result: ")
	fmt.Println(student.Name)
	fmt.Println(student.AddressDataList)
	fmt.Println(student.InstituteCode)
	fmt.Println(student.InstituteName)
	fmt.Println(student.Evaluations[0])
	fmt.Println(student.SubjectAverages[0])
	fmt.Println(student.Absences[0])
	fmt.Println(student.FormTeacher)
	fmt.Println(student.Tutelaries)
	fmt.Println()
}
