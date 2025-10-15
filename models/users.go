package models

import (
	"gorm/db"
)

type User struct {
	Id       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Users []User

func MigrateUsers() {
	db.DataBase.AutoMigrate(User{})
}
