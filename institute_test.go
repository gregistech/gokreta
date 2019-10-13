package gokreta

import (
	"fmt"
	"strconv"
	"testing"
)

func TestGetAllInstitutes(t *testing.T) {
	institutes, err := GetAllInstitutes()
	if err != nil || institutes == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get institutes!")
	}
}

func TestGetInstituteDetails(t *testing.T) {
	institutes, err := GetAllInstitutes()
	id := institutes[0].InstituteId
	institute, err := GetInstituteDetails(id)
	if err != nil || institute.Name == "" || institutes == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get institute " + strconv.Itoa(id))
	}
}
