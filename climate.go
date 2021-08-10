package goclima

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (s *service) GetClimateRainByCityID(id int64, lat, long *float64) (*ClimateRain, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "climate", "rain", "locale", fmt.Sprint(id)))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("token", s.Token)

	if lat != nil && long != nil {
		q.Add("latitude", fmt.Sprint(*lat))
		q.Add("longitude", fmt.Sprint(*long))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var climaTempoError *ClimaTempoError

	err = json.Unmarshal(body, &climaTempoError)
	if err != nil {
		return nil, err
	}

	if climaTempoError.Error {
		return nil, errors.New(climaTempoError.Detail)
	}

	var climate *ClimateRain

	err = json.Unmarshal(body, &climate)
	if err != nil {
		return nil, err
	}

	return climate, nil
}

func (s *service) GetClimateTemperatureByCityID(id int64, lat, long *float64) (*ClimateTemperature, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "climate", "temperature", "locale", fmt.Sprint(id)))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("token", s.Token)

	if lat != nil && long != nil {
		q.Add("latitude", fmt.Sprint(*lat))
		q.Add("longitude", fmt.Sprint(*long))
	}

	url.RawQuery = q.Encode()

	resp, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var climaTempoError *ClimaTempoError

	err = json.Unmarshal(body, &climaTempoError)
	if err != nil {
		return nil, err
	}

	if climaTempoError.Error {
		return nil, errors.New(climaTempoError.Detail)
	}

	var climate *ClimateTemperature

	err = json.Unmarshal(body, &climate)
	if err != nil {
		return nil, err
	}

	return climate, nil
}
