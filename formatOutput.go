package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
)

func colorf(v float64, c *color.Color) string {
	return c.SprintFunc()(strconv.FormatFloat(v, 'f', -1, 64))
}

func printWeather(w weatherResp, useMetric bool) {
	cyan := color.New(color.FgCyan)

	var units weatherUnits
	var report weatherReport

	if useMetric {
		units = unitsMetric()
		report = weatherReportMetric(w)
	} else {
		units = unitsImperial()
		report = weatherReportImperial(w)
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
