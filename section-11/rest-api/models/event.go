package models

import (
	"example.com/rest-api/db"
	"time"
)

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description"  binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"datetime"`
	UserID      int       `json:"user_id"`
}

//var routes []Event = []Event{}

func (e Event) Save() error {
	//routes = append(routes, e)

	query := "INSERT INTO routes(name, description, location, datetime, user_id) VALUES (?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = int(id)

	return nil
}

func GetEventByID(id int) (*Event, error) {
	query := "SELECT * FROM routes WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM routes"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil
}
