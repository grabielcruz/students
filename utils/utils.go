package utils

import (
	"fmt"
	"log"
	"net/http"
)

func SendInternalServerError(w http.ResponseWriter, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintf(w, "Error del servidor")
	log.Println(err)
}