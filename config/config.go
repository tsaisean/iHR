package config

import (
	"fmt"
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/pelletier/go-toml/v2"
	"github.com/spf13/viper"
)

type Config struct {
	AppName   string   `toml:"appName"`
	AppURL    string   `toml:"appURL"`
	Env       string   `toml:"env"`
	JWTSecret string   `toml:"jwtSecret" validate:"required"`
	Database  Database `toml:"database" validate:"required"`
	Redis     Redis    `toml:"redis" validate:"required"`
	Oauth     Oauth    `toml:"oauth" validate:"required"`
	Email     Email    `toml:"email" validate:"required"`
}

type Database struct {
	Host     string `toml:"host" validate:"required"`
	Port     int    `toml:"port" validate:"required"`
	Username string `toml:"username" validate:"required"`
	Password string `toml:"password" validate:"required"`
	DBName   string `toml:"dbname" validate:"required"`
}

type Redis struct {
	Host string `toml:"host" validate:"required"`
	Port int    `toml:"port" validate:"required"`
}

type Oauth struct {
	Google Google `toml:"google" validate:"required"`
}

type Google struct {
	ClientID     string `toml:"clientID" validate:"required"`
	ClientSecret string `toml:"clientSecret" validate:"required"`
}

type Email struct {
	SenderEmail string `toml:"senderEmail" validate:"required"`
	SendgridKey string `toml:"sendgridKey" validate:"required"`
}

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	// Add all possible config paths
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")
}

func LoadConfig() (*Config, error) {
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	return &config, nil
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
