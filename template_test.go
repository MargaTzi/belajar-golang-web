package golangweb

import (
	"embed"
	"fmt"
	"io"
	"text/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Template(w http.ResponseWriter, r *http.Request){
	ttext := `<html><body>{{.}}</body></html>`
	// t, err := template.New("Halo").Parse(ttext)
	// if err != nil{
	// 	log.Fatalln(err)
	// }

	t := template.Must(template.New("Halo").Parse(ttext))

	t.ExecuteTemplate(w, "Halo", "Hello World")
}

func TestTemplate(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	Template(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFile(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/halo.html"))
	t.ExecuteTemplate(w, "halo.html", "Wahyu Ndlogok")
}

func TestTemplateFile(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFile(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateGlob(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))
	t.ExecuteTemplate(w, "halo.gohtml", "Wahyu Toy Story")
}

func TestTemplateGlob(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateGlob(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var tmplFS embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(tmplFS, "templates/*.gohtml"))
	t.ExecuteTemplate(w, "halo.gohtml", "Wahyu Toy Story")
}

func TestTemplateEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateEmbed(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}