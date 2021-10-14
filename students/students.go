package students

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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
	studentsQuery := "SELECT * FROM students ORDER BY id DESC;"

	rows, err := db.Query(studentsQuery)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}

	for rows.Next() {
		student := models.Student{}
		if err := rows.Scan(&student.Id, &student.Name, &student.Surname, &student.Code, &student.Grade, &student.Section, &student.Birthdate, &student.PublicId, &student.Photo); err != nil {
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

	// posgrestDate := student.Birthdate.Format("2006-01-02")
	createStudentQuery := "INSERT INTO students (name, surname, code, grade, section, birthdate, public_id, photo) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, name, surname, code, grade, section, birthdate, public_id, photo;"
	createdStudent := models.Student{}

	row := db.QueryRow(createStudentQuery, student.Name, student.Surname, student.Code, student.Grade, student.Section, student.Birthdate, student.PublicId, student.Photo)

	err = row.Scan(&createdStudent.Id, &createdStudent.Name, &createdStudent.Surname, &createdStudent.Code, &createdStudent.Grade, &createdStudent.Section, &createdStudent.Birthdate, &createdStudent.PublicId, &createdStudent.Photo)
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

func UpdateStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var student models.Student
	studentId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Debe especificar el id del estudiante que quires actualizar")
		return
	}

	if studentId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "El id del estudiante debe ser un número mayor a cero")
		return
	}

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

	// posgrestDate := student.Birthdate.Format("2006-01-02")
	updateStudentQuery := "UPDATE students SET name = $1, surname = $2, code = $3, grade = $4, section = $5, birthdate = $6, public_id = $7, photo = $8 WHERE id = $9 RETURNING id, name, surname, code, grade, section, birthdate, public_id, photo;"
	updatedStudent := models.Student{}

	row := db.QueryRow(updateStudentQuery, student.Name, student.Surname, student.Code, student.Grade, student.Section, student.Birthdate, student.PublicId, student.Photo, studentId)

	err = row.Scan(&updatedStudent.Id, &updatedStudent.Name, &updatedStudent.Surname, &updatedStudent.Code, &updatedStudent.Grade, &updatedStudent.Section, &updatedStudent.Birthdate, &updatedStudent.PublicId, &updatedStudent.Photo)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "El id especificado no pertenece a ningún estudiante")
			return
		}
		utils.SendInternalServerError(w, err)
		return
	}

	response, err := json.Marshal(updatedStudent)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	studentId, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Debe especificar el id del estudiante que desea eliminar")
		return
	}

	if studentId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "El id del estudiante debe ser un número mayor a cero")
		return
	}

	db, err := database.ConnectDB()
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	defer db.Close()

	deleteStudentQuery := "DELETE FROM students WHERE id = $1 RETURNING id;"
	deletedStudentId := models.IdResponse{}

	row := db.QueryRow(deleteStudentQuery, studentId)

	err = row.Scan(&deletedStudentId.Id)
	if err != nil {
		if err.Error() == "sql: no rows in result set" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "El id especificado no pertenece a ningún estudiante")
			return
		}
		utils.SendInternalServerError(w, err)
		return
	}

	response, err := json.Marshal(deletedStudentId)
	if err != nil {
		utils.SendInternalServerError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}