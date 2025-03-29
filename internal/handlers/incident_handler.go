package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"JedelKomek/internal/models"
	"JedelKomek/internal/services"
)

type IncidentHandler struct {
	Service *services.IncidentService
}

func (h *IncidentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var inc models.IncidentReport
	if err := json.NewDecoder(r.Body).Decode(&inc); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	created, err := h.Service.Create(r.Context(), inc)
	if err != nil {
		http.Error(w, "Failed to create incident", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(created)
}

func (h *IncidentHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, "Error getting incidents", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *IncidentHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	inc, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(inc)
}

func (h *IncidentHandler) Update(w http.ResponseWriter, r *http.Request) {
	var inc models.IncidentReport
	if err := json.NewDecoder(r.Body).Decode(&inc); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	updated, err := h.Service.Update(r.Context(), inc)
	if err != nil {
		http.Error(w, "Failed to update", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updated)
}

func (h *IncidentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
