package app

import (
	"fmt"
	"net/http"

	// Register for connection to postgres
	"github.com/jinzhu/gorm"

	"github.com/NickUseGitHub/golang-101/app/routes"
	"github.com/NickUseGitHub/golang-101/configs"
	"github.com/gorilla/mux"
)

// App struct type
type App struct {
	Router  *mux.Router
	DBApp   *DBApp
	configs configs.Configs
}

// WrappedHandleFunc higher order func
type WrappedHandleFunc func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

// Initialize app
func (a *App) Initialize(configs configs.Configs) {
	configs.LoadEnvFromFile()

	a.configs = configs.Initialize()

	a.DBApp = &DBApp{}
	err := a.DBApp.InitializeDB(a.configs)
	if err != nil {
		fmt.Println(err)
		panic("[x]:: failed to connect database")
	}

	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.assignMethodPut("/todo", a.bindingHandleFunc(routes.TodoCreate))
	a.assignMethodGet("/todo", a.bindingHandleFunc(routes.TodoGetList))
	a.assignMethodGet("/", a.bindingHandleFunc(routes.HomeIndex))
}

// Get http method for APP
func (a *App) assignMethodGet(path string, handleFn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, handleFn).Methods("GET")
}

// Put http method for APP
func (a *App) assignMethodPut(path string, handleFn func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, handleFn).Methods("PUT")
}

func (a App) bindingHandleFunc(wrappedHandleFunc WrappedHandleFunc) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		wrappedHandleFunc(a.DBApp.DB, w, r)
	}
}

// Run with it's own port
func (a *App) Run() {
	port := a.configs.GetPort()
	fmt.Printf(fmt.Sprintf("App listen on port: %s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), a.Router)
}

// CloseDB database
func (a *App) CloseDB() {
	a.DBApp.DB.Close()
}
