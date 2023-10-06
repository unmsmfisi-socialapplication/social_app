package utils

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfig() *Config {
	return &Config{
		//DBConnectionString: "host=35.193.19.32 user=social-app-dbuser password=oc(11b-ABvU]m6um dbname=db-social-app-dev sslmode=disable",
		DBConnectionString: "host=localhost user=postgres password=root dbname=test3 sslmode=disable",
	}
}
