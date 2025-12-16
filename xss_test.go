package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//go:embed templates/*.gohtml
var tmpl1 embed.FS

var funcMap = template.FuncMap{
	"greet":  func (name string) string {
		return "hai " + name
	},
	"lower": strings.ToLower,
}

var myTmpl1 = template.Must(template.New("").Funcs(funcMap).ParseFS(tmpl1, "templates/*.gohtml"))


func TemplateAutoEscape(w http.ResponseWriter, r *http.Request){
	
	myTmpl1.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body" : "<p>Ini Adalah Body<script>alert('Anda di Heck')</script></p>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateAutoEscapeDisabled(w http.ResponseWriter, r *http.Request){
	
	myTmpl1.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body" : template.HTML("<p>Hello World</p>"),
	})
}

func TestAutoEscapeDisabled(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscapeDisabled(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestAutoEscapeServerDisabled(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisabled),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TemplateXSS(w http.ResponseWriter, r *http.Request){
	
	myTmpl1.ExecuteTemplate(w, "post.gohtml", map[string]any{
		"Title": "Golang Auto Escape",
		"Body" : template.HTML(r.URL.Query().Get("body")),
	})
}

func TestXSS(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateXSS(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TestServerXSS(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}