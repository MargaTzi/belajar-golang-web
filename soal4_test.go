package golangweb

import (
	"log"
	"net/http"
	"testing"
)

func TestSoalFileServer(t *testing.T) {
	dir := http.Dir("./public")
	fs := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets", fs))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}