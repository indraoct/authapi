package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port   string `envconfig:"HTTP_PORT"`
	DBHost string `envconfig:"DB_HOST"`
	DBUser string `envconfig:"DB_USER"`
	DBPass string `envconfig:"DB_PASS"`
	DBName string `envconfig:"DB_NAME"`
}

func Get() Config {
	c := Config{}
	envconfig.MustProcess("", &c)

	return c
}
