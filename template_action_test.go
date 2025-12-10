package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Person struct{
	Name string
	Title string
}

func TemplateActionIf(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", Person{
		Name: "",
		Title: "Home",
	})
}

func TestActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateActionOperator(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]any{
		"Nilai" : 50,
	})
}

func TestActionOperator(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionOperator(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]any{
		"Hobbies" : []string{"Game", "Coding", "Belajar"},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type Alamat struct {
    Jalan string
    KOta   string
}

type Homes struct {
    Title   string
    Name    string
    Alamat Alamat
}


func TemplateWith(w http.ResponseWriter, r * http.Request){
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	data := Homes{
		Title: "Home",
		Name: "Wahyu",
		Alamat: Alamat{
			Jalan: "BAngetayu",
			KOta: "Semarang",
		},
	}

	t.ExecuteTemplate(w, "address.gohtml", data)
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}