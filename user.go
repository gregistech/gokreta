package gokreta

type User struct {
	authDetails AuthDetails
}

func NewUser(instituteCode, userName, password string) (User, error) {
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	return User{authDetails}, err
}

func (user User) GetStudentDetails() (Student, error) {
	return GetStudentDetails(user.authDetails)
}

func (user User) GetStudentDetailsByDate(fromDate, toDate string) (Student, error) {
	return GetStudentDetailsByDate(user.authDetails, fromDate, toDate)
}

func (user User) GetAllExams() ([]Exam, error) {
	return GetAllExams(user.authDetails)
}

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

func (user User) GetTimetable() ([]Lesson, error) {
	return GetTimetable(user.authDetails)
}

func (user User) GetTimetableByDate(fromDate, toDate string) ([]Lesson, error) {
	return GetTimetableByDate(user.authDetails, fromDate, toDate)
}
