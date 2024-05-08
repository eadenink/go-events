package models

import "time"

type Event struct {
	ID          int       `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

var events = []Event{}

func (event *Event) Save() {
	events = append(events, *event)
}

func GetEvents() []Event {
	return events
}
