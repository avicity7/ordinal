package structs

type Quiz struct {
	QuizID  string
	TopicID string
	Title   string
	Order   int
}

type ReturnedQuiz struct {
	QuizID    string
	TopicID   string
	Title     string
	Order     int
	Questions []ReturnedQuestion
}

type Question struct {
	QuizID     string
	QuestionID string
	Body       string
	Answer     int
}

type ReturnedQuestion struct {
	QuizID     string
	QuestionID string
	Body       string
	Answer     int
	Options    []QuestionOption
}

type QuestionOption struct {
	OptionID   string
	QuestionID string
	Body       string
	Order      int
}

type QuizVideo struct {
	QuizID string
	Uri    string
	Order  int
}

type Topic struct {
	TopicID string
	Name    string
}

type Signup struct {
	Name     string
	Email    string
	Password string
	RoleID   string
}

type Login struct {
	Email    string
	Password string
}

type User struct {
	Name     string
	Email    string
	RoleID   string
	RoleName string
	Password string
}

type ReturnedUser struct {
	Name     string
	Email    string
	RoleID   string
	RoleName string
}

type UserRole struct {
	RoleID int
	Name   string
}
