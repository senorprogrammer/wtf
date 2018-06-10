package weather

import (
	"os"

	"github.com/andrewzolotukhin/wtf/wtf"
	owm "github.com/briandowns/openweathermap"
	"github.com/gdamore/tcell"
	"github.com/olebedev/config"
	"github.com/rivo/tview"
)

// Config is a pointer to the global config object.
var Config *config.Config

const HelpText = `
  Keyboard commands for Weather:

    /: Show/hide this help window
    h: Previous weather location
    l: Next weather location

    arrow left:  Previous weather location
    arrow right: Next weather location
`

// Widget is the container for weather data.
type Widget struct {
	wtf.TextWidget

	app   *tview.Application
	pages *tview.Pages

	APIKey string
	Data   []*owm.CurrentWeatherData
	Idx    int
}

// NewWidget creates and returns a new instance of the weather Widget.
func NewWidget(app *tview.Application, pages *tview.Pages) *Widget {
	widget := Widget{
		TextWidget: wtf.NewTextWidget(" Weather ", "weather", true),

		app:   app,
		pages: pages,

		APIKey: os.Getenv("WTF_OWM_API_KEY"),
		Idx:    0,
	}

	widget.View.SetInputCapture(widget.keyboardIntercept)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves OpenWeatherMap data from the OpenWeatherMap API.
// It takes a list of OpenWeatherMap city IDs.
// It returns a list of OpenWeatherMap CurrentWeatherData structs, one per valid city code.
func (widget *Widget) Fetch(cityIDs []int) []*owm.CurrentWeatherData {
	data := []*owm.CurrentWeatherData{}

	for _, cityID := range cityIDs {
		result, err := widget.currentWeather(widget.APIKey, cityID)
		if err == nil {
			data = append(data, result)
		}
	}

	return data
}

// Refresh fetches new data from the OpenWeatherMap API and loads the new data into the.
// widget's view for rendering
func (widget *Widget) Refresh() {
	if widget.apiKeyValid() {
		widget.Data = widget.Fetch(wtf.ToInts(Config.UList("wtf.mods.weather.cityids", widget.defaultCityCodes())))
	}

	widget.UpdateRefreshedAt()
	widget.display()
}

// Next displays data for the next city data in the list. If the current city is the last
// city, it wraps to the first city.
func (widget *Widget) Next() {
	widget.Idx = widget.Idx + 1
	if widget.Idx == len(widget.Data) {
		widget.Idx = 0
	}

	widget.display()
}

// Prev displays data for the previous city in the list. If the previous city is the first
// city, it wraps to the last city.
func (widget *Widget) Prev() {
	widget.Idx = widget.Idx - 1
	if widget.Idx < 0 {
		widget.Idx = len(widget.Data) - 1
	}

	widget.display()
}

/* -------------------- Unexported Functions -------------------- */

func (widget *Widget) apiKeyValid() bool {
	if widget.APIKey == "" {
		return false
	}

	if len(widget.APIKey) != 32 {
		return false
	}

	return true
}

func (widget *Widget) currentData() *owm.CurrentWeatherData {
	if len(widget.Data) == 0 {
		return nil
	}

	if widget.Idx < 0 || widget.Idx >= len(widget.Data) {
		return nil
	}

	return widget.Data[widget.Idx]
}

func (widget *Widget) currentWeather(apiKey string, cityCode int) (*owm.CurrentWeatherData, error) {
	weather, err := owm.NewCurrent(Config.UString("wtf.mods.weather.tempUnit", "C"), Config.UString("wtf.mods.weather.language", "EN"), apiKey)
	if err != nil {
		return nil, err
	}

	err = weather.CurrentByID(cityCode)
	if err != nil {
		return nil, err
	}

	return weather, nil
}

func (widget *Widget) defaultCityCodes() []interface{} {
	defaultArr := []int{3370352}

	var defaults = make([]interface{}, len(defaultArr))
	for i, d := range defaultArr {
		defaults[i] = d
	}

	return defaults
}

// icon returns an emoji for the current weather
// src: https://github.com/chubin/wttr.in/blob/master/share/translations/en.txt
// Note: these only work for English weather status. Sorry about that
//
// FIXME: Move these into a configuration file so they can be changed without a compile
func (widget *Widget) icon(data *owm.CurrentWeatherData) string {
	var icon string

	if len(data.Weather) == 0 {
		return ""
	}

	switch data.Weather[0].Description {
	case "broken clouds":
		icon = "☁️"
	case "clear":
		icon = "☀️"
	case "clear sky":
		icon = "☀️"
	case "cloudy":
		icon = "⛅️"
	case "few clouds":
		icon = "🌤"
	case "fog":
		icon = "🌫"
	case "haze":
		icon = "🌫"
	case "heavy rain":
		icon = "💦"
	case "heavy snow":
		icon = "⛄️"
	case "light intensity shower rain":
		icon = "☔️"
	case "light rain":
		icon = "🌦"
	case "light shower snow":
		icon = "🌦⛄️"
	case "light snow":
		icon = "🌨"
	case "mist":
		icon = "🌬"
	case "moderate rain":
		icon = "🌧"
	case "moderate snow":
		icon = "🌨"
	case "overcast":
		icon = "🌥"
	case "overcast clouds":
		icon = "🌥"
	case "partly cloudy":
		icon = "🌤"
	case "scattered clouds":
		icon = "☁️"
	case "shower rain":
		icon = "☔️"
	case "snow":
		icon = "❄️"
	case "sunny":
		icon = "☀️"
	case "thunderstorm":
		icon = "⛈"
	default:
		icon = "💥"
	}

	return icon
}

func (widget *Widget) keyboardIntercept(event *tcell.EventKey) *tcell.EventKey {
	switch string(event.Rune()) {
	case "/":
		widget.showHelp()
		return nil
	case "h":
		widget.Prev()
		return nil
	case "l":
		widget.Next()
		return nil
	}

	switch event.Key() {
	case tcell.KeyLeft:
		widget.Prev()
		return nil
	case tcell.KeyRight:
		widget.Next()
		return nil
	default:
		return event
	}
}

func (widget *Widget) showHelp() {
	closeFunc := func() {
		widget.pages.RemovePage("help")
		widget.app.SetFocus(widget.View)
	}

	modal := wtf.NewBillboardModal(HelpText, closeFunc)

	widget.pages.AddPage("help", modal, false, true)
	widget.app.SetFocus(modal)
}
