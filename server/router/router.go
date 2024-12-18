package router

import (
	"encoding/json"
	"io"
	"net/http"
)

type ResultResponse struct {
	Result interface{} `json:"result"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type Middleware func(next http.HandlerFunc) http.HandlerFunc

var (
	handlers = map[string]http.HandlerFunc{
		"/create_event": createEvent,
		"/events":       getEvents,
	}

	middlewares = []Middleware{
		logMiddleware,
	}
)

func InitRouter() {
	for path, handleFunc := range handlers {
		handler := handleFunc
		for _, middleware := range middlewares {
			handler = middleware(handler)
		}
		http.HandleFunc(path, handler)
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

func getQueryParam(req *http.Request, paramName, defaultValue string) string {
	if req.URL.Query().Has(paramName) {
		return req.URL.Query().Get(paramName)
	}
	return defaultValue
}

func readBody(req *http.Request, obj any) error {
	b, err := io.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, obj); err != nil {
		return err
	}
	return nil
}
