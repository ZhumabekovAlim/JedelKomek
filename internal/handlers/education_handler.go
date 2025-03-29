package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"JedelKomek/internal/models"
	"JedelKomek/internal/services"
)

type EducationHandler struct {
	Service *services.EducationService
}

func (h *EducationHandler) Create(w http.ResponseWriter, r *http.Request) {
	var obj models.EducationContent
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created, err := h.Service.Create(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *EducationHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *EducationHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	item, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *EducationHandler) Update(w http.ResponseWriter, r *http.Request) {
	var obj models.EducationContent
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	updated, err := h.Service.Update(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *EducationHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
