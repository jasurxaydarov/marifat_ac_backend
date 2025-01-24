package models

type TeacherReq struct {
	Teacher_name    string `json:"teacher_name"`
	Teacher_surname string `json:"teacher_surname"`
	Teacher_number  string `json:"teacher_number"`
	Teachers_tg     string `json:"teacher_tg"`
	Teacher_bio     string `json:"bio"`
	Rating          string `json:"rating"`
}

type Teacher struct {
	Teacher_id      string `json:"teacher_id"`
	Teacher_name    string `json:"teacher_name"`
	Teacher_surname string `json:"teacher_surname"`
	Teacher_number  string `json:"teacher_number"`
	Teachers_tg     string `json:"teacher_tg"`
	Teacher_bio     string `json:"bio"`
	Rating          string `json:"rating"`
}

type CoursesReq struct {
	Title       string `json:"title"`
	Teacher_id  string `json:"teacher_id"`
	Description string `json:"description"`
}

type Course struct {
	Course_id   string `json:"course_id"`
	Title       string `json:"title"`
	Teacher_id  string `json:"teacher_id"`
	Description string `json:"description"`
}

type VideoReq struct {
	Title     string `json:"title"`
	Course_id string `json:"course_id"`
}

type Id struct {
	Id string
}

type IsExists struct {
	TableName  string
	ClomunName string
	ExpValue   string
}

type IsExistsResp struct {
	IsExists bool
	Status string
}

type Claim struct {
	UserId   string
	UserRole string
}

