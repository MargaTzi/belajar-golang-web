package golangweb

// import (
// 	"embed"
// 	"fmt"
// 	"io"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"
// 	"html/template"
// )

// func TemplateSoal1(w http.ResponseWriter, r *http.Request){
// 	funcs := template.FuncMap{
// 		"greet": func (name string) string {
// 			return "Hai " + name
// 		},
// 		"upper": strings.ToUpper,
// 	}

// 	t := template.Must(template.New("soal1").Funcs(funcs).Parse(`
// 		{{greet .Name | upper}}
// 	`))

// 	t.ExecuteTemplate(w, "soal1", MyPage{
// 		Name: "Wahyu",
// 	})
// }

// func TestGlobal1(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	TemplateSoal1(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }

// //go:embed templates/*.gohtml
// var tmpl1 embed.FS

// var funcMap = template.FuncMap{
// 	"greet":  func (name string) string {
// 		return "hai " + name
// 	},
// 	"lower": strings.ToLower,
// }

// var myTmpl1 = template.Must(template.New("").Funcs(funcMap).ParseFS(tmpl1, "templates/*.gohtml"))

// func TemplateSoal2(w http.ResponseWriter, r *http.Request){
// 	myTmpl1.ExecuteTemplate(w, "halo", MyPage{
// 		Name: "Wahyu",
// 	})
// }

// func TestGlobal2(t *testing.T) {
// 	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	recorder := httptest.NewRecorder()

// 	TemplateSoal2(recorder, request)

// 	response := recorder.Result()
// 	body, _ := io.ReadAll(response.Body)

// 	fmt.Println(string(body))
// }