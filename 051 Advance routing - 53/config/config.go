package config

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct{
	Version string 
	ServiceName string
	HttpPort int
}

func LoadConfig(){
	version := os.Getenv("VERSION")
	if version ==""{
		fmt.Println("Version is reqiuired")
		os.Exit(1)
	}
	
	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == ""{
		fmt.Println("Service name is required")
		os.Exit(1)
	}
	
	httpPort := os.Getenv("HTTP_PORT")
	if  httpPort == ""{
		fmt.Println("Http Port is required")
		os.Exit(1)
	}

	port,err:=strconv.ParseInt(httpPort,10,64)

	cnf:= Config{
		Version: version,
		ServiceName: serviceName,
		HttpPort: httpPort,
	}

}