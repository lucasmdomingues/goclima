package goclima

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchLocaleByID(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able return a locale by id", func(t *testing.T) {
		_, err := service.SearchLocaleByID(3477)
		assert.Nil(t, err)
	})

	t.Run("should be  able return error passing a invalid city id", func(t *testing.T) {
		_, err := service.SearchLocaleByID(0000)
		assert.NotNil(t, err)
	})
}

func TestSearchLocaleByFilter(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able return a locale passing name", func(t *testing.T) {
		filter := map[string]string{
			"name": "São Paulo",
		}

		_, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
	})

	t.Run("should be able return a locale passing name and country", func(t *testing.T) {
		filter := map[string]string{
			"name":    "São Paulo",
			"country": "BR",
		}

		_, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
	})

	t.Run("should be able return a locale passing name and state", func(t *testing.T) {
		filter := map[string]string{
			"name":  "São Paulo",
			"state": "SP",
		}

		_, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
	})

	t.Run("should be able return a locale passing name and province", func(t *testing.T) {
		filter := map[string]string{
			"name":     "São Paulo",
			"province": "SP",
		}

		_, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
	})

	t.Run("should be able return a locale passing name,country,state and province", func(t *testing.T) {
		filter := map[string]string{
			"name":     "São Paulo",
			"country":  "BR",
			"state":    "SP",
			"province": "SP",
		}

		_, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
	})

	t.Run("should be able return a error passing a invalid locale name", func(t *testing.T) {
		filter := map[string]string{
			"name": "Onde judas perdeu as botas",
		}

		locales, err := service.SearchLocaleByFilter(filter)
		assert.Nil(t, err)
		assert.Empty(t, locales)
	})
}

func TestGetWeather(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able return current weather by city id", func(t *testing.T) {
		_, err := service.GetWeatherByCityID(3477)
		assert.Nil(t, err)
	})
}

func TestRegisterUserLocales(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able register locale to user token", func(t *testing.T) {
		err := service.RegisterUserCity([]float64{3477})
		assert.NotNil(t, err)
	})
}

func TestGetUserLocalesByToken(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able register locale to user token", func(t *testing.T) {
		city, err := service.GetUserRegisteredCity()
		assert.Nil(t, err)
		assert.NotEmpty(t, city.Locales)
	})
}

func TestGetClimateRain(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able return climate rain by city id", func(t *testing.T) {
		_, err := service.GetClimateRainByCityID(3477, nil, nil)
		assert.Nil(t, err)
	})

	lat := -23.5480
	long := -46.6360

	t.Run("should be able return climate rain by city id and latitude and longitude", func(t *testing.T) {
		_, err := service.GetClimateRainByCityID(3477, &lat, &long)
		assert.Nil(t, err)
	})
}

func TestGetClimateTemperature(t *testing.T) {
	service := NewService("1a191762d2f4def4e18438d8c93f71b5")

	t.Run("should be able return climate temperature by city id", func(t *testing.T) {
		_, err := service.GetClimateTemperatureByCityID(3477, nil, nil)
		assert.Nil(t, err)
	})

	lat := -23.5480
	long := -46.6360

	t.Run("should be able return climate temperature by city id and latitude and longitude", func(t *testing.T) {
		_, err := service.GetClimateTemperatureByCityID(3477, &lat, &long)
		assert.Nil(t, err)
	})
}
