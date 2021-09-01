package students

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/students/database"
	"example.com/students/models"
	"example.com/students/utils"
	"github.com/julienschmidt/httprouter"
)

func GetStudents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var students []models.Student
	db, err := database.ConnectDB()
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	defer db.Close()
	studentsQuery := "SELECT * FROM students;"

	rows, err := db.Query(studentsQuery)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}

	for rows.Next() {
		student := models.Student{}
		if err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Code, &student.Grade, &student.Birthdate, &student.PublicId, &student.Photo); err != nil {
			utils.SendInternalServerError(w, err)
			return
		}
		students = append(students, student)
	}

	response, err := json.Marshal(students)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func CreateStudent(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var student models.Student
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "No se pudo leer el cuerpo de la petición")
		return
	}

	err = json.Unmarshal(body, &student)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "La data recibida no es del tipo Student")
		return
	}

	if student.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Debe especificar el nombre del estudiante")
		return
	}

	if student.Surname == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Debe especificar el apellido del estudiante")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	defer db.Close()

	posgrestDate := student.Birthdate.Format("2006-01-02")
	createStudentQuery := "INSERT INTO students (name, surname, code, grade, birthdate, public_id, photo) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id, name, surname, code, grade, birthdate, public_id, photo;"
	createdStudent := models.Student{}

	row := db.QueryRow(createStudentQuery, student.Name, student.Surname, student.Code, student.Grade, posgrestDate, student.PublicId, student.Photo)

	err = row.Scan(&createdStudent.Id, &createdStudent.Name, &createdStudent.Surname, &createdStudent.Code, &createdStudent.Grade, &createdStudent.Birthdate, &createdStudent.PublicId, &createdStudent.Photo)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}

	response, err := json.Marshal(createdStudent)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
