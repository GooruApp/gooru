package logger

import (
	"github.com/GooruApp/gooru/server/internal/environment"
	"go.uber.org/zap"
)

func New(service string, env *environment.Environment) (*zap.Logger, error) {
	appEnv := env.AppEnv()

	var logger *zap.Logger
	var err error

	switch appEnv {
	case "production":
		logger, err = zap.NewProduction(zap.Fields(
			zap.String("env", appEnv),
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
