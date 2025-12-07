package testing

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Profile struct {
	Title string
	Name  string
	Age   int
}

func TesTemplateFile(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("../templates/file.gohtml"))

	t.ExecuteTemplate(w, "file.gohtml", Profile{
		Title: "Testing",
		Name: "Wahyu",
		Age: 10,
	})
}

func TestTemplateFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TesTemplateFile(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}