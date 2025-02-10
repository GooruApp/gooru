package env

import (
	"fmt"
	"os"
	"strconv"
)

const (
	defaultEnv       string = "dev"
	defaultDBConnStr string = "sqlite://booru.db"
	defaultPort      int    = 8000
)

func initAppEnv() string {
	envAppEnv := os.Getenv("APP_ENV")

	if envAppEnv == "" {
		return defaultEnv
	}

	return envAppEnv
}

func initPort() int {
	envPort := os.Getenv("PORT")

	if envPort == "" {
		return defaultPort
	}

	port, err := strconv.ParseUint(envPort, 10, 16)
	if err != nil {
		fmt.Printf("Can't parse PORT from the env, falling back to default: %d", defaultPort)
	}

	return int(port)
}

func initDBConnStr() string {
	envDBConnStr := os.Getenv("DB_CONN_STR")

	if envDBConnStr == "" {
		return defaultDBConnStr
	}

	return envDBConnStr
}
