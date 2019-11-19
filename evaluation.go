package gokreta

// Nature struct will tell you about more about
// an evaluation's nature.
type Nature struct {
	Id          int    `json:"Id"`
	Name        string `json:"Nev"`
	Description string `json:"Leiras"`
}

// Evaluation struct contains every information about
// an evaluation.
type Evaluation struct {
	Id                int    `json:"EvaluationId"`
	Form              string `json:"Form"`
	FormName          string `json:"FormName"`
	Type              string `json:"Type"`
	TypeName          string `json:"TypeName"`
	Subject           string `json:"Subject"`
	SubjectCategory   string `json:"SubjectCategory"`
	Theme             string `json:"Theme"`
	DoesCountIntoAvg  bool   `json:"IsAtlagbaBeleszamit"`
	Mode              string `json:"Mode"`
	Weight            string `json:"Weight"`
	Value             string `json:"Value"`
	NumberValue       int    `json:"NumberValue"`
	SeenByTutelaryUTC string `json:"SeenByTutelaryUTC"`
	Teacher           string `json:"Teacher"`
	Date              string `json:"Date"`
	CreatingTime      string `json:"CreatingTime"`
	Nature            Nature `json:"Jelleg"`
	NatureName        string `json:"JellegNev"`
	ValueType         Nature `json:"ErtekFajta"`
}
