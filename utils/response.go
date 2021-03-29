package utils

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
	responseWriter http.ResponseWriter
}

func NewApiResponse(w http.ResponseWriter) *ApiResponse {
	return &ApiResponse{responseWriter: w}
}

func (apiResp ApiResponse) SuccessJsonResponse(data interface{}) {
	bytes, _ := json.Marshal(struct {
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}{true, data})

	apiResp.responseWriter.WriteHeader(http.StatusOK)
	apiResp.responseWriter.Write(bytes)
}

func (apiResp ApiResponse) ErrorJsonResponse(errorsList []error, statusCode int) {
	if statusCode == 0 {
		statusCode = http.StatusNotFound
	}

	var errors []string

	for _, error := range errorsList {
		errors = append(errors, error.Error())
	}

	bytes, _ := json.Marshal(struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	}{false, errors})

	apiResp.responseWriter.WriteHeader(statusCode)
	apiResp.responseWriter.Write(bytes)
}
