package router

import "net/http"

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
