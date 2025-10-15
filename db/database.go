package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dsn = "sql8802190:vvcle5YJP3@tcp(sql8.freesqldatabase.com:3306)/sql8802190?charset=utf8mb4&parseTime=True&loc=Local"
var DataBase = func() (db *gorm.DB) {
	if db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		fmt.Println("Error conexion", err)
		panic(err)
	} else {
		fmt.Println("success !")
		return db
	}
}()
