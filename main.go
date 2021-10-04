package main

import (
	"fmt"
	"log"
	"net/http"

	"example.com/students/handle_uploads"
	"example.com/students/students"
	"github.com/julienschmidt/httprouter"
)

func CustomOptions(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// Enable Cors
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h(w, r, ps)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", index)

	router.GET("/students", CustomOptions(students.GetStudents))
	router.POST("/students", CustomOptions(students.CreateStudent))
	router.PUT("/students/:id", CustomOptions(students.UpdateStudent))
	router.DELETE("/students/:id", CustomOptions(students.DeleteStudent))

	router.ServeFiles("/public/*filepath", http.Dir("./public"))
	router.POST("/uploadphoto/", CustomOptions(handle_uploads.UploadPhoto))

	log.Fatal(http.ListenAndServe(":8080", router))
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "index route")
}
