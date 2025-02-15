package config

type CacheConfig struct {
	Host string
	Port uint16

	Username string
	Password string
}

func NewCacheConfig() CacheConfig {
	return CacheConfig{
		// Host:   "localhost",
		Host: "redis",
		Port: 6379,

		Username: "redis",
		Password: "password",
	}
}
