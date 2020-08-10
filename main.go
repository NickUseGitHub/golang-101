package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/NickUseGitHub/golang-101/app"
)

func main() {
	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")
	appTest := &app.App{}
	appTest.Initialize()

	fmt.Printf("App listen on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), appTest.Router)
}
