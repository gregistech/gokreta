package gokreta

// User struct contains the auth details,
// and abstractions to lower level requests.
type User struct {
	authDetails AuthDetails
}

// NewUser will create a new User.
func NewUser(instituteCode, userName, password string) (User, error) {
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	return User{authDetails}, err
}

// GetStudentDetails will return a Student from User.
func (user User) GetStudentDetails() (Student, error) {
	return GetStudentDetails(user.authDetails)
}

// GetStudentDetailsByDate will return a Student from User restricted by the given duration.
func (user User) GetStudentDetailsByDate(fromDate, toDate string) (Student, error) {
	return GetStudentDetailsByDate(user.authDetails, fromDate, toDate)
}

// GetAllExams will return an array of Exam(s) from User.
func (user User) GetAllExams() ([]Exam, error) {
	return GetAllExams(user.authDetails)
}

// GetAllExamsByDate will return an array of Exam(s) from User restricted by the given duration.
func (user User) GetAllExamsByDate(fromDate, toDate string) ([]Exam, error) {
	return GetAllExamsByDate(user.authDetails, fromDate, toDate)
}

func (user User) GetInstituteDetails() (Institute, error) {
	institutes, err := GetAllInstitutes()
	for i := range institutes {
		e := institutes[i]
		if e.InstituteCode == user.authDetails.InstituteCode {
			return e, err
		}
	}
	return Institute{}, err
}

// GetTimetable will return an array of Lesson(s) from User.
func (user User) GetTimetable() ([]Lesson, error) {
	return GetTimetable(user.authDetails)
}

// GetTimetableByDate will return an array of Lesson(s) from User restricted by the given duration.
func (user User) GetTimetableByDate(fromDate, toDate string) ([]Lesson, error) {
	return GetTimetableByDate(user.authDetails, fromDate, toDate)
}
