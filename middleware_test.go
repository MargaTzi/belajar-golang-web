package golangweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

// func (mid *LogMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("PATH:", r.URL.Path)
// 	fmt.Println("Sebelum eksekusi handler")
// 	mid.Handler.ServeHTTP(w, r)
// 	fmt.Println("Sudah Eksekusi handler")
// }

type ErorHandler struct{
	Handler http.Handler
}

// func (errorHand *ErorHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
// 	defer func ()  {
// 		err := recover()
// 		if err != nil {
// 			fmt.Println("Error COk")
// 			w.WriteHeader(http.StatusInternalServerError)
// 			fmt.Fprintf(w, "Error : %s", err)
// 		}
// 	}()
// 	errorHand.Handler.ServeHTTP(w, r)
// }

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("Handler Di eksekusi")
		fmt.Fprint(w, "Halo Middleware !!")
	})
	mux.HandleFunc("/halo", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("halo Di eksekusi")
		fmt.Fprint(w, "Halo Dunia !!")
	})
	mux.HandleFunc("/panic", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Println("panic exceuted")
		panic("Ups Eror")
	})

	log := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErorHandler{
		Handler: log,
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	err := server.ListenAndServe()
	if err != nil{
		panic(err)
	}
}