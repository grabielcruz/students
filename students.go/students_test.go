package students

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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

func TestGetStudents(t *testing.T)  {
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

