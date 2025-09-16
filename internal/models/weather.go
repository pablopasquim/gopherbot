package models

type WeatherResponse struct {
	Name    string    `json:"name"`
	Main    Main      `json:"main"`
	Weather []Weather `json:"weather"`
}

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

type Weather struct {
	Description string `json:"description"`
}
