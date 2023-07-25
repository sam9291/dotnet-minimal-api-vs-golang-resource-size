package main

import (
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"time"
)

type ResponseMapper interface {
	ToResponse() Response
}

type Response struct {
	Date         string `json:"date"`
	TemperatureC int    `json:"temperatureC"`
	Summary      string `json:"summary"`
	TemperatureF int    `json:"temperatureF"`
}

type WeatherForecast struct {
	TemperatureC int
	Date         time.Time
	Summary      string
}

func (forecast WeatherForecast) TemperatureF() int {
	return int(math.Round((32 + float64(forecast.TemperatureC)/0.5556)))
}

func (forecast WeatherForecast) ToResponse() Response {
	return Response{
		TemperatureC: forecast.TemperatureC,
		Date:         forecast.Date.Format("2006-01-02"),
		Summary:      forecast.Summary,
		TemperatureF: forecast.TemperatureF(),
	}
}

var summaries = []string{
	"Freezing",
	"Bracing",
	"Chilly",
	"Cool",
	"Mild",
	"Warm",
	"Balmy",
	"Hot",
	"Sweltering",
	"Scorching",
}

const amountOfForecastsToGenerate int = 5

func handleWeatherForecast(w http.ResponseWriter, req *http.Request) {

	forecasts := make([]Response, amountOfForecastsToGenerate)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(forecasts); i++ {
		forecasts[i] = WeatherForecast{
			TemperatureC: rand.Intn(55-20) + 20,
			Date:         time.Now().AddDate(0, 0, i),
			Summary:      summaries[rand.Intn(len(summaries))],
		}.ToResponse()
	}

	respondJson(w, forecasts)
}

func respondJson(w http.ResponseWriter, data []Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func main() {

	http.HandleFunc("/weatherforecast", handleWeatherForecast)

	http.ListenAndServe(":8090", nil)
}
