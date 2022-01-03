package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/unexpectedtokens/ocur_api/db"
	event "github.com/unexpectedtokens/ocur_api/model"
)




func CreateEvent(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	ev := event.Event{}
	err := json.NewDecoder(r.Body).Decode(&ev)
	if err != nil{
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request: unable to parse body"))
		return
	}
	err = event.Insert(db.DBCon, ev)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error: query error"))
		fmt.Println(err)
		return
	}else{
		w.WriteHeader(http.StatusCreated)
	}
}




func GetEvent(w http.ResponseWriter, r * http.Request, ps httprouter.Params){
	id := ps.ByName("id")
	idInt, err := strconv.Atoi(id)
	
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("bad request: invalid id param"))
		return
	}
	fmt.Println(id, idInt)
	ev, err := event.GetSingle(db.DBCon, idInt)
	if err != nil {
		if err == sql.ErrNoRows{
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("not found"))
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error: query error"))
		return
	}
	mev, err := json.Marshal(ev)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error: query error"))
		return
	}
	w.Write(mev)
}

func GetEvents(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	evs, err := event.GetList(db.DBCon)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error: query error"))
		return
	}
	mEvs, err := json.Marshal(evs)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internal server error: query error"))
		return
	}
	w.Write(mEvs)
}