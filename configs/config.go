package configs

import (
	config_helper "github.com/Some-Trash-Stuff/infra-helper/settings"
)

type AppSettings struct {
	Database struct {
		Server   string `json:"Server" env:"DB_SERVER"`
		Port     int    `json:"Port" env:"DB_PORT"`
		User     string `json:"User" env:"DB_USER"`
		Password string `json:"Password" env:"DB_PASSWORD"`
	} `json:"Database"`
}

var Configs AppSettings

// init carrega as configurações automaticamente ao importar o package
func init() {
	Configs = config_helper.Load[AppSettings]()
}
