package utils

import (
	"encoding/json"
	"net/http"
)

func SendJSONResponse(w http.ResponseWriter, statusCode int, status string, response string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	jsonResponse := map[string]string{
		"status":   status,
		"response": response,
	}
	jsonBytes, err := json.Marshal(jsonResponse)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}
	w.Write(jsonBytes)

}
