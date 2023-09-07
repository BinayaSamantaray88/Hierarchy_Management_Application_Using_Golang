package api

import (
	"encoding/json"
	"net/http"
)

func sendErrorResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}

func sendSuccessResponse(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resp)
}
