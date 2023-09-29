package config

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfig() *Config {
	return &Config{
		DBConnectionString: "host=localhost user=postgres password=1234 dbname=unmsm_social_app port=5432",
	}
}
