package gokreta

type Tutelary struct {
	Id          int    `json:"TutelaryId"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}

type Teacher struct {
	Id          int    `json:"TeacherId"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}
