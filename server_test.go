package golangweb

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080", 
	}
	fmt.Println("server starting")
	
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}