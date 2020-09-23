package goclima

import "testing"

func TestGetLocaleById(t *testing.T) {
	tests := []struct {
		id   int64
		name string
	}{
		{3477, "São Paulo"},
		{5959, "Rio de Janeiro"},
	}

	service := NewService("TOKEN")

	for _, test := range tests {
		locale, err := service.SearchByID(test.id)
		if err != nil {
			t.Error(err)
			return
		}

		if locale.Name != test.name {
			t.Errorf("Locale name is incorrect, got: %s, want: %s.", locale.Name, test.name)
			return
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

	service := NewService("TOKEN")

	for _, test := range tests {
		locales, err := service.SearchByNameState(test.name, test.state)
		if err != nil {
			t.Error(err)
			return
		}

		for _, locale := range locales {
			if locale.Name != test.result {
				t.Errorf("Locale name is incorrect, got: %s, want: %s.", locale.Name, test.result)
				return
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

	service := NewService("TOKEN")

	for _, test := range tests {
		_, err := service.GetWeather(test.id)
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

	service := NewService("TOKEN")

	for _, test := range tests {
		_, err := service.GetClimate(test.id)
		if err != nil {
			t.Error(err)
			return
		}
	}
}
