package middleware

import "net/http"

func CORS(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "127.0.0.1")
		next.ServeHTTP(w, r)
	})
}