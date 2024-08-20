package config

import (
	"flag"

	"github.com/spf13/viper"
)

var (
	migrate bool
	port    = 50015
)

type Config struct {
	App App
	Db  Db
}

type App struct {
	Port int
}

type Db struct {
	Migrate bool
	Uri     string
	Name    string
	User    string
	Passwd  string
	Host    string
}

func LoadConfig() *Config {
	return &Config{
		App: App{
			Port: port,
		},
		Db: Db{
			Migrate: migrate,
			Name:    viper.GetString("database.name"),
			User:    viper.GetString("database.user"),
			Passwd:  viper.GetString("database.password"),
			Host:    viper.GetString("database.host"),
		},
	}
}

func InitViper(filename string) {
	flag.BoolVar(&migrate, "migrate", false, "enables database auto migration")
	flag.Parse()
	viper.SetConfigName(filename)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./conf")

	if err := viper.ReadInConfig(); err != nil {
		panic("error reading config file")
	}
}
