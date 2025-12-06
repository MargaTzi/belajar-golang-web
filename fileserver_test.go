package golangweb

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	dir := http.Dir("./resource")
	fs := http.FileServer(dir)

	mux := http.NewServeMux()
	mux.Handle("/statis/", http.StripPrefix("/statis", fs))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

//go:embed resource
var resource embed.FS

func TestFileServerEmbed(t *testing.T) {
	dir, _ := fs.Sub(resource, "resource")
	fs := http.FileServer(http.FS(dir))

	mux := http.NewServeMux()
	mux.Handle("/statis/", http.StripPrefix("/statis", fs))

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}