package api

import (
	"encoding/json"
	"net/http"
)

func (a *api) getFile(w http.ResponseWriter, r *http.Request) {
	p := &debug{Message: "This is a debug message", Code: 777}

	w.Header().Set("Content-Type", "application/json")

	_ = json.NewEncoder(w).Encode(p)
}
