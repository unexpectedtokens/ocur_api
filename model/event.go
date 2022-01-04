package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)


type Event struct{
	ID int `db:"_id" json:"id"`
	Max_Participants int `json:"maxparticipants" db="max_participants"`
	Title string
	Description string 
	Location string `json:"location" `
	Date time.Time
	ParticipationCount int
}



func Insert(db *sqlx.DB, ev Event) error{
	_, err := db.Exec(
		"INSERT INTO events (title, description, date, location, max_participants) VALUES ($1, $2, $3, $4, $5);", 
		ev.Title, 
		ev.Description, 
		time.Now(), 
		"amsterdam",
		ev.Max_Participants,
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
	res, err := db.Queryx("SELECT * FROM events;")
	if err !=nil && err != sql.ErrNoRows{
		fmt.Println(err)
		return evs, fmt.Errorf("system error fetching event list: %s", err.Error())
	}

	for res.Next(){
		ev := Event{}
		err = res.StructScan(&ev)
		if err == nil{
			evs = append(evs, ev)
		}else{
			fmt.Println(err)
		}
	}
	return evs, nil

}
