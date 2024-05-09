package models

import (
	"time"

	"github.com/eadenink/go-events/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time"`
	UserID      int       `json:"user_id"`
}

func (event *Event) Save() error {
	query := `
	INSERT INTO events (title, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(event.Title, event.Description, event.Location, event.DateTime, event.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	event.ID = id
	return nil
}

func GetEvents() []Event {
	return []Event{}
}
