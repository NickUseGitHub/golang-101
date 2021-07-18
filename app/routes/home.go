package routes

import (
	"net/http"

	"github.com/jinzhu/gorm"
)

// HomeIndex of app
func HomeIndex(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok"))
}
