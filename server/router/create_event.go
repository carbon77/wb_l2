package router

import (
	"fmt"
	"net/http"
)

func createEvent(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		name := "User"
		if req.URL.Query().Has("name") {
			name = req.URL.Query().Get("name")
		}
		message := fmt.Sprintf("Hello, %s!", name)
		w.Write([]byte(message))
	}
}
