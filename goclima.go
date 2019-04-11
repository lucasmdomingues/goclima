package goclima

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiClimaTempoPrefix = "http://apiadvisor.climatempo.com.br/api/v1"

func GetLocaleByID(token string, id int64) (*Locale, error) {

	url := fmt.Sprintf("%s/locale/city/%d?token=%s", apiClimaTempoPrefix, id, token)

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
		url = fmt.Sprintf("%s/locale/city?state=%s&token=%s", apiClimaTempoPrefix, state, token)
	} else if state == "" {
		url = fmt.Sprintf("%s/locale/city?name=%s&token=%s", apiClimaTempoPrefix, name, token)
	} else {
		url = fmt.Sprintf("%s/locale/city?name=%s&state=%s&token=%s", apiClimaTempoPrefix, name, state, token)
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

	url := fmt.Sprintf("%s/weather/locale/%d/current?token=%s", apiClimaTempoPrefix, id, token)

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

	url := fmt.Sprintf("%s/climate/rain/locale/%d?token=%s", apiClimaTempoPrefix, id, token)

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
