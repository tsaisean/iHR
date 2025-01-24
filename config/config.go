package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

type Config struct {
	AppName   string   `toml:"appName"`
	Env       string   `toml:"env"`
	JWTSecret string   `toml:"jwtSecret" validate:"required"`
	Database  Database `toml:"database" validate:"required"`
}

type Database struct {
	Host     string `toml:"host" validate:"required"`
	Port     int    `toml:"port" validate:"required"`
	Username string `toml:"username" validate:"required"`
	Password string `toml:"password" validate:"required"`
	DBName   string `toml:"dbname" validate:"required"`
}

func LoadConfig() (*Config, error) {
	return loadConfig("config/config.toml")
}

func loadConfig(filePath string) (*Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := toml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	fmt.Println("Config loaded successfully:", config)

	return &config, nil
}
