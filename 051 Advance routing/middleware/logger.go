package middleware

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start:= time.Now()//8.55.20

		log.Println("Ami middleware age print hobo")
		next.ServeHTTP(w,r) //10s
		//curr 8.55.30
		log.Println("Amio middleware pore print hobo")
		log.Println(r.Method, r.URL.Path, time.Since(start))

	})



}
