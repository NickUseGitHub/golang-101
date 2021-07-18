package routes

import (
	"encoding/json"
	"net/http"

	"github.com/NickUseGitHub/golang-101/app/models"
	"github.com/NickUseGitHub/golang-101/app/utils"
	"github.com/jinzhu/gorm"
)

// TodoCreate route
func TodoCreate(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	todo := models.Todo{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	db.Create(&todo)

	w.WriteHeader(200)
	w.Write([]byte(utils.ToJSON(todo)))
}

// TodoGetList
func TodoGetList(db *gorm.DB, w http.ResponseWriter, _ *http.Request) {
	todos := []models.Todo{}

	db.Where("name is not null").Find(&todos)

	w.WriteHeader(200)
	w.Write([]byte(utils.ToJSON(todos)))
}
