package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("name")
	if name == ""{
		fmt.Fprint(w, "Cahyo")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?name=", nil)
	simpan_request := httptest.NewRecorder()

	SayHello(simpan_request, request)
	response := simpan_request.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
}

func MultipleQuery(w http.ResponseWriter, r *http.Request){
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")
	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func TestMultiple(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/home?first_name=Dito&last_name=Golang", nil)
	recorder := httptest.NewRecorder()

	MultipleQuery(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
}


func MultipleQueryParam(w http.ResponseWriter, r *http.Request){
	names := r.URL.Query()["nama"]

	for _, row := range names{
		fmt.Fprintln(w, row)
	}

}

func TestMultipleParam(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/home?nama=Dito&nama=Wahyu", nil)
	recorder := httptest.NewRecorder()

	MultipleQueryParam(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.Status)
}

