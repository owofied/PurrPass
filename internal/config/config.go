package config

import (
	"encoding/json"
	"os"
)

type ServerConfig struct {
	Port string `json:"port"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"db_name"`
}

type Config struct {
	Server   ServerConfig   `json:"server"`
	Database DatabaseConfig `json:"database"`
}

func (c Config) New() Config {
	return Config{
		Server: ServerConfig{
			Port: "3000",
		},
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "",
			Password: "",
			DBName:   "purrpass",
		},
	}
}

var userConfigDirFunc = func() (string, error) {
	return os.UserConfigDir()
}

func InitializeConfig() {
	dir, err := userConfigDirFunc()
	if err != nil {
		panic(err)
	}

	err = os.Mkdir(dir+"/PurrPass", 0o755)
	if err != nil {
		panic(err)
	}

	prettyJSON, err := json.MarshalIndent(Config{}.New(), "", "  ")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dir+"/PurrPass/config.json", prettyJSON, 0o600)
	if err != nil {
		panic(err)
	}
}

func GetConfig() *Config {
	cfgDir, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	var conf []byte
	conf, err = os.ReadFile(cfgDir + "/PurrPass/config.json")
	if err != nil {
		panic(err)
	}

	var uMarshal *Config

	err = json.Unmarshal(conf, &uMarshal)
	if err != nil {
		panic(err)
	}

	return uMarshal
}
