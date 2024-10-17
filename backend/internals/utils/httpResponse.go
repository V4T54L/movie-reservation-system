package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, statusCode int, obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println("[-] error marshaling the object")
	}
	w.WriteHeader(statusCode)
	w.Write(data)
}

func ErrorResponse(w http.ResponseWriter, statusCode int, err string) {
	JSONResponse(w, statusCode, map[string]string{
		"error": err,
	})
}

func MessageResponse(w http.ResponseWriter, statusCode int, msg string) {
	JSONResponse(w, statusCode, map[string]string{
		"message": msg,
	})
}
