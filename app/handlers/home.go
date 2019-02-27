package handlers

import (
	"encoding/json"
	"fmt"
	"goclima/app/types"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	api := make(map[string]string)
	api["Title"] = "CoClima (Clima Tempo)"
	api["Version"] = "1.0"

	json, err := json.Marshal(api)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
