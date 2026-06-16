package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/IvanGelium/main-service/internal/domain"
)

type Handler struct {
	svc domain.RoomService
}

func NewHandler(s domain.RoomService) *Handler {
	return &Handler{svc: s}
}

func (h *Handler) GetRooms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	rooms, err := h.svc.GetActiveRooms()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rooms)
}
