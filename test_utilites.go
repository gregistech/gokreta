package gokreta

import (
	"bufio"
	"os"
)

func getCredetinalsFromFile() (string, string, string, error) {
	var instituteCode, userName, password string
	file, err := os.Open("credetinals.txt")
	if err != nil {
		return instituteCode, userName, password, err
	}
	reader := bufio.NewReader(file)
	instituteCodeByte, isPrefix, err := reader.ReadLine()
	userNameByte, isPrefix, err := reader.ReadLine()
	passwordByte, isPrefix, err := reader.ReadLine()
	defer file.Close()
	instituteCode, userName, password = string(instituteCodeByte), string(userNameByte), string(passwordByte)
	if isPrefix {
		return instituteCode, userName, password, err
	}
	return instituteCode, userName, password, err
}
