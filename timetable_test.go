package gokreta

import (
	"fmt"
	"testing"
)

func TestGetTimetable(t *testing.T) {
	t.Run("low-level", TGetTimetable)
	t.Run("high-level", TUserGetTimetable)
}

func TestGetTimetableByDate(t *testing.T) {
	t.Run("low-level", TGetTimetableByDate)
	t.Run("high-level", TUserGetTimetableByDate)
}

func TGetTimetable(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable!")
		return
	}
	lessons, err := GetTimetable(authDetails)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable!")
		return
	}
	fmt.Println("TGetTimetable passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}

func TGetTimetableByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable by date!")
		return
	}
	lessons, err := GetTimetableByDate(authDetails, "null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable by date!")
		return
	}
	fmt.Println("TGetTimetableByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}

func TUserGetTimetable(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable (from User)!")
		return
	}
	lessons, err := user.GetTimetable()
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable (from User)!")
		return
	}
	fmt.Println("TUserGetTimetable passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}

func TUserGetTimetableByDate(t *testing.T) {
	instituteCode, userName, password, err := getCredetinalsFromFile()
	user, err := NewUser(instituteCode, userName, password)
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable (from User) by date!")
		return
	}
	lessons, err := user.GetTimetableByDate("null", "null")
	if err != nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get the timetable (from User) by date!")
		return
	}
	fmt.Println("TUserGetTimetableByDate passed!")
	fmt.Println("Result: ")
	fmt.Println(lessons[0])
	fmt.Println()
}
