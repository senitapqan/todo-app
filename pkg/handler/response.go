package handler

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(wr http.ResponseWriter, statusCode int, message string) {
	json.NewEncoder(wr).Encode(errorResponse{message})
}
