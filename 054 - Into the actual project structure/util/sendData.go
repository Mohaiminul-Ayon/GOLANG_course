package util

import (
	"encoding/json"
	"net/http"
)

func SendData(w http.ResponseWriter, data interface{}, statusCode int) {
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
}

func Senderror (w http.ResponseWriter,msg string, statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(msg)
	
}