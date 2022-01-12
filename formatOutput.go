package main

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strconv"
	"math"
)

func clamp(val, min, max float64) float64 {
	return math.Min(math.Max(val,min), max)

}

func colorlerp(val, min, max float64) color.Attribute {
	var heatColorArray[5]color.Attribute
	heatColorArray[0] =	color.FgCyan
	heatColorArray[1] =	color.FgBlue
	heatColorArray[3] =	color.FgGreen
	heatColorArray[3] =	color.FgYellow
	heatColorArray[4] =	color.FgRed
	percent := clamp(val-min, 0, max)/max
	index := math.Round(percent*4)
	return heatColorArray[int(index)]
}

func colorTemp(v float64) string {
	color := color.New(colorlerp(v, -10,32))
	return color.SprintFunc()(strconv.FormatFloat(v, 'f', -1, 64))
}

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
		colorTemp(report.Temp),
		colorTemp(report.FeelsLike), units.temp)
	fmt.Fprintf(os.Stdout, "\t%s - %s %s\n", colorf(report.Wind, cyan), colorf(report.WindGust, cyan), units.speed)
	fmt.Fprintf(os.Stdout, "\t%s %s\n", colorf(report.Visiblity, cyan), units.distance)
	fmt.Fprintf(os.Stdout, "\t%s %s", colorf(report.Percip, cyan), units.volume)
}
