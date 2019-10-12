package gokreta

import (
	"bufio"
	"fmt"
	"os"
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

func TestGetAPIUrls(t *testing.T) {
	urls, err := GetAPIUrls()
	if err != nil || urls == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get API URLs!")
	}
}

func TestGetAuthDetailsByCredetinals(t *testing.T) {
	file, err := os.Open("credetinals.txt")
	if err != nil {
		fmt.Println("Skipping GetAuthDetailsByCredetinals because credetinals.txt doesn't exist!")
		return
	}
	reader := bufio.NewReader(file)
	instituteCode, isPrefix, err := reader.ReadLine()
	userName, isPrefix, err := reader.ReadLine()
	password, isPrefix, err := reader.ReadLine()
	if isPrefix {
		fmt.Println("Yeaa...")
	}
	defer file.Close()
	if err != nil {
		fmt.Println("Skipping GetAuthDetailsByCredetinals because credetinals.txt doesn't exist!")
		return
	}
	authDetails, err := GetAuthDetailsByCredetinals(string(instituteCode), string(userName), string(password))
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get login details!")
	}
	if authDetails.AccessToken == "" {
		fmt.Println(authDetails)
		t.Errorf("An error happened while trying to get login details!")
	}
}

func TestRefreshAuthDetails(t *testing.T) {
	file, err := os.Open("credetinals.txt")
	if err != nil {
		fmt.Println("Skipping RefreshAuthDetails because credetinals.txt doesn't exist!")
		return
	}
	reader := bufio.NewReader(file)
	instituteCode, isPrefix, err := reader.ReadLine()
	userName, isPrefix, err := reader.ReadLine()
	password, isPrefix, err := reader.ReadLine()
	if isPrefix {
		fmt.Println("Yeaa...")
	}
	defer file.Close()
	if err != nil {
		fmt.Println("Skipping RefreshAuthDetails because credetinals.txt doesn't exist!")
		return
	}
	authDetails, err := GetAuthDetailsByCredetinals(string(instituteCode), string(userName), string(password))
	refreshedDetails, err := RefreshAuthDetails(string(instituteCode), authDetails.RefreshToken)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to refresh login details!")
	}
	if refreshedDetails.AccessToken == "" {
		fmt.Println(refreshedDetails)
		t.Errorf("An error happened while trying to refresh login details!")
	}
}
