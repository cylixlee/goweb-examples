package main

import "net/http"

type myHandler struct{}

func (m *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("No matter the URL, hello!"))
}

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: &myHandler{},
	}
	server.ListenAndServe()
}
