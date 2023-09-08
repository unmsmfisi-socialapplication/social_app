package config

type Config struct {
	DBConnectionString string
	AppPort            string
}

//Defines the connection string to the database
func LoadConfig() *Config {
	return &Config{
		DBConnectionString: "",
	}
}
