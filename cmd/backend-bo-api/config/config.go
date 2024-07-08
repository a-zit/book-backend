package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	MariaURI string
	BOPort   string
}

var c Config

func InitConfig() *Config {
	c.MariaURI = os.Getenv("MARIA_DB_URI")
	c.BOPort = os.Getenv("BO_PORT")

	return &c
}
