package config

import (
	"log"
	"os"
)

type Config struct {
	Version     string
	ServiceName string
	HttpPort    int
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

}
