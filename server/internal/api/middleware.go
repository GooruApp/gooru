package api

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
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

func (a *api) requestIdMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-Id", id)
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

		log.Println(lwr.statusCode, r.Method, r.URL.Path, time.Since(start))
	})
}
