package api

import (
	"net/http"

	"github.com/bytedance/sonic"
)

func JSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := sonic.ConfigDefault.NewEncoder(w).Encode(v)
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
