package api

import (
	"encoding/json"
	"net/http"
)

// RespondWithError sends an error response with a status code and message
func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}

// HandlerErr handles generic errors
func HandlerErr(w http.ResponseWriter, r *http.Request) {
	RespondWithError(w, http.StatusBadRequest, "something went wrong")
}
