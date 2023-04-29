package gorm

import (
	"github.com/kelseyhightower/envconfig"
)

var mc *MainMySQLConfig

// MainMySQLConfig ... 環境変数からセット
type MainMySQLConfig struct {
	User      string `envconfig:"MYSQL_USER" default:"root"`
	Password  string `envconfig:"MYSQL_PASSWORD" default:"password"`
	Host      string `envconfig:"MYSQL_HOST" default:"mysql"`
	Port      int    `envconfig:"MYSQL_PORT" default:"3307"`
	Database  string `envconfig:"MYSQL_DB" default:"mydb"`
	DebugMode bool   `envconfig:"MYSQL_DEBUGMODE" default:"true"`
}

// LoadMainMySQLConfig ...
func LoadMainMySQLConfig() *MainMySQLConfig {
	var cnf MainMySQLConfig
	if mc != nil {
		return mc
	}
	err := envconfig.Process("", &cnf)
	if err != nil {
		panic(err)
	}
	mc = &cnf
	return mc
}

// GetMainMySQLConfig ...
func GetMainMySQLConfig() *MainMySQLConfig {
	if mc == nil {
		LoadMainMySQLConfig()
	}
	return mc
}
