package apiserver

type Config struct {
	BindAddr    string `toml: "bind_addr`
	LogLevel    string `toml: "log_level"`
	DatabaseURL string `tom: "database_url"`
	SessionKey  string `toml: "session_key"`
}

func NewConfig() *Config {
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
	}
}
