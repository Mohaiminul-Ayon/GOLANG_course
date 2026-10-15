package config

import (
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
)

var configaration Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
	JwtSecreteKey string
}

func loadConfig() {
	err := godotenv.Load()
	if err !=nil{
		fmt.Println("Faild to load env variable",err)
	}
	version := os.Getenv("VERSION")
	if version == "" {
		fmt.Println("Version is reqiuired")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		fmt.Println("Service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		fmt.Println("Http Port is required")
		os.Exit(1)
	}
	
	port, err := strconv.ParseInt(httpPort, 10, 64) //string ,base , bit
	if err != nil {
		fmt.Println("Invalid HTTP_PORT:", err)
		os.Exit(1)
	}
	jwtSecreteKey := os.Getenv("JWT_SECRET_KEY")
	if jwtSecreteKey == "" {
		fmt.Println("secret key is required")
		os.Exit(1)
	}
	
	configaration = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    int(port),
		JwtSecreteKey: jwtSecreteKey,
	}

}

func Getconfig() Config {
	loadConfig()
	return configaration
}
