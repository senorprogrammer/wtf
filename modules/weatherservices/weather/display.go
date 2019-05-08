package weather

import (
	"fmt"
	"strings"

	owm "github.com/briandowns/openweathermap"
	"github.com/wtfutil/wtf/wtf"
)

func (widget *Widget) display() {
	err := ""
	if widget.apiKeyValid() == false {
		err += " Environment variable WTF_OWM_API_KEY is not set\n"
	}

	cityData := widget.currentData()
	if cityData == nil {
		err += " Weather data is unavailable: no city data\n"
	}

	if len(cityData.Weather) == 0 {
		err += " Weather data is unavailable: no weather data"
	}

	title := widget.CommonSettings.Title
	var content string
	if err != "" {
		content = err
	} else {
		title = widget.title(cityData)
		_, _, width, _ := widget.View.GetRect()
		content = widget.settings.common.SigilStr(len(widget.Data), widget.Idx, width) + "\n"
		content = content + widget.description(cityData) + "\n\n"
		content = content + widget.temperatures(cityData) + "\n"
		content = content + widget.sunInfo(cityData)
	}

	widget.Redraw(title, content, false)
}

func (widget *Widget) description(cityData *owm.CurrentWeatherData) string {
	descs := []string{}
	for _, weather := range cityData.Weather {
		descs = append(descs, fmt.Sprintf(" %s", weather.Description))
	}

	return strings.Join(descs, ",")
}

func (widget *Widget) sunInfo(cityData *owm.CurrentWeatherData) string {
	return fmt.Sprintf(
		" Rise: %s   Set: %s",
		wtf.UnixTime(int64(cityData.Sys.Sunrise)).Format("15:04 MST"),
		wtf.UnixTime(int64(cityData.Sys.Sunset)).Format("15:04 MST"),
	)
}

func (widget *Widget) temperatures(cityData *owm.CurrentWeatherData) string {
	str := fmt.Sprintf("%8s: %4.1f° %s\n", "High", cityData.Main.TempMax, widget.settings.tempUnit)

	str = str + fmt.Sprintf(
		"%8s: [%s]%4.1f° %s[white]\n",
		"Current",
		widget.settings.colors.current,
		cityData.Main.Temp,
		widget.settings.tempUnit,
	)

	str = str + fmt.Sprintf("%8s: %4.1f° %s\n", "Low", cityData.Main.TempMin, widget.settings.tempUnit)

	return str
}

func (widget *Widget) title(cityData *owm.CurrentWeatherData) string {
	str := fmt.Sprintf("%s %s", widget.emojiFor(cityData), cityData.Name)
	return widget.ContextualTitle(str)
}
