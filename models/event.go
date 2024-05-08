package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (event Event) Save() {
	events = append(events, event)
}

func GetEvents() []Event {
	return events
}
