package types

type Weather struct {
	ID        int64        `json:"id"`
	CityName  string       `json:"name"`
	CityState string       `json:"state"`
	Country   string       `json:"country"`
	Data      *WeatherData `json:"data"`
}

type WeatherData struct {
	Temperature   int    `json:"temperature"`
	WindDirection string `json:"wind_direction"`
	WindVelocity  int    `json:"wind_velocity"`
	Humidity      int    `json:"humidity"`
	Condition     string `json:"condition"`
	Pressure      int    `json:"pressure"`
	Icon          string `json:"icon"`
	Sensation     int    `json:"sensation"`
	Date          string `json:"date"`
}
