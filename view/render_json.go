package view

import (
	"net/http"
	"encoding/json"
)

func RenderJson(w http.ResponseWriter, res ApiResponse) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	for key, value := range res.Headers {
		w.Header().Set(key, value)
	}
	w.WriteHeader(res.Code)

	if res.Code == http.StatusNoContent {
		return
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
