package utils

import (
	"encoding/json"
	"net/http"
)

func Serve(w http.ResponseWriter, p any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(p)
}
