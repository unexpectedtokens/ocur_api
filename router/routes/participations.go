package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unexpectedtokens/ocur_api/db"
	"github.com/unexpectedtokens/ocur_api/model"
	"github.com/unexpectedtokens/ocur_api/util"
)

func CreateParticipation(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idString := ps.ByName("id")
	var eventID int
	var err error
	if eventID, err = util.IDStringToINT(idString); err != nil {
		util.BadRequest(w, "invalid id parameter")
		return
	}
	var participation model.Participation
	err = json.NewDecoder(r.Body).Decode(&participation)
	if err != nil {
		fmt.Println(err.Error())
		util.BadRequest(w, "unable to parse body")
		return
	}
	event, err := model.GetSingle(db.DBCon, eventID)
	if err != nil {
		util.ServerError(w)
		return
	}
	err = model.CreateParticipation(db.DBCon, participation, event)
	if err != nil {
		if lerr, ok := err.(model.LogicError); ok {
			util.BadRequest(w, lerr.Error())
			return
		} else {
			util.ServerError(w)
			return
		}
	}
	w.WriteHeader(http.StatusCreated)

}

func GetParticipations(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := util.IDStringToINT(ps.ByName("id"))
	if err != nil {
		util.BadRequest(w, "unable to parse id")
	}
	participations, err := model.GetParticipations(db.DBCon, id)
	if err != nil {
		if err == sql.ErrNoRows {
			util.NotFound(w)
		} else {
			util.ServerError(w)
			fmt.Printf("error fetching participations: %s\n", err.Error())
		}
	}
	marshaled, err := json.Marshal(participations)
	if err != nil {
		util.ServerError(w)
		return
	}
	w.Write(marshaled)
}
