package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request){
	file := r.URL.Query().Get("name")
	if file == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Ora iso")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./resource/"+file)
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}