package main

import (
	"fmt"
	c "github.com/matteeyao/http-server/config"
	"github.com/spf13/viper"
	"log"
)

func viperEnvVariable(key string) string {
	viper.SetConfigName("config")

	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	viper.SetConfigType("yml")
	var configuration c.Configurations

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	err := viper.Unmarshal(&configuration)
	if err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion")
	}
	return value
}

func main() {
	host := viperEnvVariable("server.host")
	port := viperEnvVariable("server.port")
	address := host + ":" + port

	mux := &myMux{}
	listenAndServe(address, mux)
}
