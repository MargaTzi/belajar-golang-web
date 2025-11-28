package testing

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func Test1(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo, ini method: %s\n", r.Method)
		fmt.Fprintf(w, "Dan path: %s\n", r.RequestURI)
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}
}

func Test2(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Warung")
	})
	mux.HandleFunc("/produk", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "nasi rames\n")
		fmt.Fprint(w, "nasi pecel")
	})
	mux.HandleFunc("/produk/meja", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "meja\n")
		fmt.Fprint(w, "kamar mandi")
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

func Test3(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}
}

func Test4(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Selamat datang di website kami !!!")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini halaman about")
	})
	mux.HandleFunc("/hello/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo, Kamu sedang membuka %s", r.RequestURI)
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}
}

func Test5(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method{
		case "GET":
			fmt.Fprint(w, "Silahkan isi form logim")
		case "POST":
			fmt.Fprint(w, "Berhasil login")
		default:
			fmt.Fprint(w, "Method tidak diizinkan")
		}
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