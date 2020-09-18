package goclima

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const URL = "http://apiadvisor.climatempo.com.br/api/v1"

func GetLocaleByID(token string, id int64) (*Locale, error) {
	url := fmt.Sprintf("%s/locale/city/%d?token=%s", URL, id, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newClimaTempoError(body)
	}

	var locale *Locale

	err = json.Unmarshal(body, &locale)
	if err != nil {
		return nil, err
	}

	return locale, nil
}

func GetLocaleByNameState(token, name, state string) ([]*Locale, error) {
	var url string

	name = strings.Replace(name, " ", "+", len(name))

	if name == "" {
		url = fmt.Sprintf("%s/locale/city?state=%s&token=%s", URL, state, token)
	} else if state == "" {
		url = fmt.Sprintf("%s/locale/city?name=%s&token=%s", URL, name, token)
	} else {
		url = fmt.Sprintf("%s/locale/city?name=%s&state=%s&token=%s", URL, name, state, token)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
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
	url := fmt.Sprintf("%s/weather/locale/%d/current?token=%s", URL, id, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
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
	url := fmt.Sprintf("%s/climate/rain/locale/%d?token=%s", URL, id, token)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, newClimaTempoError(body)
	}

	var climate *Climate

	err = json.Unmarshal(body, &climate)
	if err != nil {
		return nil, err
	}

	return climate, nil
}

func newClimaTempoError(errBytes []byte) error {
	var climaTempoError *ClimaTempoError

	err := json.Unmarshal(errBytes, &climaTempoError)
	if err != nil {
		return err
	}

	return fmt.Errorf("%s", climaTempoError.Detail)
}
