package gokreta

// Absence struct stores details about an absence.
type Absence struct {
	Id                     int    `json:"AbsenceId"`
	Type                   string `json:"Type"`
	TypeName               string `json:"TypeName"`
	Mode                   string `json:"Mode"`
	ModeName               string `json:"ModeName"`
	Subject                string `json:"Subject"`
	SubjectCategory        string `json:"SubjectCategory"`
	SubjectCategoryName    string `json:"SubjectCategoryName"`
	DelayTimeMinutes       int    `json:"DelayTimeMinutes"`
	Teacher                string `json:"Teacher"`
	LessonStartTime        string `json:"LessonStartTime"`
	NumberOfLessons        int    `json:"NumberOfLessons"`
	CreatingTime           string `json:"CreatingTime"`
	JustificationState     string `json:"JustificationState"`
	JustificationStateName string `json:"JustificationStateName"`
	JustificationType      string `json:"JustificationType"`
	JustificationTypeName  string `json:"JustificationTypeName"`
}
