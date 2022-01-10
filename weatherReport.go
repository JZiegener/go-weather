package main

//Json weather object from weatherapi.com
type weatherResp struct {
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

// Units agnostic
type weatherReport struct {
	Condition  string
	Temp       float64
	Wind       float64
	WindGust   float64
	WindDegree int
	WindDir    string
	Pressure   float64
	Percip     float64
	Humidity   int
	Cloud      int
	FeelsLike  float64
	Visiblity  float64
	Uv         float64
}

//Convert from JSON to Metrics representation
func weatherReportMetric(w weatherResp) weatherReport {
	return weatherReport{w.Current.Condition.Text,
		w.Current.TempC,
		w.Current.WindKph,
		w.Current.GustKph,
		w.Current.WindDegree,
		w.Current.WindDir,
		w.Current.PressureMb,
		w.Current.PrecipMm,
		w.Current.Humidity,
		w.Current.Cloud,
		w.Current.FeelslikeC,
		w.Current.VisKm,
		w.Current.Uv}
}

//Convert from JSON to Imperial representation
func weatherReportImperial(w weatherResp) weatherReport {
	return weatherReport{w.Current.Condition.Text,
		w.Current.TempF,
		w.Current.WindMph,
		w.Current.GustMph,
		w.Current.WindDegree,
		w.Current.WindDir,
		w.Current.PressureIn,
		w.Current.PrecipIn,
		w.Current.Humidity,
		w.Current.Cloud,
		w.Current.FeelslikeF,
		w.Current.VisMiles,
		w.Current.Uv}
}

//unit labels to use for temp/speed/...
type weatherUnits struct {
	temp     string
	speed    string
	volume   string
	distance string
	pressure string
}

//Metric units
func unitsMetric() weatherUnits {
	return weatherUnits{"C", "KpH", "mm^3", "KM", "mB"}
}

//Imperial units
func unitsImperial() weatherUnits {
	return weatherUnits{"F", "MpH", "in^3", "Mi", "In"}
}
