package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	firstName := r.PostForm.Get("firstName")
	lastName := r.PostForm.Get("lastName")
	fmt.Fprintf(w, "%s %s", firstName, lastName)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("firstName=Dito&lastName=Mada")
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)
	
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}