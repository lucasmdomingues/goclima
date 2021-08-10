package goclima

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"
)

func (s *service) GetUserRegisteredCity() (*UserCity, error) {
	url, err := s.BaseURL.Parse(path.Join("api-manager", "user-token", s.Token, "locales"))
	if err != nil {
		return nil, err
	}

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

	var userCity *UserCity

	err = json.Unmarshal(body, &userCity)
	if err != nil {
		return nil, err
	}

	return userCity, nil
}

func (s *service) RegisterUserCity(locales []float64) error {
	rel, err := s.BaseURL.Parse(path.Join("api-manager", "user-token", s.Token, "locales"))
	if err != nil {
		return err
	}

	q := rel.Query()
	q.Add("token", s.Token)

	rel.RawQuery = q.Encode()

	form := url.Values{}

	for _, locale := range locales {
		form.Add("localeId[]", fmt.Sprint(locale))
	}

	req, err := http.NewRequest(http.MethodPut, rel.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	var client http.Client

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var climaTempoError *ClimaTempoError

	err = json.Unmarshal(body, &climaTempoError)
	if err != nil {
		return err
	}

	if climaTempoError.Error {
		return errors.New(climaTempoError.Detail)
	}

	return nil
}
