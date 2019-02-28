package goclima

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const CLIMATEMPO_PREFIX = "http://apiadvisor.climatempo.com.br"

func GetLocaleByID(token string, id int64) (*Locale, error) {

	url := fmt.Sprintf("%s/api/v1/locale/city/%d?token=%s", CLIMATEMPO_PREFIX, id, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, newClimaTempoError(body)
	}

	var locale *Locale

	err = json.Unmarshal(body, &locale)
	if err != nil {
		return nil, err
	}

	return locale, nil
}

func GetLocaleByNameState(token string, name, state string) ([]*Locale, error) {

	var url string

	name = strings.Replace(name, " ", "+", len(name))

	if name == "" {
		url = fmt.Sprintf("%s/api/v1/locale/city?state=%s&token=%s", CLIMATEMPO_PREFIX, state, token)
	} else if state == "" {
		url = fmt.Sprintf("%s/api/v1/locale/city?name=%s&token=%s", CLIMATEMPO_PREFIX, name, token)
	} else {
		url = fmt.Sprintf("%s/api/v1/locale/city?name=%s&state=%s&token=%s", CLIMATEMPO_PREFIX, name, state, token)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, newClimaTempoError(body)
	}

	var locales []*Locale

	err = json.Unmarshal(body, &locales)
	if err != nil {
		return nil, err
	}

	return locales, nil
}

func GetWeather(token string, id int64) (*Weather, error) {

	url := fmt.Sprintf("%s/api/v1/weather/locale/%d/current?token=%s", CLIMATEMPO_PREFIX, id, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, newClimaTempoError(body)
	}

	var weather *Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func GetClimate(token string, id int64) (*Climate, error) {

	url := fmt.Sprintf("%s/api/v1/climate/rain/locale/%d?token=%s", CLIMATEMPO_PREFIX, id, token)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, newClimaTempoError(body)
	}

	var climate *Climate

	err = json.Unmarshal(body, &climate)
	if err != nil {
		return nil, err
	}

	return climate, nil
}

func newClimaTempoError(error []byte) error {

	var climaTempoError *ClimaTempoError

	err := json.Unmarshal(error, &climaTempoError)
	if err != nil {
		return err
	}

	return fmt.Errorf("%s", climaTempoError.Detail)
}
