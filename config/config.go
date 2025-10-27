package config

type Config struct {
	Host       string
	Port       string
	AppName    string
	DBHOST     string `mapstructure:"DB_HOST"`
	DBUSERNAME string `mapstructure:"DB_USERNAME"`
	DBPASSWORD string `mapstructure:"DB_PASSWORD"`
	DBPORT     string `mapstructure:"DB_PORT"`
	DBNAME     string `mapstructure:"DB_NAME"`
}

type App struct {
	Name string
}
type Server struct {
	Host string
	Port string
}
