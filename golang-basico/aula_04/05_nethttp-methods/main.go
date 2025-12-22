package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
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
		} else {
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		}
	},
	)

	fmt.Println("Servindo na porta 8080...")
	http.ListenAndServe(":8080", nil) // iniciando o servidor na porta 8080
}
