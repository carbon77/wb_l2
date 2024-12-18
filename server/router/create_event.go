package router

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"ru/zakat/server/events"
)

func createEvent(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		b, err := io.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Failed to read body"))
			log.Printf("ERROR: bad request: %s\n", err)
			return
		}

		var event events.Event
		if err = json.Unmarshal(b, &event); err != nil {
			sendErrorResponse(w, err.Error(), http.StatusBadRequest)
			log.Printf("ERROR: bad request: %s\n", err)
			return
		}

		events.Repository().AddEvent(&event)
		sendResultResponse(w, event)
	}
}
