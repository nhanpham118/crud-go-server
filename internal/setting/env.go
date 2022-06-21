package setting

import "os"

func NewServerConfig() Config {
	return Config{
		MYSQL_PASS: os.Getenv("MYSQL_PASS"),
		MYSQL_HOST: os.Getenv("MYSQL_HOST"),
		MYSQL_PORT: os.Getenv("MYSQL_PORT"),
		MYSQL_DB:   os.Getenv("MYSQL_DB"),
	}
}

type Config struct {
	MYSQL_PASS string
	MYSQL_HOST string
	MYSQL_PORT string
	MYSQL_DB   string
}
