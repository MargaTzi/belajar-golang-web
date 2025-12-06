package golangweb

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(w http.ResponseWriter, r *http.Request){
	cookie := new(http.Cookie)
	cookie.Name ="Oreon-Prom"
	cookie.Value = r.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(w, cookie)
	fmt.Fprint(w, "succes create cookie")
}

func GetCookie(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie("Oreon-Prom")
	if err != nil {
		log.Fatalln(err)
	} else {
		name := cookie.Value
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=Dito", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()
	for _, cookie := range cookies{
		fmt.Printf("Cookie %s:%s \n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "Oreon-Prom"
	cookie.Value = "Wahyu"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}