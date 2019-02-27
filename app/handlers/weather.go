package handlers

import (
	"encoding/json"
	"fmt"
	"goclima/app/types"
	"net/http"
	"strconv"
)

func WeatherHandler(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	if err := r.ParseForm(); err != nil {
		log.Error(w, err)
		return
	}

	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)
	if err != nil {
		log.Error(w, err)
		return
	}

	token := r.Header.Get("token")
	if token == "" {
		log.Info(w, nil, "Token inválido.")
		return
	}

	route := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/weather/locale/%d/current?token=%s", id, token)

	request := types.NewRequest(w, route, "GET", nil)

	response, err := request.SendRequest()
	if err != nil {
		log.Error(w, err)
		return
	}

	var weather types.Weather

	err = json.Unmarshal(response, &weather)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(weather)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
