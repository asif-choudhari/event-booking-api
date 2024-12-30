package services

import (
	"event-booking-api/db"
	"event-booking-api/models"
)

func GetAllEvents() (*[]models.Event, error) {
	query := `
	SELECT * FROM events ORDER BY dateTime DESC
	`
	rows, err := db.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return &events, nil
}

func GetEvent(id int64) (*models.Event, error) {
	query := `
	SELECT * FROM events
	WHERE id = ?
	`

	row := db.Connection.QueryRow(query, id)

	var event models.Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, nil
	}
	return &event, err
}

func AddEvent(event *models.Event) (int64, error) {
	query := `
	INSERT INTO events (name, description, location, dateTime, userId)
	VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := db.Connection.Prepare(query)
	if err != nil {
		return 0, err
	}

	defer stmt.Close()
	result, err := stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.UserId)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func UpdateEvent(event *models.Event) error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, userId = ?
	WHERE id = ?
	`
	_, err := db.Connection.Exec(query, event.Name, event.Description, event.Location, event.DateTime, event.UserId, event.Id)
	if err != nil {
		return err
	}
	return nil
}

func DeleteEvent(id int64) error {
	query := `
	DELETE FROM events
	WHERE id = ?
	`

	_, err := db.Connection.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
