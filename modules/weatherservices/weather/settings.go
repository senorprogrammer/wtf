package weather

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "weather"

type colors struct {
	current string
}

type Settings struct {
	colors
	common *cfg.Common

	apiKey   string
	cityIDs  []interface{}
	language string
	tempUnit string
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),

		apiKey:   ymlConfig.UString("apiKey", os.Getenv("WTF_OWM_API_KEY")),
		cityIDs:  ymlConfig.UList("cityids"),
		language: ymlConfig.UString("language", "EN"),
		tempUnit: ymlConfig.UString("tempUnit", "C"),
	}

	settings.colors.current = ymlConfig.UString("colors.current", "green")

	return &settings
}
