package golangweb

// import (
// 	"embed"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"
// 	"text/template"
// )

// //go:embed templates/*.gohtml
// var tmpl embed.FS

// var myTmpl = template.Must(template.ParseFS(tmpl, "templates/*.gohtml"))

// func TemplateCaching(w http.ResponseWriter, r *http.Request){
// 	myTmpl.ExecuteTemplate(w, "halo.gohtml", "nil")
// }

// func TestTemplateCaching(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	TemplateCaching(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }