package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/NickUseGitHub/golang-101/configs"
	"github.com/gorilla/mux"
)

// App struct type
type App struct {
	Router *mux.Router
}

// Initialize app
func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.setRouters()
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
func (a *App) Run(configs configs.Configs) {
	port := configs.GetPort()
	fmt.Printf(fmt.Sprintf("App listen on port: %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)
}
