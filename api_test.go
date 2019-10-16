package gokreta

import (
	"fmt"
	"testing"
)

func TestGetAPIUrls(t *testing.T) {
	urls, err := GetAPIUrls()
	if err != nil || urls == nil {
		fmt.Println(err)
		t.Errorf("An error happened while trying to get API URLs!")
		return
	}
	fmt.Println("TestGetAPIUrls passed!")
	fmt.Println("Result: ")
	fmt.Println(urls)
	fmt.Println()
}
