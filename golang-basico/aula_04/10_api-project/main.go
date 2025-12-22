package main

import (
	"fmt"
	"net/http"

	"github.com/Paulinhoh/golang-basico/aula_04/10_api-project/handler"
	"github.com/Paulinhoh/golang-basico/aula_04/10_api-project/middleware"
)

func main() {
	mux := http.NewServeMux()

	// configurar middleware
	loggerMux := middleware.LoggingMiddleware(mux)

	// configurar rotas
	mux.HandleFunc("/users", handler.ListUsers)
	mux.HandleFunc("/users/", handler.GetUserByID)

	fmt.Println("Servidor iniciado na porta 8080")
	http.ListenAndServe(":8080", loggerMux)
}
