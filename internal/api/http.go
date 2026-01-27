package api

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusBadRequest)
		return err
	}
	return nil
}

func Text(w http.ResponseWriter, status int, s string) error {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	_, err := w.Write([]byte(s))
	if err != nil {
		return err
	}
	return nil
}

func DecodeJSON(w http.ResponseWriter, r *http.Request, message string, data any) error {
	if message == "" {
		err := json.NewDecoder(r.Body).Decode(data)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return err
		}
	}

	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		http.Error(w, message, http.StatusBadRequest)
		return err
	}
	return nil
}
