package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"JedelKomek/internal/models"
	"JedelKomek/internal/services"
)

type EmergencyHandler struct {
	Service       *services.EmergencyService
	PoliceService *services.PoliceDepartmentService
}

func (h *EmergencyHandler) Create(w http.ResponseWriter, r *http.Request) {
	var obj models.EmergencyCall
	if err := json.NewDecoder(r.Body).Decode(&obj); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	created, err := h.Service.Create(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	police, err := h.PoliceService.FindNearestPolice(obj.Latitude, obj.Longitude)
	if err != nil {
		http.Error(w, "Failed to find nearest police", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"emergency":      created,
		"police_name":    police.Name,
		"police_phone":   police.PhoneNumber,
		"police_address": police.Address,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *EmergencyHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.Service.GetAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(list)
}

func (h *EmergencyHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	item, err := h.Service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(item)
}

func (h *EmergencyHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get(":id"))
	if err := h.Service.Delete(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
