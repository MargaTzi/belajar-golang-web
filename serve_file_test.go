package golangweb

import (
	_ "embed"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func ServeFile(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		http.ServeFile(w, r, "./resource/ok.html")
	} else {
		http.ServeFile(w, r, "./resource/notFound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

//go:embed resource/ok.html
var resourceOk string

//go:embed resource/notFound.html
var resourceNotFound string

func ServeFileEmbed(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != ""{
		fmt.Fprint(w, resourceOk)
	} else {
		fmt.Fprint(w, resourceNotFound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed), 
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}