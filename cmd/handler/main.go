package main

import "net/http"

type helloHandler struct{}

func (h *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello, world!"))
}

type aboutHandler struct{}

func (a *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("about this..."))
}

func simple(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("simple!"))
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/hello", &helloHandler{})
	http.Handle("/about", &aboutHandler{})

	http.Handle("/simple", http.HandlerFunc(simple))

	http.HandleFunc("/simplest", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("simplest!"))
	})
	server.ListenAndServe()
}
