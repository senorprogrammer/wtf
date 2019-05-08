package weather

import (
	owm "github.com/briandowns/openweathermap"
	"github.com/rivo/tview"
	"github.com/wtfutil/wtf/wtf"
)

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
	wtf.HelpfulWidget
	wtf.KeyboardWidget
	wtf.TextWidget

	// APIKey   string
	Data []*owm.CurrentWeatherData
	Idx  int

	settings *Settings
}

// NewWidget creates and returns a new instance of the weather Widget
func NewWidget(app *tview.Application, pages *tview.Pages, settings *Settings) *Widget {
	widget := Widget{
		HelpfulWidget:  wtf.NewHelpfulWidget(app, pages, HelpText),
		KeyboardWidget: wtf.NewKeyboardWidget(),
		TextWidget:     wtf.NewTextWidget(app, settings.common, true),

		Idx: 0,

		settings: settings,
	}

	widget.initializeKeyboardControls()
	widget.View.SetInputCapture(widget.InputCapture)

	widget.HelpfulWidget.SetView(widget.View)

	return &widget
}

/* -------------------- Exported Functions -------------------- */

// Fetch retrieves OpenWeatherMap data from the OpenWeatherMap API.
// It takes a list of OpenWeatherMap city IDs.
// It returns a list of OpenWeatherMap CurrentWeatherData structs, one per valid city code.
func (widget *Widget) Fetch(cityIDs []int) []*owm.CurrentWeatherData {
	data := []*owm.CurrentWeatherData{}

	for _, cityID := range cityIDs {
		result, err := widget.currentWeather(cityID)
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
		widget.Data = widget.Fetch(wtf.ToInts(widget.settings.cityIDs))
	}

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
	if widget.settings.apiKey == "" {
		return false
	}

	if len(widget.settings.apiKey) != 32 {
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

func (widget *Widget) currentWeather(cityCode int) (*owm.CurrentWeatherData, error) {
	weather, err := owm.NewCurrent(
		widget.settings.tempUnit,
		widget.settings.language,
		widget.settings.apiKey,
	)
	if err != nil {
		return nil, err
	}

	err = weather.CurrentByID(cityCode)
	if err != nil {
		return nil, err
	}

	return weather, nil
}
