package router

import (
	"encoding/json"
	"net/http"
)

type ResultResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"result"`
}

var (
	handlers = map[string]http.HandlerFunc{
		"/create_event": createEvent,
	}
)

func InitRouter() {
	for path, handleFunc := range handlers {
		http.HandleFunc(path, logMiddleware(handleFunc))
	}
}

func sendResultResponse(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(ResultResponse{data})
}

func sendErrorResponse(w http.ResponseWriter, message string, status int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(ErrorResponse{message})
}
