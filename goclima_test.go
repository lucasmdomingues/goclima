package goclima

import "testing"

const token = "8d600aefb4ff5b9ebbc08b907444d6a"

func TestGetLocaleById(t *testing.T) {

	tests := []struct {
		id   int64
		name string
	}{
		{3477, "São Paulo"},
		{5959, "Rio de Janeiro"},
	}

	for _, test := range tests {

		locale, err := GetLocaleByID(token, test.id)
		if err != nil {
			t.Error(err)
			return
		}

		if locale.Name != test.name {
			t.Errorf("Locale name is incorrect, got: %s, want: %s.", locale.Name, test.name)
		}
	}
}

func TestGetLocaleByNameState(t *testing.T) {

	tests := []struct {
		name   string
		state  string
		result string
	}{
		{"São Paulo", "", "São Paulo"},
		{"São Paulo", "SP", "São Paulo"},
	}

	for _, test := range tests {

		locales, err := GetLocaleByNameState(token, test.name, test.state)
		if err != nil {
			t.Error(err)
			return
		}

		for _, locale := range locales {
			if locale.Name != test.result {
				t.Errorf("Locale name is incorrect, got: %s, want: %s.", locale.Name, test.result)
			}
		}
	}
}

func TestGetWeather(t *testing.T) {

	tests := []struct {
		id int64
	}{
		{3477},
		{5959},
	}

	for _, test := range tests {
		_, err := GetWeather(token, test.id)
		if err != nil {
			t.Error(err)
			return
		}
	}
}

func TestGetClimate(t *testing.T) {

	tests := []struct {
		id int64
	}{
		{3477},
		{5959},
	}

	for _, test := range tests {
		_, err := GetClimate(token, test.id)
		if err != nil {
			t.Error(err)
		}
	}
}
