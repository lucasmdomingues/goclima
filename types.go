package goclima

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

type Locale struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	State   string `json:"state"`
	Country string `json:"country"`
}

type Weather struct {
	ID        int64        `json:"id"`
	CityName  string       `json:"name"`
	CityState string       `json:"state"`
	Country   string       `json:"country"`
	Data      *WeatherData `json:"data"`
}

type WeatherData struct {
	Temperature   int64  `json:"temperature"`
	WindDirection string `json:"wind_direction"`
	WindVelocity  int64  `json:"wind_velocity"`
	Humidity      int64  `json:"humidity"`
	Condition     string `json:"condition"`
	Pressure      int64  `json:"pressure"`
	Icon          string `json:"icon"`
	Sensation     int64  `json:"sensation"`
	Date          string `json:"date"`
}

type ClimaTempoError struct {
	Error  bool   `json:"error"`
	Detail string `json:"detail"`
}
