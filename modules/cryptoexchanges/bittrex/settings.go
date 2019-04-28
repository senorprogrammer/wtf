package bittrex

import (
	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const configKey = "bittrex"

type colors struct {
	base struct {
		name        string
		displayName string
	}
	market struct {
		name  string
		field string
		value string
	}
}

type currency struct {
	displayName string
	market      []interface{}
}

type summary struct {
	currencies map[string]*currency
}

type Settings struct {
	colors
	common *cfg.Common
	summary
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, ymlConfig, globalConfig),
	}

	settings.colors.base.name = ymlConfig.UString("colors.base.name")
	settings.colors.base.displayName = ymlConfig.UString("colors.base.displayName")

	settings.colors.market.name = ymlConfig.UString("colors.market.name")
	settings.colors.market.field = ymlConfig.UString("colors.market.field")
	settings.colors.market.value = ymlConfig.UString("colors.market.value")

	settings.summary.currencies = make(map[string]*currency)
	for key, val := range ymlConfig.UMap("summary") {
		coercedVal := val.(map[string]interface{})

		currency := &currency{
			displayName: coercedVal["displayName"].(string),
			market:      coercedVal["market"].([]interface{}),
		}

		settings.summary.currencies[key] = currency
	}

	return &settings
}
