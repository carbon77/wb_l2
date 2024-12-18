package events

import (
	"encoding/json"
	"fmt"
	"time"
)

type UserId uint

type Event struct {
	UserId   UserId    `json:"user_id"`
	Title    string    `json:"title"`
	DateTime time.Time `json:"datetime"`
}

func (e Event) String() string {
	return fmt.Sprintf("Event[user_id=%d, title=%s, datetime=%v]", e.UserId, e.Title, e.DateTime)
}

func (e *Event) UnmarshalJSON(data []byte) error {
	_event := &struct {
		UserId   UserId `json:"user_id"`
		Title    string `json:"title"`
		DateTime string `json:"datetime"`
	}{}

	err := json.Unmarshal(data, &_event)
	if err != nil {
		return err
	}

	e.UserId = _event.UserId
	e.Title = _event.Title

	datetime, err := time.Parse("2006-01-02", _event.DateTime)
	if err != nil {
		return err
	}
	e.DateTime = datetime
	return nil
}
