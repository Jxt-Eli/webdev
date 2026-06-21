package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		start := time.Now()
		f.ServeHTTP(w, r)
		end := time.Since(start)
		fmt.Printf("Time taken: %s", end)
	})
}
func APIKeyMiddleware(f http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		if key != "secretkey" {
			http.Error(w, "Bad Request", http.StatusUnauthorized)
			return
		}
		fmt.Println("Key accepted, unlocking ....")
		f.ServeHTTP(w, r)
	})
}
