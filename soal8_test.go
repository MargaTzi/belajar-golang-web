package golangweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

type Home struct {
	Title string
	Name string
}

//go:embed views/*.gohtml
var viewfs embed.FS

func TesSoalEmbed(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.ParseFS(viewfs, "views/*.gohtml"))

	t.ExecuteTemplate(w, "home.gohtml", Home{
		Title: "Home",
		Name: "Dito House",
	})
}

func TestSoalEmbed(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TesSoalEmbed(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
