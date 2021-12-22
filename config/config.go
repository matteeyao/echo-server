package config

type Configurations struct {
	Server       	  ServerConfigurations
	REDIRECT_HOST	  string
	REDIRECT_ENDPOINT string
}

type ServerConfigurations struct {
	Host string
	Port string
}
