package model

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)



type Participation struct{
	FirstName string
	LastName string
	Email string
	Event_id int
}



func CreateParticipation(db *sqlx.DB, participation Participation, event Event) (error){
	partCount, err := GetParticipations(db, event.ID)
	if err != nil{
		return fmt.Errorf(
			"error retrieving amount of participants: %s",
			err.Error(),
		)
	}
	if partCount + 1 > event.Max_Participants{
		return NewLogicError("error creating participations: max participations reached")
	}
	_, err = db.Exec(
		"INSERT INTO participations (firstname, event_id, lastname, email) VALUES ($1, $2, $3, $4", 
		participation.FirstName, 
		participation.Event_id,
		participation.LastName,
		participation.Email,
	)
	if err != nil{
		return fmt.Errorf("error saving participation: %s", err.Error())
	}
	return nil
}


func GetParticipations(db *sqlx.DB, eventID int) (int, error){
	var count int
	err := db.QueryRow(
		"SELECT COUNT(event_id) FROM participations WHERE event_id = $1;",
		eventID,
	).Scan(&count)
	if err != nil {
		if err != sql.ErrNoRows{
			return 0, fmt.Errorf(
				"error retrieving participations for event with id %d: %s", 
				eventID, 
				err.Error(),
			)
		}
	}
	return count, nil
}