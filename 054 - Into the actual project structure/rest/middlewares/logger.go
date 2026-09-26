package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()//8.55.20

		next.ServeHTTP(w,r) //10s
		//curr 8.55.30
		log.Println(r.Method, r.URL.Path, time.Since(start))

	})



}
