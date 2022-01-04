package util

import (
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, message string) {
	_, err := w.Write([]byte(fmt.Sprintf("Bad Request: %s", message)))
	if err != nil {
		fmt.Println(err.Error())
	}
	w.WriteHeader(http.StatusBadRequest)
}

func ServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
