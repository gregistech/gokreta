package gokreta

import (
	"encoding/json"
	"errors"
)

func CheckResponseErrors(resp []byte) error {
	jsonMap := make(map[string]string)
	err := json.Unmarshal(resp, &jsonMap)
	if val, ok := jsonMap["error"]; ok {
		return errors.New(val)
	} else {
		return err
	}
}
