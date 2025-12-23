package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func (middlewareCostume *MiddlewareCostume) Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// before
		start := time.Now()

		next.ServeHTTP(w, r)

		// after
		duration := time.Since(start)
		middlewareCostume.Log.Info("Activity route", zap.String("Method", r.Method), zap.String("URL", r.URL.String()), zap.Duration("Duration", duration))
		// log.Printf("Method : %s, URL : %s, Duration : %s\n", r.Method, r.URL.String(), duration.String())
	})
}
