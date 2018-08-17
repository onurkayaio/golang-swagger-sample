package common

import (
	"net/http"
	"encoding/json"
	"go-swagger-sample/models"
)

func ErrorResponse(
	w http.ResponseWriter,
	code int,
	message string,
	statusCode int,
	statusText string,
) {
	err := models.Error{
		Code:       code,
		Message:    message,
		StatusCode: statusCode,
		StatusText: statusText,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.StatusCode)
	json.NewEncoder(w).Encode(&err)
}
