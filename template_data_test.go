package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]any{
		"Title": "Template Data Struct",
		"Name": "Wahyu",
		"Address": map[string]any{
			"Street": "Jalan Rusak !!",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type Address struct{
	Street string
}

type Page struct{
	Title string
	Name string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Struct",
		Name: "Wahyu",
		Address: Address{
			Street: "Jalan Sedang Rusak",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}