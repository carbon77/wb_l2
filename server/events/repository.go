package events

import (
	"sync"

	"github.com/google/uuid"
)

type EventsRepository interface {
	AddEvent(event *Event) *Event
	DeleteEvent(eventId string)
	GetEvents() []*Event
	GetEventsByUserId(userId UserId) []*Event
}

type eventsRepository struct {
	mu     *sync.Mutex
	events map[string]*Event
}

var (
	er = &eventsRepository{
		events: make(map[string]*Event),
		mu:     &sync.Mutex{},
	}
)

func Repository() EventsRepository {
	return er
}

func (ep *eventsRepository) AddEvent(event *Event) *Event {
	ep.mu.Lock()
	event.Id = uuid.New().String()
	ep.events[event.Id] = event
	ep.mu.Unlock()
	return event
}

func (ep *eventsRepository) DeleteEvent(eventId string) {
	ep.mu.Lock()
	delete(ep.events, eventId)
	ep.mu.Unlock()
}

func (ep *eventsRepository) getEventsByPredicate(predicate func(*Event) bool) []*Event {
	events := make([]*Event, 0)
	for _, event := range ep.events {
		if predicate(event) {
			events = append(events, event)
		}
	}
	return events
}

func (ep *eventsRepository) GetEvents() []*Event {
	events := make([]*Event, 0, len(ep.events))
	for _, event := range ep.events {
		events = append(events, event)
	}
	return events
}

func (ep *eventsRepository) GetEventsByUserId(userId UserId) []*Event {
	return ep.getEventsByPredicate(func(e *Event) bool { return userId == e.UserId })
}
