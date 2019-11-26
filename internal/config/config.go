package config

import "github.com/ilyakaznacheev/cleanenv"

type Application struct {
	Server   Server   `yml:"server"`
	Database Database `yml:"database"`
}

type Server struct {
	Port int    `yml:"port" env:"PORT" env-default:"8080" env-description:"server port"`
	Host string `yml:"host" env:"HOST" env-default:"localhost" env-description:"server host"`
}

type Database struct {
	MongoURI string `yml:"mongo-uri" env:"MONGO_URI" env-description:"MongoDB connection URI"`
}

// Read reads configuration from file and environment
func Read(filePath string) (*Application, error) {
	var cfg Application
	err := cleanenv.ReadConfig(filePath, &cfg)
	return &cfg, err
}
