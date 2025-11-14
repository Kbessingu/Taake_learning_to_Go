/*
Simple weather forecast CLI application
Includes Fahrenheit if wanted
Learning API and JSON usage with https://weatherapi.com
BYOAK (Bring Your Own API Key)
Author: Taake
*/
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		TempF     float64 `json:"temp_f"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	} `json:"current"`
	Forecast struct {
		Forecastday []struct {
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				TempF     float64 `json:"temp_f"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
				ChanceOfSnow float64 `json:"chance_of_snow"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func main() {
	res, err := http.Get("api key here")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		panic("Weather API not available")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		panic(err)
	}

	location, current, hours := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour

	fmt.Printf(
		"%s, %s: %.0f Celsius, %.0f Fahrenheit, %s\n",
		location.Name,
		location.Country,
		current.TempC,
		current.TempF,
		current.Condition.Text,
	)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)

		if date.Before(time.Now()) {
			continue
		}

		fmt.Printf("%s - ", date.Format("15:04"))

		if hour.TempC >= 15 {
			green := color.New(color.FgGreen)
			green.Printf("%.0f Celsius", hour.TempC)
		} else if hour.TempC >= 0 {
			hiblue := color.New(color.FgHiBlue)
			hiblue.Printf("%.0f Celsius", hour.TempC)
		} else if hour.TempC < 0 {
			cyan := color.New(color.FgCyan)
			cyan.Printf("%.0f Celsius", hour.TempC)
		}

		if hour.ChanceOfSnow > 0 {
			if hour.ChanceOfSnow >= 50 {
				blue := color.New(color.FgBlue)
				blue.Printf(" - %.0f%% chance of snow", hour.ChanceOfSnow)
			} else {
				fmt.Printf(" - %.0f%% chance of snow", hour.ChanceOfSnow)
			}
		} else if hour.ChanceOfRain > 0 {
			if hour.ChanceOfRain >= 40 {
				red := color.New(color.FgRed)
				red.Printf(" - %.0f%% chance of rain", hour.ChanceOfRain)
			} else {
				fmt.Printf(" - %.0f%% chance of rain", hour.ChanceOfRain)
			}
		}

		fmt.Printf(" - %s\n", hour.Condition.Text)

	}
}
