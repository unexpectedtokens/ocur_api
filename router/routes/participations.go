package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/unexpectedtokens/ocur_api/db"
	"github.com/unexpectedtokens/ocur_api/model"
	"github.com/unexpectedtokens/ocur_api/util"
)




func CreateParticipation(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	idString := ps.ByName("id")
	var eventID int
	var err error
	if eventID, err = util.IDStringToINT(idString); err != nil{
		util.BadRequest(w, "invalid id parameter")
	}
	var participation model.Participation
	err = json.NewDecoder(r.Body).Decode(&participation)
	if err != nil{
		fmt.Println(err.Error())
		util.BadRequest(w, "unable to parse body")
		return
	}
	event, err := model.GetSingle(db.DBCon, eventID)
	if err != nil{
		util.ServerError(w)
		return
	}
	err = model.CreateParticipation(db.DBCon, participation, event)
	if err != nil{
		if lerr, ok := err.(model.LogicError); !ok {
			util.ServerError(w)
			return
		}else{
			util.BadRequest(w, lerr.Error())
		}
	}
	
}


	