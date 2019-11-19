package gokreta

// Tutelary struct will contaion every information about a tutelary.
type Tutelary struct {
	Id          int    `json:"TutelaryId"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}

// Teacher struct will contaion every information about a teacher.
type Teacher struct {
	Id          int    `json:"TeacherId"`
	Name        string `json:"Name"`
	Email       string `json:"Email"`
	PhoneNumber string `json:"PhoneNumber"`
}
