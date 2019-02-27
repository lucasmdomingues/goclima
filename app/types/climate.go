package types

type Climate struct {
	ID        int64          `json:"id"`
	CityName  string         `json:"name"`
	CityState string         `json:"state"`
	Country   string         `json:"country"`
	Data      []*ClimateData `json:"data"`
}

type ClimateData struct {
	Date        string       `json:"date"`
	DateBR      string       `json:"date_br"`
	ClimateRain *ClimateRain `json:"climate_rain"`
}

type ClimateRain struct {
	LastYear int64 `json:"last_year"`
	Normal   int64 `json:"normal"`
	Forecast int64 `json:"forecast"`
}
