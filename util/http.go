package util

import (
	"fmt"
	"net/http"
)


func BadRequest(w http.ResponseWriter, message string){
	w.Write([]byte(fmt.Sprintf("Bad Request: %s", message)))
	w.WriteHeader(http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter){
	w.WriteHeader(http.StatusInternalServerError)
}