package gokreta

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetAllInstitutes(t *testing.T) {
	t.Run("low-level", TGetAllInstitutes)
}

func TestGetInstituteDetails(t *testing.T) {
	t.Run("low-level", TGetInstituteDetails)
	t.Run("high-level", TUserGetInstituteDetails)
}

func TGetAllInstitutes(t *testing.T) {
	institutes, err := GetAllInstitutes()
	if err != nil || institutes == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get institutes!")
		return
	}
	fmt.Println("TGetAllInstitutes passed!")
	fmt.Println("Result: ")
	fmt.Println(institutes[2])
	fmt.Println()
}

func TGetInstituteDetails(t *testing.T) {
	institutes, err := GetAllInstitutes()
	id := institutes[0].InstituteId
	institute, err := GetInstituteDetails(id)
	if err != nil || institute.Name == "" || institutes == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get institute " + strconv.Itoa(id))
		return
	}
	fmt.Println("TGetInstituteDetails passed!")
	fmt.Println("Result: ")
	fmt.Println(institute)
	fmt.Println()
}

func TUserGetInstituteDetails(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	institute, err := user.GetInstituteDetails()
	if err != nil || institute.Name == "" {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get institute from User!")
		return
	}
	fmt.Println("TUserGetInstituteDetails passed!")
	fmt.Println("Result: ")
	fmt.Println(institute)
	fmt.Println()
}
