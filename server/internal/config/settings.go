package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const appEnvPrefix = "GOORU"

var Settings = struct {
	Mode      setting[string]
	DBConnStr setting[string]
	Port      setting[int]
}{
	Mode: setting[string]{
		envVar:       "MODE",
		defaultValue: "dev",
		initValue: func(mode string) (string, error) {
			if mode != "dev" && mode != "production" {
				return "", errors.New("must be 'dev' or 'production'")
			}
			return mode, nil
		},
	},
	DBConnStr: setting[string]{
		envVar:       "DB_CONNECTION",
		defaultValue: "sqlite://booru.db",
		initValue: func(conn string) (string, error) {
			if !strings.HasPrefix(conn, "sqlite://") && !strings.HasPrefix(conn, "postgresql://") {
				return "", errors.New("invalid protocol, must be 'sqlite' or 'postgresql'")
			}
			return conn, nil
		},
	},
	Port: setting[int]{
		envVar:       "PORT",
		defaultValue: 8000,
		initValue: func(portStr string) (int, error) {
			port, err := strconv.Atoi(portStr)
			if err != nil {
				return 0, err
			}
			if port < 0 || port > 65535 {
				return 0, errors.New("out of valid port range (0 - 65535)")
			}
			return port, nil
		},
	},
}

type setting[T any] struct {
	envVar       string
	defaultValue T
	initValue    func(string) (T, error)
	value        *T
}

func (s *setting[T]) init() *T {
	envVar := fmt.Sprintf("%v_%v", appEnvPrefix, s.envVar)
	envValue, found := os.LookupEnv(envVar)

	if !found {
		fmt.Printf("WARN: environment variable %v not set, defaulting to '%v'\n", envVar, s.defaultValue)
		return &s.defaultValue
	} else {
		validValue, err := s.initValue(envValue)
		if err != nil {
			fmt.Printf("WARN: invalid value (%v) for environment variable %v: %v; reverting to default value '%v'\n", envValue, envVar, err, s.defaultValue)
			return &s.defaultValue
		} else {
			return &validValue
		}
	}
}

func (s *setting[T]) Get() T {
	if s.value != nil {
		return *s.value
	}

	s.value = s.init()
	return *s.value
}
