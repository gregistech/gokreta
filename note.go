package gokreta

// Note struct contains every information about a note.
type Note struct {
	Id                int    `json:"NoteId"`
	Type              string `json:"Type"`
	Title             string `json:"Title"`
	Content           string `json:"Content"`
	SeenByTutelaryUTC string `json:"SeenByTutelaryUTC"`
	Teacher           string `json:"Teacher"`
	Date              string `json:"Date"`
	CreatingTime      string `json:"CreatingTime"`
}
