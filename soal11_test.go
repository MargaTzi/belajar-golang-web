package golangweb

import (
	"embed"
	"fmt"
	// "fmt"
	// "io"
	"net/http"
	// "net/http/httptest"
	"html/template"
	"testing"
)

//go:embed templates/tes/*.gohtml
var testtmpl embed.FS

var SoalTes = template.Must(template.ParseFS(testtmpl, "templates/tes/*.gohtml"))

// func SoalCaching(w http.ResponseWriter, r *http.Request){
// 	SoalTes.ExecuteTemplate(w, "home.gohtml", map[string]any{
// 		"Name": "Home Page",
// 	})
// }

// func TestCachingSoal(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)
// 	rec := httptest.NewRecorder()

// 	SoalCaching(rec, req)
// 	res := rec.Result()
// 	body, _ := io.ReadAll(res.Body) 

// 	fmt.Println(string(body))
// }

func SoalXssRaw(w http.ResponseWriter, r *http.Request){
	msg := r.URL.Query().Get("msg")
	err := SoalTes.ExecuteTemplate(w, "xss.gohtml", map[string]any{
		"Message": template.HTML(msg),
	})

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SoalXssRawsafe(w http.ResponseWriter, r *http.Request){
	msg := r.URL.Query().Get("msg")
	err := SoalTes.ExecuteTemplate(w, "xss.gohtml", map[string]any{
		"Message": msg,
	})

	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func TestXssRaw(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/comment/raw", SoalXssRaw)
	mux.HandleFunc("/comment/safe", SoalXssRawsafe)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err:= server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}

func SoalEscape(w http.ResponseWriter, r *http.Request){
	admin := r.URL.Query().Get("admin")
	user := r.URL.Query().Get("user")

	err := SoalTes.ExecuteTemplate(w, "escape.gohtml", map[string]any{
		"Admin": template.HTML(admin),
		"User": user,
	})

	if err != nil{
		panic(err)
	}
}

func TestEscape(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/escape", SoalEscape)
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err:= server.ListenAndServe()
	if err !=nil{
		panic(err)
	}
}

func SoalRedirect(w http.ResponseWriter, r *http.Request){
	next := r.URL.Query().Get("next")
	target := "/"

	if next == "/dashboard"{
		target = next
	}
	if next == "/profile"{
		target = next
	}

	http.Redirect(w, r, target, http.StatusFound)
}

func TestSoalRedirect(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/login", SoalRedirect)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func SoalM(w http.ResponseWriter, r *http.Request){
	to := r.URL.Query().Get("to")
	target := "/"

	if to == "/dashboard"{
		target = to
	} 
	if to == "/profile"{
		target = to
	}

	http.Redirect(w, r, target, http.StatusFound)
}


func TestSoalM(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/go", SoalM)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func Jump(w http.ResponseWriter, r *http.Request) {
	dest := r.URL.Query().Get("dest")
	target := "/"

	if dest == "home" {
		target = "/"
	} else if dest == "about" {
		target = "/about"
	} else if dest == "kontak" {
		target = "/kontak"
	}

	http.Redirect(w, r, target, http.StatusFound)
}

func Rumah(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About Page")
}

func Kontak(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Kontak Page")
}


func TestJump(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/jump", Jump)
	// target pages
	mux.HandleFunc("/", Rumah)
	mux.HandleFunc("/about", About)
	mux.HandleFunc("/kontak", Kontak)


	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}