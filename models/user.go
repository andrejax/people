package models

import "database/sql"

type User struct {
	Id       string  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	GroupId  sql.NullString `json:"group_id"`
}