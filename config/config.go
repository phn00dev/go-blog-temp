package config

type Config struct {
	App      App
	Server   Server
	Database Database
}

type App struct {
	AppName string
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Username string
	Password string
	Host     string
	Port     string
	DbName   string
}
