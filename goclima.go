package goclima

import (
	"net/url"
)

type service struct {
	Token   string
	BaseURL url.URL
}

type Service interface {
	SearchLocaleByID(id int64) (*Locale, error)
	SearchLocaleByFilter(filter map[string]string) ([]Locale, error)
	GetWeatherByCityID(id int64) (*Weather, error)
	GetClimateRainByCityID(id int64, lat, long *float64) (*ClimateRain, error)
	GetClimateTemperatureByCityID(id int64, lat, long *float64) (*ClimateTemperature, error)
	GetUserRegisteredCity() (*UserCity, error)
	RegisterUserCity(locales []float64) error
}

func NewService(token string) Service {
	return &service{
		Token: token,
		BaseURL: url.URL{
			Scheme: "https",
			Host:   "apiadvisor.climatempo.com.br",
		},
	}
}
