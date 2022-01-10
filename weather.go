package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

var (
	apiKeyName string = "WEATHER_APIKEY"
	weatherURL string = "http://api.weatherapi.com/v1/current.json"
)

func getWeather(apikey, location string, useMetric bool) string {
	baseURL, err := url.Parse(weatherURL)

	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return ""
	}

	query := url.Values{}
	query.Add("key", apikey)
	query.Add("q", location)
	baseURL.RawQuery = query.Encode()

	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error with request:: ", err.Error())
		return ""
	}
	defer resp.Body.Close()
	var w weatherResp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&w)

	if err != nil {
		panic(err)
	}
	printWeather(w, useMetric)
	return ""
}

func main() {
	apiKey := os.Getenv("WEATHER_APIKEY")
	location := flag.String("location", "auto:ip", "city to get weather for")
	units := flag.String("unit", "c", "Units to uses, c for Metric, f for Imperial")
	flag.Parse()

	if apiKey == "" {
		fmt.Println("Cannot find API key. Shutting down.")
		os.Exit(1)
	}
	useMetric := true
	if *units == "f" {
		useMetric = false
	}

	fmt.Println(getWeather(apiKey, *location, useMetric))

}
