// AUTO-GENERATED: DO NOT EDIT

package config

import (
	"strings"

	"github.com/go-playground/validator/v10"

	"project-v/pkg/l"

	"github.com/spf13/viper"
)

// Tracing ...
type Tracing struct {
	Enabled   bool   `json:"enabled" mapstructure:"enabled"`
	TeleToken string `json:"tele_token" mapstructure:"tele_token"`
	TeleID    int64  `json:"tele_id" mapstructure:"tele_id"`
	Name      string `json:"name" mapstructure:"name"`
}

// Base ...
type Base struct {
	HTTPAddress string `json:"http_address" mapstructure:"http_address"  validate:"required"`
	Environment string `json:"environment" mapstructure:"environment"  validate:"required"`

	Tracing Tracing `json:"tracing" mapstructure:"tracing"`
}

// Load ...
func Load(ll l.Logger, cPath ...string) *Config {
	var cfg = &Config{}

	v := viper.NewWithOptions(viper.KeyDelimiter("__"))

	customConfigPath := "."
	if len(cPath) > 0 {
		customConfigPath = cPath[0]
	}

	v.SetConfigType("env")
	v.SetConfigFile(".env")
	if len(cPath) > 0 {
		v.SetConfigName(".env")
	}
	v.AddConfigPath(customConfigPath)
	v.AddConfigPath(".")
	v.AddConfigPath("/app")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "__"))
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		ll.S.Debugf("Error reading config file, %s", err)
	}

	err := v.Unmarshal(&cfg)
	if err != nil {
		ll.Fatal("Failed to unmarshal config", l.Error(err))
	}

	ll.Debug("Config loaded", l.Object("config", cfg))

	validate := validator.New()
	if err := validate.Struct(cfg); err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			ll.S.Fatalf(
				"Invalid config [%+v], tag [%+v], value [%+v]",
				err.StructNamespace(), err.Tag(), err.Value(),
			)
		}
	}

	return cfg
}
