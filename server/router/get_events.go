package router

import (
	"net/http"
	"ru/zakat/server/events"
	"strconv"
)

func getEvents(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		userIdParam := getQueryParam(req, "user_id", "")
		var result []*events.Event

		if userIdParam == "" {
			result = events.Repository().GetEvents()
		} else {
			userId, err := strconv.Atoi(userIdParam)
			if err != nil {
				sendErrorResponse(w, "invalid user id", http.StatusBadRequest)
				return
			}
			result = events.Repository().GetEventsByUserId(events.UserId(userId))
		}

		sendResultResponse(w, result)
	}
}
