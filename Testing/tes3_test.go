package testing

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func LoginForm(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if username == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "username wajib diisi")
	} 
	if password == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "password wajib diisi")
	}  

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login berhasil : %s", username)
	
}

func TestLoginForm(t *testing.T) {
	requestBody := strings.NewReader("username=dito&password=")
	request := httptest.NewRequest("POST", "http://localhost/", requestBody)
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded") 
	recorder := httptest.NewRecorder()

	LoginForm(recorder, request)
	res := recorder.Result()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(string(body))
	fmt.Println(res.StatusCode)
}