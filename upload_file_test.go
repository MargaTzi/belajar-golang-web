package golangweb

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

//go:embed templates/tes/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/tes/*.gohtml"))

func UploadForm(w http.ResponseWriter, r *http.Request){
	myTemplates.ExecuteTemplate(w, "upload.form.gohtml", nil)
}

func Upload(w http.ResponseWriter, r *http.Request){
	file, fileHeader, err := r.FormFile("file")
	if err != nil{
		panic(err)
	}
	fileDestination, err := os.Create("./resource/" + fileHeader.Filename)
	if err != nil{
		panic(err)
	}
	_, err = io.Copy(fileDestination, file)
	if err != nil{
		panic(err)
	}

	name := r.PostFormValue("name")
	myTemplates.ExecuteTemplate(w, "upload.success.gohtml", map[string]any{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUpload(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}

//go:embed resource/larva.jpg
var uploadfile []byte

func TestUploadFile(t *testing.T){
	body := new(bytes.Buffer)

	writer := multipart.NewWriter(body)
	writer.WriteField("name", "Dito Golang")
	file, _ := writer.CreateFormFile("file", "udang.png")

	file.Write(uploadfile)
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "http://localhost:8080/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rec := httptest.NewRecorder()

	Upload(rec, req)
	bodyRes, _ := io.ReadAll(rec.Result().Body)

	fmt.Println(string(bodyRes))
}