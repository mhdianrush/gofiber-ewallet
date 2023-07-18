package dto

type UserData struct {
	Id       int64  `json:"id"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Username string `json:"username"`
}