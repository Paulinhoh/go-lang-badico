package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/users/")
	w.Header().Set("Conteent-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "ok"})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/users/", getUserHandler)
	http.ListenAndServe(":8080", mux)
}
