package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"first_service/models"
	"first_service/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func New(srv *service.Service) *Handler {
	return &Handler{
		service: srv,
	}
}

type HealthResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type Notebook struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Rev         string `json:"_rev,omitempty"`
}

func (h *Handler) MountHandlers(r *mux.Router) {
	r.HandleFunc("/health", h.Health).Methods("GET")
	r.HandleFunc("/notebooks", h.Create).Methods("POST")
	r.HandleFunc("/notebooks/{notebook_id}", h.Get).Methods("GET")
	r.HandleFunc("/notebooks/{notebook_id}", h.Delete).Methods("DELETE")
	// r.HandleFunc("/notebooks/{notebook_id}", h.Update).Methods("PUT")
	// r.HandleFunc("/notebooks-list", h.List).Methods("GET")
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(HealthResponse{Status: "ok", Message: "Everthing is cool here..."}); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
	}
}

// Create
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	ctx := context.TODO()

	w.Header().Set("Content-Type", "application/json")

	var input models.CreateNotebookInputDTO
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, fmt.Sprintf("JSON inválido: %v", err), http.StatusBadRequest)
		return
	}

	resp, err := h.service.Create(ctx, input)
	if err != nil {
		http.Error(w, fmt.Sprintf("erro ao criar notebook: %v", err), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, fmt.Sprintf("falha ao codificar a resposta: %v", err), http.StatusInternalServerError)
		return
	}
}

// Update
// func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	ctx := context.TODO()

// 	vars := mux.Vars(r)
// 	id := vars["notebook_id"]

// 	var existingNB Notebook
// 	err := h.couchdb.Get(ctx, id).ScanDoc(&existingNB)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao buscar Notebook: %v", err), http.StatusNotFound)
// 		return
// 	}

// 	rev := existingNB.Rev // Aqui pega a revisão do documento
// 	if rev == "" {
// 		http.Error(w, "Falha ao obter a rev", http.StatusInternalServerError)
// 		return
// 	}

// 	var updated Notebook
// 	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
// 		http.Error(w, "Erro ao decodificar JSON de entrada", http.StatusBadRequest)
// 		return
// 	}

// 	updated.ID = id
// 	updated.Rev = rev

// 	newRev, err := h.couchdb.Put(ctx, updated.ID, updated)
// 	if err != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao atualizar documento: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	updated.Rev = newRev
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(updated)
// }

// Get
func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vars := mux.Vars(r)
	id := vars["notebook_id"]

	nb, err := h.service.Get(ctx, id)
	if err != nil {
		http.Error(w, "falha ao buscar notebook", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(nb); err != nil {
		http.Error(w, "falha ao codificar resposta", http.StatusInternalServerError)
		return
	}
}

// Delete
func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx := context.TODO()

	vars := mux.Vars(r)
	id := vars["notebook_id"]

	if err := h.service.Delete(ctx, id); err != nil {
		http.Error(w, "falha ao deletar entidade", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":   "Sucesso ao deletar o documento",
		"documento": id,
	})
}

// List
// func (h *Handler) List(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	ctx := context.TODO()

// 	rows := h.couchdb.AllDocs(ctx, kivik.Params(map[string]interface{}{"include_docs": true}))

// 	if rows != nil {
// 		http.Error(w, fmt.Sprintf("Erro ao listar documentos: %v", rows), http.StatusInternalServerError)
// 		return
// 	}
// 	defer rows.Close()

// 	var notebooks []Notebook

// 	for rows.Next() {
// 		var nb Notebook
// 		if err := rows.ScanDoc(&nb); err != nil {
// 			http.Error(w, fmt.Sprintf("Erro ao ler documento: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		id, err := rows.ID()
// 		if err != nil {
// 			http.Error(w, fmt.Sprintf("Erro ao obter ID do documento: %v", err), http.StatusInternalServerError)
// 			return
// 		}
// 		nb.ID = id
// 		fmt.Printf("couch resp: %+v\n", nb)

// 		notebooks = append(notebooks, nb)
// 	}

// 	if err := rows.Err(); err != nil {
// 		http.Error(w, fmt.Sprintf("Erro durante iteração: %v", err), http.StatusInternalServerError)
// 		return
// 	}

// 	if err := json.NewEncoder(w).Encode(notebooks); err != nil {
// 		http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
// 		return
// 	}
// }
