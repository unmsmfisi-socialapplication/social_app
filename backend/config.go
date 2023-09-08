package config

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfig() *Config {
	return &Config{
		DBConnectionString: "host=localhost user=postgres password=root dbname=redsocial sslmode=disable",
	}
}
