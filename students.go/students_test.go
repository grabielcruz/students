package students

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"example.com/students/models"
	"github.com/julienschmidt/httprouter"
)

func init() {
	os.Setenv("DBUSER", "postgres")
	os.Setenv("PASSWORD", "1234")
	os.Setenv("HOST", "localhost")
	os.Setenv("PORT", "5432")
	os.Setenv("DBNAME", "students")
}

func TestGetStudents(t *testing.T) {
	router := httprouter.New()
	router.GET("/students", GetStudents)

	req, err := http.NewRequest("GET", "/students", nil)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not make a get request to /students")
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	t.Log("testing OK status code")
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status = %v, want %v", status, http.StatusOK)
	}

	t.Log("testing for an array of students")
	students := []models.Student{}
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not read body of response")
	}

	err = json.Unmarshal(body, &students)
	if err != nil {
		t.Error("Response body does not contain an array of type Student")
	}
}


func TestCreateStudent(t *testing.T) {
	router := httprouter.New()
	router.POST("/students", CreateStudent)

	student := models.Student{}
	student.Name = "Name x"
	student.Surname = "Surname Y"

	jsonStudent, err := json.Marshal(student)
	if err != nil {
		t.Error("Could not marshal json")
	}
	
	requestBody := strings.NewReader(string(jsonStudent))
	req, err := http.NewRequest("POST", "/students", requestBody)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not make a post request to /students")
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("status = %v, want %v", status, http.StatusOK)
	}

	t.Log("testing create student success")
	requestResponse := models.Student{}
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not read body of response")
	}

	err = json.Unmarshal(body, &requestResponse)
	if err != nil {
		log.Fatal(err)
		t.Error("Response body does not containe a student type")
	}
}

func TestCreateStudentNoName(t *testing.T) {
	router := httprouter.New()
	router.POST("/students", CreateStudent)

	student := models.Student{}
	student.Surname = "Surname Y"

	jsonStudent, err := json.Marshal(student)
	if err != nil {
		t.Error("Could not marshal json")
	}
	
	requestBody := strings.NewReader(string(jsonStudent))
	req, err := http.NewRequest("POST", "/students", requestBody)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not make a post request to /students")
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("status = %v, want %v", status, http.StatusBadRequest)
	}

	t.Log("testing create student success")
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not read body of response")
	}

	response := string(body)
	wanted := "Debe especificar el nombre del estudiante"
	if response != wanted {
		t.Errorf("response = %v, wanted %v", response, wanted)
	}
}

func TestCreateStudentNoSurname(t *testing.T) {
	router := httprouter.New()
	router.POST("/students", CreateStudent)

	student := models.Student{}
	student.Name = "some name"
	jsonStudent, err := json.Marshal(student)
	if err != nil {
		t.Error("Could not marshal json")
	}
	
	requestBody := strings.NewReader(string(jsonStudent))
	req, err := http.NewRequest("POST", "/students", requestBody)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not make a post request to /students")
	}

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("status = %v, want %v", status, http.StatusBadRequest)
	}

	t.Log("testing create student success")
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		log.Fatal(err)
		t.Error("Could not read body of response")
	}

	response := string(body)
	wanted := "Debe especificar el apellido del estudiante"
	if response != wanted {
		t.Errorf("response = %v, wanted %v", response, wanted)
	}
}