package api

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	payload, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(payload)
}

func RespondWithError(w http.ResponseWriter, status int, err string) {
	RespondWithJSON(w, status, ErrorResponse{Error: err})
}
