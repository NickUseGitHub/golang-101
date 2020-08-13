package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NickUseGitHub/golang-101/configs"
	"github.com/gorilla/mux"
)

// App struct type
type App struct {
	Router  *mux.Router
	DB      *gorm.DB
	configs configs.Configs
}

// Initialize app
func (a *App) Initialize(configs configs.Configs) {
	a.configs = configs.Initialize()

	db, err := a.initialDB(a.configs)
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	defer db.Close()

	a.DB = db

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) initialDB(cnfs configs.Configs) (*gorm.DB, error) {
	return gorm.Open("postgres", cnfs.GetDbConfigConnection(cnfs))
}

func (a *App) setRouters() {
	a.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})
}

// Get http method for APP
func (a *App) Get(path string, handleFn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, handleFn).Methods("GET")
}

// Run with it's own port
func (a *App) Run() {
	port := a.configs.GetPort()
	fmt.Printf(fmt.Sprintf("App listen on port: %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)
}
