package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(map[string]string{"id": id, "status": "ok"})
}

func postUserHandler(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Erro ao ler o corpo da requisição", http.StatusBadRequest)
		return
	}
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		http.Error(w, "Dados invalidos", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Usuario %s com %d anos recebido!\n", user.Name, user.Age)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/users/{id}", getUserHandler).Methods(http.MethodGet)
	r.HandleFunc("/users", postUserHandler).Methods(http.MethodPost)
	http.ListenAndServe(":8080", r)
}
