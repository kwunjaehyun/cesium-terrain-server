package handlers

import (
	"fmt"
	"net/http"	
)

type Bytes uint64

type ResponseLimiter interface {
	http.ResponseWriter
	LimitExceeded() bool
}

type LimiterFactory func(writer http.ResponseWriter, limit Bytes) ResponseLimiter

// Return HTTP middleware which allows CORS requests from any domain
func AddCorsHeader(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headers := w.Header()
		headers.Set("Access-Control-Allow-Origin", "*")
		headers.Set("Cache-Control", "max-age=86400")
		
		fmt.Sprintf("test %d", "abc")
		next.ServeHTTP(w, r)
	})
}
