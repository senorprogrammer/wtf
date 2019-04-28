package ipinfo

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "ipinfo"

type colors struct {
	name  string
	value string
}

type Settings struct {
	colors
	common *cfg.Common
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),
	}

	settings.colors.name = ymlConfig.UString("colors.name", "red")
	settings.colors.value = ymlConfig.UString("colors.value", "white")

	return &settings
}
