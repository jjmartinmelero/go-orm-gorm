package handlers

import (
	"encoding/json"
	"gorm/db"
	"gorm/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}

	db.DataBase.Find(&users)
	sendData(rw, users, http.StatusOK)
}


func GetUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		sendData(rw, user, http.StatusOK)
	}
}

func getUserById(r *http.Request) (models.User, *gorm.DB) {
	
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])

	user := models.User{}
	if err := db.DataBase.First(&user, userId); err.Error != nil {
		return user, err
	} 

	return user, nil
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {

	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		sendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.DataBase.Save(&user)
		sendData(rw, user, http.StatusCreated)
	}
}


func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	
	if userPrev, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		userId := userPrev.Id
		user := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			sendError(rw, http.StatusUnprocessableEntity)
		} else {
			user.Id = userId
			db.DataBase.Save(&user)
			sendData(rw, user, http.StatusOK)
		}
	}

}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {

	if user, err := getUserById(r); err != nil {
		sendError(rw, http.StatusNotFound)
	} else {
		db.DataBase.Delete(&user)
		sendData(rw, user, http.StatusOK)
	}
}
