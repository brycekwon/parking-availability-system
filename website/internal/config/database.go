package config

type DatabaseConfig struct {
	Driver string
	Name   string
	Host   string
	Port   uint16

	Username string
	Password string
	SSLMode  string

	MaxConnectionPool      uint
	MaxIdleConnections     uint
	ConnectionsMaxLifetime string
}

func NewDatabaseConfig() DatabaseConfig {
	return DatabaseConfig{
		Driver: "pgx",
		Name:   "pas-website",
		// Host:   "localhost",
		Host: "postgres",
		Port: 5432,

		Username: "postgres",
		Password: "password",
		SSLMode:  "disable",

		MaxConnectionPool:      4,
		MaxIdleConnections:     4,
		ConnectionsMaxLifetime: "300s",
	}
}
