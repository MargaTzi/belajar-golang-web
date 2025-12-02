package testing

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome")
	})
	
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello From Go Web")
	})

	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	mux.ServeHTTP(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)

	requestHello := httptest.NewRequest("GET", "http://localhost:8080/hello", nil)
	recorderhello := httptest.NewRecorder()

	mux.ServeHTTP(recorderhello, requestHello)
	response2 := recorderhello.Result()
	body2, _ := io.ReadAll(response2.Body)
	fmt.Println(string(body2))
	fmt.Println(response2.Status)
}

func Halo(w http.ResponseWriter, r *http.Request){
	nama := r.URL.Query().Get("nama")
	if nama == ""{
		fmt.Fprintln(w, "Hello Guest")
	} else {
		fmt.Fprintf(w, "Hello %s", nama)
	}
}

func TestHelloGuest(t *testing.T){
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	Halo(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)
}

func TestHelloWithName(t *testing.T){
	request := httptest.NewRequest("GET", "http://localhost:8080/hello?nama=Dito", nil)
	recorder := httptest.NewRecorder()

	Halo(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)
}

func ShowTags(w http.ResponseWriter, r *http.Request){
	tags := r.URL.Query()["tag"]

	hasil := strings.Join(tags, ", ")
	fmt.Fprintln(w, hasil)
}

func TestTags(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/article?tag=go&tag=web&tag=backend", nil)
	recorder := httptest.NewRecorder()

	ShowTags(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)}