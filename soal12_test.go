package golangweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/tes/*.gohtml
var tes12 embed.FS

var tmpl12 = template.Must(template.ParseFS(tes12, "templates/tes/*.gohtml"))

func TesUploadForm(w http.ResponseWriter, r *http.Request){
	tmpl12.ExecuteTemplate(w, "upload.test.gohtml", nil)
}

func TesUploadFile(w http.ResponseWriter, r *http.Request){
	file, fd, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file_to, err := os.Create("./resource/" + fd.Filename)
	if err != nil{
		panic(err)
	}
	defer file_to.Close() 
	_, err = io.Copy(file_to, file)
	if err != nil {
		panic(err)
	}

	tmpl12.ExecuteTemplate(w, "upload.success.gohtml", map[string]any{
		"File" : "/static/" + fd.Filename,
	})
}

func SoalDownload(w http.ResponseWriter, r *http.Request){
	file := r.URL.Query().Get("name")
	if file == ""{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(w, r, "./resource/"+file)
}

func TestSoalDownload(t *testing.T){
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(SoalDownload),
	}

	err := server.ListenAndServe()
	if err != nil{
		log.Fatalln(err)
	}
}

func (middleware *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request){
	middleware.Handler.ServeHTTP(w, r)
}

func (erorHand *ErorHandler ) ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer func ()  {
		err := recover()
		if err != nil{
			fmt.Println("Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error : %s", err)
		}
	}()
	erorHand.Handler.ServeHTTP(w,r)
}

func TestSoalMiddleware(t *testing.T){
	handler := http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprint(w, "OK DARI HANDLER")
	})

	log := &LogMiddleware{
		Handler: handler,
	}

	erormid := &ErorHandler{
		Handler: log,
	}

	req := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	rec := httptest.NewRecorder()

	erormid.ServeHTTP(rec,req)

	if rec.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rec.Code)
	}

	if rec.Body.String() != "OK DARI HANDLER" {
		t.Errorf("unexpected body: %s", rec.Body.String())
	}
}