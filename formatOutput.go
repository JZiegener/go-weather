package main

import (
	"fmt"
	"os"
	"strconv"
	"github.com/fatih/color"
)
func colorf(v float64, c *color.Color) string {
	return c.SprintFunc()(strconv.FormatFloat(v, 'f', -1, 64))
}

func printWeather(w WeatherResp, useMetric bool) {
	cyan := color.New(color.FgCyan)

	var units WeatherUnits
	var report WeatherReport

	if useMetric {
		units = UnitsMetric()
		report = WeatherReportMetric(w)
	} else {
		units = UnitsImperial()
		report = WeatherReportImperial(w)
	}

	fmt.Fprintf(os.Stdout, "Weather report: %s, %s, %s\n", w.Location.Name, w.Location.Region, w.Location.Country)
	fmt.Fprintf(os.Stdout, "\t%s\n", report.Condition)
	fmt.Fprintf(os.Stdout, "\t%s (%s) %s\n",
		colorf(report.Temp, cyan),
		colorf(report.FeelsLike, cyan), units.temp)
	fmt.Fprintf(os.Stdout, "\t%s - %s %s\n", colorf(report.Wind, cyan), colorf(report.WindGust, cyan), units.speed)
	fmt.Fprintf(os.Stdout, "\t%s %s\n", colorf(report.Visiblity, cyan), units.distance)
	fmt.Fprintf(os.Stdout, "\t%s %s", colorf(report.Percip, cyan), units.volume)
}


