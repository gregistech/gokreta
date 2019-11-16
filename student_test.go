package gokreta

import (
	"fmt"
	"testing"
)

func TestGetStudentDetails(t *testing.T) {
	t.Run("low-level", TGetStudentDetails)
	t.Run("high-level", TGetDetails)
}

func TestGetStudentDetailsByDate(t *testing.T) {
	t.Run("low-level", TGetStudentDetailsByDate)
	t.Run("high-level", TGetDetailsByDate)
}

func TGetStudentDetailsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details by date!")
		return
	}
	accessToken := authDetails.AccessToken
	student, err := GetStudentDetailsByDate(instituteCode, accessToken, "null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details by date!")
		return
	}
	fmt.Println("TGetStudentDetailsByDate passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TGetStudentDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details!")
		return
	}
	accessToken := authDetails.AccessToken
	student, err := GetStudentDetails(instituteCode, accessToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details!")
		return
	}
	fmt.Println("TGetStudentDetails passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TGetDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User)!")
		return
	}
	student, err := user.GetDetails()
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User)!")
		return
	}
	fmt.Println("TGetDetails passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TGetDetailsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User) by date!")
		return
	}
	student, err := user.GetDetailsByDate("null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User) by date!")
		return
	}
	fmt.Println("TGetDetailsByDate passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func printStudentResults(student Student) {
	fmt.Println(student.Name)
	fmt.Println(student.AddressDataList)
	fmt.Println(student.InstituteCode)
	fmt.Println(student.InstituteName)
	if len(student.Evaluations) > 0 {
		fmt.Println(student.Evaluations[0])
	}
	if len(student.SubjectAverages) > 0 {
		fmt.Println(student.SubjectAverages[0])
	}
	if len(student.Absences) > 0 {
		fmt.Println(student.Absences[0])
	}
	fmt.Println(student.FormTeacher)
	fmt.Println(student.Tutelaries)
	fmt.Println()
}
