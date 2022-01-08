package main

import (
	"fmt"
	"net/url"
	"net/http"
	"os"
	"io"
)

var (
	apiKeyName string = "WEATHER_APIKEY"
	baseURL string = "http://api.weatherapi.com/v1/current.json"
)


func getWeather(apikey string) string {
	baseUrl, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println("Malformed URL: ", err.Error())
		return ""
	}

	query := url.Values{}
	//query.Add("domain", baseURL)
	query.Add("key", apikey)
	query.Add("q", "London")

	baseUrl.RawQuery = query.Encode()


	fmt.Println(baseUrl.String())


	resp, err := http.Get(baseUrl.String())

	//if err != nil {
	//	fmt.Println("Error with request:: ", err.Error())
	//	return ""
	//}
	//fmt.Println(resp)
	b, err := io.ReadAll(resp.Body)
	fmt.Println(string(b))
	return ""
}

func main() {
	var apiKey string
	apiKey = os.Getenv("WEATHER_APIKEY")

	if len(os.Args) >= 2 {
		apiKey = os.Args[1]
	}

	if apiKey == "" {
		fmt.Println("Cannot find API key. Shutting down.")
		os.Exit(1)
	}

	fmt.Println(getWeather(apiKey))

}

