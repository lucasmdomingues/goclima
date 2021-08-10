package goclima

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (s *service) GetWeatherByCityID(id int64) (*Weather, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "weather", "locale", fmt.Sprint(id), "current"))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("token", s.Token)

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

	var weather *Weather

	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
