package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config *Config

type DBConfig struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	EnableSSLMODE bool
}

type Config struct {
	Version      string
	ServiceName  string
	HttpPort     int
	JwtSecretKey string
	DB           *DBConfig
}

func loadConfig() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Fail to load .env variables: ", err)
		os.Exit(1)
	}

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

	jwtSecretKey := os.Getenv("JWT_SECTER_KEY")
	if jwtSecretKey == "" {
		log.Fatal("JWT_SECTER_KEY is required")
		os.Exit(1)
	}

	// start DB config

	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("DB HOST is required")
		os.Exit(1)
	}

	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		log.Fatal("DB PORT is required")
		os.Exit(1)
	}

	dbPrt, err := strconv.ParseInt(dbPort, 10, 64)
	if err != nil {
		fmt.Println("DB post must be a number")
		os.Exit(1)
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("DB NAME is required")
		os.Exit(1)
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("DB USER is required")
		os.Exit(1)
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("PASSWORD is required")
		os.Exit(1)
	}

	enableSslMode := os.Getenv("DB_ENABLE_SSL_MODE")
	enblSSMode, err := strconv.ParseBool(enableSslMode)
	if err != nil {
		fmt.Println("Invalid ssl mode value", err)
		os.Exit(1)
	}

	fmt.Println(dbPrt)

	dbConfig := &DBConfig{
		Host:          dbHost,
		Port:          int(dbPrt),
		Name:          dbName,
		User:          dbUser,
		Password:      dbPassword,
		EnableSSLMODE: enblSSMode,
	}

	config = &Config{
		Version:      version,
		ServiceName:  serviceName,
		HttpPort:     int(port), // type casting ()
		JwtSecretKey: jwtSecretKey,
		DB:           dbConfig,
	}

}

func GetConfig() *Config { // tight coupling

	if config == nil {
		loadConfig()
	}
	return config
}
