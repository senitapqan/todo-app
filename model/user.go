package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" `
	Username string `json:"username"`
	Password string `json:"password"`
}