package golangweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Halo Dunia")
}

func TestHelloHandler(t *testing.T){
	request := httptest.NewRequest("GET", "http://localhost:8080/halo", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
	fmt.Println(response.Header)
}