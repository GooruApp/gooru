package api

import (
	"encoding/json"
	"net/http"
)

type debug struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (a *api) getDebugHandler(w http.ResponseWriter, r *http.Request) {
	p := &debug{Message: "This is a debug message", Code: 777}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(p)
}
