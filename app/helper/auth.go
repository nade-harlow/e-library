package helper

var session string

func SaveSession(studentID string) {
	session = studentID
}

func Session() string {
	return session
}
