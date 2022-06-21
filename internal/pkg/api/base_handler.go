package api

import (
	"encoding/json"
	"net/http"
)

type BaseHandler struct {
}

type baseHandler interface {
	Success(w http.ResponseWriter, r *http.Request, data interface{})
}

func (h *BaseHandler) Success(w http.ResponseWriter, r *http.Request, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
