package testing

import (
	"log"
	"net/http"
	"testing"
)

func TesServeFile(w http.ResponseWriter, r *http.Request){
	if r.URL.Query().Get("name") != "" {
		http.ServeFile(w, r, "../resource/ok.html")
	} else {
		http.ServeFile(w, r, "../resource/notFound.html")
	}
}

func TestSoalServeFile(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TesServeFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}