package gokreta

import (
	"fmt"
	"testing"
)

func TestGetAuthDetailsByCredetinals(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	if err != nil {
		fmt.Println("An error happened while reading credetinals.txt skipping TestGetAuthDetailsByCredetinals!")
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
	instituteCode, userName, password, err := getCredetinalsFromFile()
	if err != nil {
		fmt.Println("An error happened while reading credetinals.txt skipping TestRefreshAuthDetails!")
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
