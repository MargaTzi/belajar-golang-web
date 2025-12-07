package testing

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Home struct{
	Title string
	Name string
}

func TesTemplateGlob(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseGlob("../views/*.gohtml"))

	data := Home{
		Title: "Inventory",
		Name: "Wahyu",
	}

	t.ExecuteTemplate(w, "home.gohtml", data)
}

func TestTemplateGlob(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TesTemplateGlob(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

