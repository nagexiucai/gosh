package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	IP string
	Port int
	Dialect string
	Username string
	Password string
	Name string
	Charset string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			IP: "127.0.0.1",
			Port: 3306,
			Dialect: "mysql",
			Username: "bob",
			Password: "nagexiucai.com",
			Name: "employee",
			Charset: "utf8",
		},
	}
}
