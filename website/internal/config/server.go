package config

type ServerConfig struct {
	Name string
	Host string
	Port uint16
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		Name: "Parking Availability System Website",
		Host: "0.0.0.0",
		Port: 3000,
	}
}
