package config

type ConfigExample struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfigExample() *Config {
	return &Config{
		DBConnectionString: "host=localhost user=postgres password=root dbname=redsocial sslmode=disable",
	}
}
