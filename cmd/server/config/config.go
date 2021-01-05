package config

import (
	"point/internal/store/shared/db"

	"github.com/kelseyhightower/envconfig"
	"gopkg.in/yaml.v2"
)

type (
	// Config provides the system configuration.
	Config struct {
		Server   Server
		Logging  Logging
		Database db.Config
	}

	// Server provides the server configuration.
	Server struct {
		Addr  string `envconfig:"-"`
		Host  string `envconfig:"POINT_SERVER_HOST" default:"localhost:8080"`
		Port  string `envconfig:"POINT_SERVER_PORT" default:":8080"`
		Proto string `envconfig:"POINT_SERVER_PROTO" default:"http"`
	}

	// Logging provides the logging configuration.
	Logging struct {
		Debug  bool `envconfig:"POINT_LOGS_DEBUG"`
		Trace  bool `envconfig:"POINT_LOGS_TRACE"`
		Color  bool `envconfig:"POINT_LOGS_COLOR"`
		Pretty bool `envconfig:"POINT_LOGS_PRETTY"`
		Text   bool `envconfig:"POINT_LOGS_TEXT"`
	}
)

// Environ returns the settings from the environment.
func Environ() (Config, error) {
	cfg := Config{}
	err := envconfig.Process("", &cfg)
	return cfg, err
}

// String returns the configuration in string format.
func (c *Config) String() string {
	out, _ := yaml.Marshal(c)
	return string(out)
}
