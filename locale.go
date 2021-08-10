package goclima

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func (s *service) SearchLocaleByID(id int64) (*Locale, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "locale", "city", fmt.Sprint(id)))
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

	var locale *Locale

	err = json.Unmarshal(body, &locale)
	if err != nil {
		return nil, err
	}

	return locale, nil
}

func (s *service) SearchLocaleByFilter(filter map[string]string) ([]Locale, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "locale", "city"))
	if err != nil {
		return nil, err
	}

	q := url.Query()

	for k, v := range filter {
		q.Add(k, v)
	}

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

	if resp.StatusCode != http.StatusOK {
		var climaTempoError *ClimaTempoError

		err = json.Unmarshal(body, &climaTempoError)
		if err != nil {
			return nil, err
		}

		if climaTempoError.Error {
			return nil, errors.New(climaTempoError.Detail)
		}
	}

	var locales []Locale

	err = json.Unmarshal(body, &locales)
	if err != nil {
		return nil, err
	}

	return locales, nil
}

func (s *service) SearchLocaleByLatitudeAndLongitude(lat, long float64) (*Locale, error) {
	url, err := s.BaseURL.Parse(path.Join("api", "v1", "locale", "city"))
	if err != nil {
		return nil, err
	}

	q := url.Query()
	q.Add("token", s.Token)
	q.Add("latitude", fmt.Sprint(lat))
	q.Add("longitude", fmt.Sprint(long))

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

	var locale *Locale

	err = json.Unmarshal(body, &locale)
	if err != nil {
		return nil, err
	}

	return locale, nil
}
