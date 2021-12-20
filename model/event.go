package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)


type Event struct{
	ID, MaxPart int
	Title, Description, Loc string
	Date time.Time
}



func Insert(db *sqlx.DB, ev Event) error{
	_, err := db.Exec(
		"INSERT INTO events (title, description, date, location, max_participants) VALUES ($1, $2, $3, $4, $5);", 
		ev.Title, 
		ev.Description, 
		ev.Date, 
		ev.MaxPart,
	)
	if err != nil {
		return fmt.Errorf("error inserting event: %s", err.Error())
	}
	return nil
}

func GetSingle(db *sqlx.DB, id int) (Event, error){
	ev := Event{}
	err := db.QueryRowx(
		"SELECT * FROM events WHERE _id = $1",
		id,
	).StructScan(&ev)
	return ev, err
}



func GetList(db *sqlx.DB) ([]Event, error){
	evs := []Event{}
	res, err := db.Queryx("SELECT title FROM events;")
	if err !=nil && err != sql.ErrNoRows{
		return evs, fmt.Errorf("system error fetching event list: %s", err.Error())
	}

	for res.Next(){
		ev := Event{}
		err = res.StructScan(&ev)
		if err == nil{
			evs = append(evs, ev)
		}
	}
	return evs, nil

}
