package config

type Config struct {
	ServerConfig
	DatabaseConfig
	CacheConfig
}

func New() *Config {
	return &Config{
		ServerConfig:   NewServerConfig(),
		DatabaseConfig: NewDatabaseConfig(),
		CacheConfig:    NewCacheConfig(),
	}
}
