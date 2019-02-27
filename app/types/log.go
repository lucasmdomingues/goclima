package types

import (
	"encoding/json"
	"log"
	"net/http"
)

type Log struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func (l *Log) Info(w http.ResponseWriter, data interface{}, message string) {

	l.Data = data
	l.Message = message

	json, err := json.Marshal(l)
	if err != nil {
		w.Write([]byte(json))
		return
	}

	w.Write([]byte(json))
}

func (l *Log) Error(w http.ResponseWriter, err error) {
	log.Printf("Error: %+v", err)
	http.Error(w, http.StatusText(500), 500)
}
