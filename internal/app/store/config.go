package store

type Config struct {
	DatabaseURL string `toms: "database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
