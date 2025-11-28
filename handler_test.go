package golangweb

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestHander(t *testing.T){
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Halo, ini method: ", r.Method)
		fmt.Fprintln(w, "Dan path: ", r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "Halo Dunia")
	})
	mux.HandleFunc("/cok", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "Wahyu Toy Story")
	})
	mux.HandleFunc("/dashboards/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "Hai Dashboard")
	})
	mux.HandleFunc("/dashboards/profile/", func(w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "Hai Profile")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}