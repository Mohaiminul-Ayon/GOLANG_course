package handlers

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {

	log.Println("Ami handler middle e print hobo")
}
