package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const APIKey = "your_api_key_here"

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
}

func GetWeather(city string) (*WeatherData, error) {
	resp, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + city + "&appid=" + APIKey)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func main() {
	var city string
	fmt.Print("地名を入力：")
	fmt.Scanln(&city)
	// city := "Tokyo"
	weather, err := GetWeather(city)
	if err != nil {
		log.Fatalf("Could not get weather data for %s: %v", city, err)
	}

	fmt.Printf("%s!\n", city)
	fmt.Printf("Temperature in %s: %.2f°C\n", city, weather.Main.Temp - 273.15