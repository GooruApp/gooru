package api

import (
	"context"
	"fmt"
	"net/http"

	"go.uber.org/zap"
)

type api struct {
	ctx    context.Context
	logger *zap.Logger
}

func NewAPI(ctx context.Context, logger *zap.Logger) *api {
	return &api{
		ctx:    ctx,
		logger: logger,
	}
}

func (a *api) Server(port int) *http.Server {
	stack := createMiddlewareStack(
		a.loggingMiddleware,
		a.requestIdMiddleware,
		a.corsMiddleware,
	)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: stack(a.Routes()),
	}
}

func (a *api) Routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /", a.getDebugHandler)

	return r
}
