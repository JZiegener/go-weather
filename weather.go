package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	
	//fmt.Println("query URL: %s", baseURL.String())

	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error with request:: ", err.Error())
		return ""
	}
	defer resp.Body.Close()
	if(resp.StatusCode != 200) {
		fmt.Println("Error making request: ", resp.Status)

		return "";
	}


	
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
	location := flag.String("location", "auto:ip", "Location to get weather at. Can be City name, or postal code. Defaults to geo ip")
	units := flag.String("units", "m", "Units to display, m for Metric, i for Imperial")
	flag.Parse()

	if apiKey == "" {
		fmt.Println("Cannot find API key. Shutting down.")
		os.Exit(1)
	}

	useMetric := true
	switch {
	case strings.EqualFold(*units, "i") ||
		strings.EqualFold(*units, "imperial"):
		useMetric = false
	case strings.EqualFold(*units, "m") ||
		strings.EqualFold(*units, "metric"):
		useMetric = true
	default:
		fmt.Println("Invalid unit units specified: %s", units)
		fmt.Println("Please use m for metric, i for imperial")
		os.Exit(1)

	}

	fmt.Println(getWeather(apiKey, *location, useMetric))
}
