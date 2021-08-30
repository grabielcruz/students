package students

import (
	"encoding/json"
	"fmt"
	"net/http"

	"example.com/students/database"
	"example.com/students/models"
	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var students []models.Student
	db := database.ConnectDB()
	defer db.Close()
	studentsQuery := "SELECT * FROM students;"

	rows, err := db.Query(studentsQuery)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	for rows.Next() {
		student := models.Student{}
		if err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Birthdate, &student.PublicId, &student.Photo); err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		students = append(students, student)
	}

	response, err := json.Marshal(students)
	if err != nil {
		fmt.Fprint(w, err.Error())
		return
	}

	w.Write(response)

}