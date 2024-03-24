package middleware

import (
	"log"
	"net/http"
	"time"
)

//todo

type writerWrapper struct {
	http.ResponseWriter
	Status int
}

func (w *writerWrapper) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.Status = statusCode
}
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		status := &writerWrapper{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		next.ServeHTTP(status, r)
		log.Println(status.Status, r.Method, r.URL.Path, time.Since(start))
	})
}
