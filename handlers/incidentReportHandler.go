package handlers

import (
	"JedelKomek/internal/models"
	"JedelKomek/internal/services"
	"encoding/json"
	"net/http"
	"strconv"
)

type IncidentReportHandler struct {
	Service *services.IncidentReportService
}

func (h *IncidentReportHandler) CreateIncidentReport(w http.ResponseWriter, r *http.Request) {
	var report models.IncidentReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	createdReport, err := h.Service.CreateIncidentReport(r.Context(), report)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdReport)
}

func (h *IncidentReportHandler) GetAllIncidentReports(w http.ResponseWriter, r *http.Request) {
	reports, err := h.Service.GetAllIncidentReports(r.Context())
	if err != nil {
		http.Error(w, "Failed to retrieve incident reports", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reports)
}

func (h *IncidentReportHandler) GetIncidentReportById(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid incident report ID", http.StatusBadRequest)
		return
	}

	report, err := h.Service.GetIncidentReportById(r.Context(), id)
	if err != nil {
		http.Error(w, "Incident report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

func (h *IncidentReportHandler) DeleteIncidentReport(w http.ResponseWriter, r *http.Request) {
	idParam := r.URL.Query().Get(":id")
	id, err := strconv.Atoi(idParam)
	if err != nil || id <= 0 {
		http.Error(w, "Invalid incident report ID", http.StatusBadRequest)
		return
	}

	err = h.Service.DeleteIncidentReport(r.Context(), id)
	if err != nil {
		http.Error(w, "Incident report not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *IncidentReportHandler) UpdateIncidentReport(w http.ResponseWriter, r *http.Request) {
	var report models.IncidentReport
	err := json.NewDecoder(r.Body).Decode(&report)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	updatedReport, err := h.Service.UpdateIncidentReport(r.Context(), report)
	if err != nil {
		http.Error(w, "Incident report not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedReport)
}
