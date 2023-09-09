package config

type Config struct {
	DBConnectionString string
	AppPort            string
}

// Defines the connection string to the database
func LoadConfig() *Config {
	return &Config{
		DBConnectionString: "host=localhost user=postgres password=root dbname=redsocial sslmode=disable",
	}
}
