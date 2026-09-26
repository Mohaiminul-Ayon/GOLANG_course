package middleware

import (
	"net/http"
)


func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//handle_corse_issue
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,PATCH,DELETE,POST,OPTION")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Ayon")
		w.Header().Set("Content-Type", "application/json")
		
		next.ServeHTTP(w, r)
	})
}
