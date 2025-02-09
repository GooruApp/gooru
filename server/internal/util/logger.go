package util

import (
	"os"

	"go.uber.org/zap"
)

func NewLogger(service string) (*zap.Logger, error) {
	env := os.Getenv("ENV")

	var logger *zap.Logger
	var err error

	switch env {
	case "production":
		logger, err = zap.NewProduction(zap.Fields(
			zap.String("env", env),
			zap.String("service", service),
		))
	default:
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return logger, nil
}
