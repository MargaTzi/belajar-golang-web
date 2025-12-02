package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request){
	name := r.URL.Query().Get("nama")
	if name == ""{
		w.WriteHeader(400)
		fmt.Fprint(w, "nama kosong cok")
	} else {
		fmt.Fprintf(w, "Halo Su %s", name)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}
func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?nama=Dito", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.Status)
	fmt.Println(response.StatusCode)
}