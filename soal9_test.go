package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

func Soal9IfElse(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/grade.gohtml"))

	t.ExecuteTemplate(w, "grade.gohtml", map[string]any{
		"score": 50,
	})
}

func Test9(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	Soal9IfElse(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func Soal9Hobi(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/rangetes.gohtml"))

	t.ExecuteTemplate(w, "rangetes.gohtml", map[string]any{
		"Hobi": []string{"Game", "Memancing", "Coding"},
	})
}

func Test9Hobi(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	Soal9Hobi(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

type Produks struct{
	Produk string
	Details *Detail
}

type Detail struct{
	Price int
	Stok int
}

func Soal9Nested(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles("./templates/nested.gohtml"))

	t.ExecuteTemplate(w, "nested.gohtml", Produks{
		Produk: "Laptop TUF",
		Details: &Detail{
			Price: 1200000,
			Stok: 22,
		},
	})
}

func Test9Nested(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	Soal9Nested(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
} 

func TesTemplateSoal9(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFiles(
		"./templates/header.gohtml",
		"./templates/layout.gohtml",
		"./templates/footer.gohtml",
	))
	t.ExecuteTemplate(w, "layout", map[string]any{
		"Name": "Home Template",
	})
}

func TestTemplateSoal9(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TesTemplateSoal9(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}