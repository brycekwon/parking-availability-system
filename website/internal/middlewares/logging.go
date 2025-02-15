package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

type wrappedWriter struct {
	http.ResponseWriter
	status int
}

func (w *wrappedWriter) WriteHeader(status int) {
	w.ResponseWriter.WriteHeader(status)
	w.status = status
}

func Logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		wrapped := &wrappedWriter{
			ResponseWriter: w,
			status:         http.StatusOK,
		}

		defer func() {
			logger.Info(
				"handled request",
				slog.Int("status", wrapped.status),
				slog.String("remoteAddr", r.RemoteAddr),
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.Any("duration", time.Since(start)),
			)
		}()

		next.ServeHTTP(wrapped, r)
	})
}
