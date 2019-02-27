package handlers

import (
	"encoding/json"
	"fmt"
	"goclima/app/types"
	"net/http"
	"strconv"
	"strings"
)

func LocaleHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.Split(r.URL.Path, "/")

	switch len(path) {
	case 2:
		GetLocaleByID(w, r)
		break
	case 3:
		GetLocaleByNameState(w, r)
		break
	}
}

// Tempo no momento por ID da cidade.
func GetLocaleByID(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	if err := r.ParseForm(); err != nil {
		log.Error(w, err)
		return
	}

	if r.FormValue("id") == "" {
		log.Info(w, nil, "ID da cidade inválido.")
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

	route := fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/locale/city/%d?token=%s", id, token)

	request := types.NewRequest(w, route, "GET", nil)

	response, err := request.SendRequest()
	if err != nil {
		log.Error(w, err)
		return
	}

	var locale types.Locale

	err = json.Unmarshal(response, &locale)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(locale)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}

// Busca dados de cidades por Nome e/ou Estado.
func GetLocaleByNameState(w http.ResponseWriter, r *http.Request) {

	var log types.Log

	if err := r.ParseForm(); err != nil {
		log.Error(w, err)
		return
	}

	name := strings.Replace(r.FormValue("name"), " ", "+", len(r.FormValue("name")))
	state := r.FormValue("state")

	token := r.Header.Get("token")
	if token == "" {
		log.Info(w, nil, "Token inválido.")
		return
	}

	var route string

	if name == "" {
		route = fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/locale/city?state=%s&token=%s", state, token)
	} else if state == "" {
		route = fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/locale/city?name=%s&token=%s", name, token)
	} else if name != "" && state != "" {
		route = fmt.Sprintf("http://apiadvisor.climatempo.com.br/api/v1/locale/city?name=%s&state=%s&token=%s", name, state, token)
	}

	request := types.NewRequest(w, route, "GET", nil)

	response, err := request.SendRequest()
	if err != nil {
		log.Error(w, err)
		return
	}

	if len(response) == 2 {
		log.Info(w, nil, "Nenhum dado foi retornado.")
		return
	}

	var locales []types.Locale

	err = json.Unmarshal(response, &locales)
	if err != nil {
		log.Error(w, err)
		return
	}

	json, err := json.Marshal(locales)
	if err != nil {
		log.Error(w, err)
		return
	}

	fmt.Fprint(w, string(json))
}
