package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"github.com/fatih/color"
)

var (
	apiKeyName string = "WEATHER_APIKEY"
	baseURL    string = "http://api.weatherapi.com/v1/current.json"
)

type WeatherResp struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

func colorf(v float64, c *color.Color) string {
	return c.SprintFunc()(strconv.FormatFloat(v, 'f', -1, 64))
}

func printWeather(w WeatherResp) {
	cyan := color.New(color.FgCyan)
	cyanSprint := cyan.SprintFunc()


	fmt.Fprintf(os.Stdout, "Name: %s Region: %s Country: %s\n", cyanSprint(w.Location.Name), cyanSprint(w.Location.Region), cyanSprint(w.Location.Country))
	fmt.Fprintf(os.Stdout, "Temp: %sc Feels Like: %sc\n",
		colorf(w.Current.TempC, cyan),
		colorf(w.Current.FeelslikeC, cyan))
	fmt.Fprintf(os.Stdout, "Wind speed: %s Kph\n", colorf(w.Current.WindKph,cyan))
	fmt.Fprintf(os.Stdout, "Gust speed: %s Kph\n", colorf(w.Current.GustKph,cyan))
}

func getWeather(apikey, location string) string {
	baseUrl, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return ""
	}

	query := url.Values{}
	query.Add("key", apikey)
	query.Add("q", location)
	baseUrl.RawQuery = query.Encode()

	resp, err := http.Get(baseUrl.String())
	if err != nil {
		fmt.Println("Error with request:: ", err.Error())
		return ""
	}
	defer resp.Body.Close()
	var w WeatherResp
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&w)

	if err != nil {
		panic(err)
	}
	printWeather(w)
	return ""
}

func main() {
	apiKey := os.Getenv("WEATHER_APIKEY")
	location := flag.String("location", "auto:ip", "city to get weather for")
	flag.Parse()

	if apiKey == "" {
		fmt.Println("Cannot find API key. Shutting down.")
		os.Exit(1)
	}

	fmt.Println(getWeather(apiKey, *location))

}
