package conf

import (
	"encoding/json"
	"os"
	"sync"
)

var (
	configPath     = "."
	configFileName = "config.json"
)

type Config struct {
	MariaDbDataSource MariaDbDataSource `json:"mariadbDataSource"`
}

type MariaDbDataSource struct {
	DriverName        string `json:"driverName"`
	Addr              string `json:"addr"`
	Port              string `json:"port"`
	Database          string `json:"database"`
	User              string `json:"user"`
	Password          string `json:"password"`
	EnableAutoMigrate bool   `json:"enableAutoMigrate"`
}

var config Config
var configOnce sync.Once

func GetConfig() *Config {
	configOnce.Do(func() {
		bytes, err := os.ReadFile(configPath + "/" + configFileName)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &config)
		if err != nil {
			panic(err)
		}
	})
	return &config
}
