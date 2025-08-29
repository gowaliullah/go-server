package config

import (
	"log"
	"os"
	"strconv"
)

var config Config

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int64
}

func LoadConfig() {
	version := os.Getenv("VERSION")
	if version == "" {
		log.Fatal("version is required")
		os.Exit(1)
	}

	serviceName := os.Getenv("SERVICE_NAME")
	if serviceName == "" {
		log.Fatal("service name is required")
		os.Exit(1)
	}

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		log.Fatal("http port is required")
		os.Exit(1)
	}

	port, err := strconv.ParseInt(httpPort, 10, 64)
	if err != nil {
		log.Fatal("post must be a number")
		os.Exit(1)
	}

	config = Config{
		Version:     version,
		ServiceName: serviceName,
		HttpPort:    port,
	}

}
