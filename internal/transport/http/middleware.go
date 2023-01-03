package http

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
		next.ServeHTTP(writer, request)
	})
}
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		logrus.WithFields(
			logrus.Fields{
				"method": request.Method,
				"path":   request.URL.Path,
			}).Info("handled request")
		next.ServeHTTP(writer, request)
	})
}

func TimeoutMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(request.Context(), 8*time.Second)
		defer cancel()
		next.ServeHTTP(writer, request.WithContext(ctx))
	})
}
