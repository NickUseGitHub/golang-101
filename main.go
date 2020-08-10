package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	os.Setenv("PORT", "8080")
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})

	fmt.Printf("App listen on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
