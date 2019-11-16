package gokreta

type User struct {
	authDetails AuthDetails
}

func NewUser(instituteCode, userName, password string) (User, error) {
	authDetails, err := GetAuthDetailsByCredetinals(instituteCode, userName, password)
	return User{authDetails}, err
}

func (user User) GetStudentDetails() (Student, error) {
	return GetStudentDetails(user.authDetails.InstituteCode, user.authDetails.AccessToken)
}

func (user User) GetStudentDetailsByDate(fromDate, toDate string) (Student, error) {
	return GetStudentDetailsByDate(user.authDetails.InstituteCode, user.authDetails.AccessToken, fromDate, toDate)
}
