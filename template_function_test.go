package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"text/template"
)

type MyPage struct {
	Name string
}

func (Asu MyPage) SayHello(name string) string {
	return "Hai " + name + ", My Name Is " + Asu.Name
}

func TemplateFunction(w http.ResponseWriter, r *http.Request){
	t := template.Must(template.New("Function").Parse(`{{.SayHello "Wahyu"}}`))

	t.ExecuteTemplate(w, "Function", MyPage{
		Name: "Dito",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateFunctionGlobal(w http.ResponseWriter, r *http.Request){
	
	t := template.Must(template.New("Asu").Parse(`{{len .Name}}`))

	t.ExecuteTemplate(w, "Asu", MyPage{
		Name: "Yanto",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateMapGlobal(w http.ResponseWriter, r *http.Request){
	// funcMap := template.FuncMap{
	// 	"lower": strings.ToLower,
	// }

	// t := template.New("tes").Funcs(funcMap)
	// t = template.Must(t.Parse(`{{lower .Name}}`))

	// t.ExecuteTemplate(w, "tes", MyPage{
	// 	Name: "WAHYU",
	// })

	t := template.New("tes")
	t = t.Funcs(map[string]any{
		"lower": func (s string) string {
			return strings.ToLower(s)
		},
	})

	t = template.Must(t.Parse(`{{lower .Name}}`))
	t.ExecuteTemplate(w, "tes", MyPage{
		Name: "Home",
	})
}

func TestTemplateMapGlobal(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateMapGlobal(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func TemplateGlobalPipeline(w http.ResponseWriter, r *http.Request){
	// funcs := template.FuncMap{
	// 	"hai" : func (v string) string {
	// 		return "Hai " + v
	// 	},
	// 	"upper": func (s string) string  {
	// 		return strings.ToUpper(s)
	// 	},
	// }

	funcs := template.FuncMap{
		"hai" : func (v string) string {
			return "Wahyu cok" + v
		},
		"upper": strings.ToUpper,
	}

	t := template.Must(template.New("tes").Funcs(funcs).Parse(`
		{{hai .Name | upper}}
	`))

	t.ExecuteTemplate(w, "tes", MyPage{
		Name: "Asu",
	})
}

func TestTemplateGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateGlobalPipeline(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}