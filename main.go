package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/students/database"
	"example.com/students/students.go"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/", index)

	router.GET("/students", students.GetStudents)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "index route")
	db := database.ConnectDB()
	defer db.Close()
}
