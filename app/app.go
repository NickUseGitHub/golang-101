package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	// Register for connection to postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/NickUseGitHub/golang-101/configs"
	"github.com/gorilla/mux"
)

// App struct type
type App struct {
	Router  *mux.Router
	configs configs.Configs
}

// Initialize app
func (a *App) Initialize(configs configs.Configs) {
	a.configs = configs.Initialize()

	dbApp := &DBApp{}
	err := dbApp.InitializeDB(a.configs)
	if err != nil {
		fmt.Println(err)
		panic("[x]:: failed to connect database")
	}
	defer dbApp.DB.Close()

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.assignMethodGet("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})
}

// Get http method for APP
func (a *App) assignMethodGet(path string, handleFn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, handleFn).Methods("GET")
}

// Run with it's own port
func (a *App) Run() {
	port := a.configs.GetPort()
	fmt.Printf(fmt.Sprintf("App listen on port: %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)
}
