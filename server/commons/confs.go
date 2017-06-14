package commons

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Config struct holds all the configurations
type Config struct {
	Debug bool
	SQL   sql
	API   api
}

type sql struct {
	Host     string
	Port     int64
	User     string
	Password string
	DBName   string
}

type api struct {
	Port int64
}

var cfg Config

// GetConfiguration returns the configuration object
func GetConfiguration() Config {
	return cfg
}

func loadConfiguration(cfg *Config) {
	conffile := "confs.toml"

	if _, err := os.Stat(conffile); err != nil {
		panic("Configuration file is missing: " + conffile)
	}

	if _, err := toml.DecodeFile(conffile, &cfg); err != nil {
		panic(err)
	}
}

func init() {
	loadConfiguration(&cfg)
}
