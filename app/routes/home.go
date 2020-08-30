package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jinzhu/gorm"
)

// HomeIndex of app
func HomeIndex(_ *gorm.DB, w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
}
