package api

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Middleware func(http.Handler) http.Handler

func createMiddlewareStack(ms ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(ms) - 1; i >= 0; i-- {
			m := ms[i]
			next = m(next)
		}

		return next
	}
}

type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lwr *LoggingResponseWriter) WriteHeader(statusCode int) {
	lwr.ResponseWriter.WriteHeader(statusCode)
	lwr.statusCode = statusCode
}

func (a *api) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next.ServeHTTP(w, r)
	})
}

func (a *api) requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("gooru-request-id", id)
		next.ServeHTTP(w, r)
	})
}

func (a *api) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lwr := &LoggingResponseWriter{
			ResponseWriter: w,
			statusCode:     http.StatusOK,
		}

		next.ServeHTTP(lwr, r)

		duration := time.Since(start).Milliseconds()

		fields := []zap.Field{
			zap.Int64("duration", duration),
			zap.String("method", r.Method),
			zap.Int("status", lwr.statusCode),
			zap.String("uri", r.RequestURI),
			zap.String("requestId", lwr.Header().Get("gooru-request-id")),
		}

		if lwr.statusCode == http.StatusOK {
			a.logger.Info("", fields...)
		} else {
			// We should log the error message if one occurs on an endpoint
			// This should be attached to the response in the http handler as a header
			// named "gooru-error"
			//
			// err := lwr.Header().Get("gooru-error")
			a.logger.Error("<err>", fields...)
		}
	})
}
