package rest

import (
	"encoding/json"
	commonpb "grpc-vs-rest-poc/proto"
	"grpc-vs-rest-poc/services/service-d/internal/service"
	"net/http"
)

type Handler struct {
	processor *service.Processor
}

func New(processor *service.Processor) *Handler {
	return &Handler{processor: processor}
}

func (h *Handler) Process(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var req commonpb.Request

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := h.processor.Process(ctx, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
