package bot

type Location struct {
	Name      string `json:"name"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	LocalTime string `json:"localtime"`
}

type Condition struct {
	Text string `json:"text"`
}

type Current struct {
	TempC        float64   `json:"temp_c"`
	TempF        float64   `json:"temp_f"`
	Condition    Condition `json:"condition"`
	WindSpeedMph float64   `json:"wind_mph"`
	WindSpeedKmh float64   `json:"wind_kph"`
	PressureMb   float64   `json:"pressure_mb"`
	Humidity     float64   `json:"humidity"`
	WindChillC   float64   `json:"windchill_c"`
}

type Astro struct {
	Sunrise   string `json:"sunrise"`
	Sunset    string `json:"sunset"`
	MoonRise  string `json:"moonrise"`
	MoonSet   string `json:"moonset"`
	MoonPhase string `json:"moon_phase"`
}

type HourlyForecast struct {
	Time          string    `json:"time"`
	TempC         float64   `json:"temp_c"`
	TempF         float64   `json:"temp_f"`
	Condition     Condition `json:"condition"`
	WindMph       float64   `json:"wind_mph"`
	WindKmh       float64   `json:"wind_kph"`
	WindDirection string    `json:"wind_dir"`
	PressureMb    float64   `json:"pressure_mb"`
	Humidity      float64   `json:"humidity"`
	FeelsLikeC    float64   `json:"feelslike_c"`
	FeelsLikeF    float64   `json:"feelslike_f"`
}

type Day struct {
	MaxTempC   float64 `json:"maxtemp_c"`
	MaxTempF   float64 `json:"maxtemp_f"`
	MaxWindMph float64 `json:"maxwind_mph"`
	MaxWindKmh float64 `json:"maxwind_kph"`
}

type ForecastDay struct {
	Date           string           `json:"date"`
	Day            Day              `json:"day"`
	Astro          Astro            `json:"astro"`
	HourlyForecast []HourlyForecast `json:"hour"`
}

type Forecast struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
	Forecast struct {
		ForecastDay []ForecastDay `json:"forecastday"`
	} `json:"forecast"`
}