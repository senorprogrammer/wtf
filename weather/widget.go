package weather

import (
	"fmt"

	owm "github.com/briandowns/openweathermap"
	"github.com/rivo/tview"
)

func Widget() tview.Primitive {
	data := Fetch()

	widget := tview.NewTextView()
	widget.SetBorder(true)
	widget.SetDynamicColors(true)
	widget.SetTitle(fmt.Sprintf(" %s Weather - %s ", icon(data), data.Name))

	str := fmt.Sprintf("\n")
	for _, weather := range data.Weather {
		str = str + fmt.Sprintf("%16s\n\n", weather.Description)
	}

	str = str + fmt.Sprintf("%10s: %4.1f° C\n\n", "Current", data.Main.Temp)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "High", data.Main.TempMax)
	str = str + fmt.Sprintf("%10s: %4.1f° C\n", "Low", data.Main.TempMin)

	fmt.Fprintf(widget, " %s ", str)

	return widget
}

// icon returns an emoji for the current weather
func icon(data *owm.CurrentWeatherData) string {
	var icon string

	switch data.Weather[0].Description {
	case "clear":
		icon = "☀️"
	case "cloudy":
		icon = "⛅️"
	case "heavy rain":
		icon = "🌧"
	case "light rain":
		icon = "🌧"
	case "moderate rain":
		icon = "🌧"
	case "overcast":
		icon = "🌥"
	case "overcast clouds":
		icon = "🌥"
	case "partly cloudy":
		icon = "🌤"
	case "snow":
		icon = "❄️"
	case "sunny":
		icon = "☀️"
	default:
		icon = "🌈"
	}

	return icon
}
