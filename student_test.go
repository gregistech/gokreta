package gokreta

import (
	"fmt"
	"testing"
)

func TestGetStudentDetails(t *testing.T) {
	t.Run("low-level", TGetStudentDetails)
	t.Run("high-level", TUserGetStudentDetails)
}

func TestGetStudentDetailsByDate(t *testing.T) {
	t.Run("low-level", TGetStudentDetailsByDate)
	t.Run("high-level", TUserGetStudentDetailsByDate)
}

func TGetStudentDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details!")
		return
	}
	student, err := GetStudentDetails(authDetails)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details!")
		return
	}
	fmt.Println("TGetStudentDetails passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TGetStudentDetailsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details by date!")
		return
	}
	student, err := GetStudentDetailsByDate(authDetails, "null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get student details by date!")
		return
	}
	fmt.Println("TGetStudentDetailsByDate passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TUserGetStudentDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User)!")
		return
	}
	student, err := user.GetStudentDetails()
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User)!")
		return
	}
	fmt.Println("TUserGetStudentDetails passed!")
	fmt.Println("Result: ")
	printStudentResults(student)
}

func TUserGetStudentDetailsByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User) by date!")
		return
	}
	student, err := user.GetStudentDetailsByDate("null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get details (from User) by date!")
		return
	}
	fmt.Println("TUserGetStudentDetailsByDate passed!")
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
