package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	port := 8080
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"Hello": "World"})
	})

	fmt.Printf("App listen on port %d", port)
	http.ListenAndServe(fmt.Sprintf(":%d", port), r)
}
