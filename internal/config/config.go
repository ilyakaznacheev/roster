package config

type Application struct {
	Server   Server   `yml:"server"`
	Database Database `yml:"database"`
}

type Server struct {
	Port    int    `yml:"port" env:"PORT" env-default:"8080" env-description:"server port"`
	Host    string `yml:"host" env:"HOST" env-default:"localhost" env-description:"server host"`
	AuthKey string `yml:"auth-key" env:"AUTH_KEY" env-description:"authentication signing key"`
}

type Database struct {
	MongoURI string `yml:"mongo-uri" env:"MONGO_URI" env-description:"MongoDB connection URI"`
}
