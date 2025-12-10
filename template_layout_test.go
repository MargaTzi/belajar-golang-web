package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateLayout(w http.ResponseWriter, r *http.Request){
	t:= template.Must(template.ParseFiles(
		"./templates/header.gohtml", 
		"./templates/footer.gohtml", 
		"./templates/layout.gohtml", 
	))

	t.ExecuteTemplate(w, "layout", map[string]any{
		"Name" : "Wahyu",
		"Title" : "Template Layout",
	})
}

func TestTemplateLayout(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateLayout(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}