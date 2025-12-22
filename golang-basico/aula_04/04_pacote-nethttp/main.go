package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc(
		"/status",
		func(w http.ResponseWriter, r *http.Request) {
			// configurando o cabeçalho de resposta
			w.Header().Set("Content-Type", "application/json")

			// respondendo com JSON
			json.NewEncoder(w).Encode(map[string]bool{"status": true})
		},
	)

	fmt.Println("Servindo na porta 8080...")
	http.ListenAndServe(":8080", nil) // iniciando o servidor na porta 8080
}
