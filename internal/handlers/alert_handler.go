package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"JedelKomek/internal/models"
	"JedelKomek/internal/services"
)

type AlertHandler struct {
	Service *services.AlertService
}

// @Summary Create Alert
// @Tags Alerts
// @Accept json
// @Produce json
// @Param input body models.Alert true "alert"
// @Success 200 {object} models.Alert
// @Router /api/alerts [post]
func (h *AlertHandler) Create(w http.ResponseWriter, r *http.Request) {
	var a models.Alert
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	created, err := h.Service.Create(r.Context(), a)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(created)
}

func (h *AlertHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	alerts, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(alerts)
}

func (h *AlertHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	alert, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(alert)
}

func (h *AlertHandler) Update(w http.ResponseWriter, r *http.Request) {
	var a models.Alert
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if err := h.Service.Update(r.Context(), a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *AlertHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
