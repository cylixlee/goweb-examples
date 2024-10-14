package main

import (
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func longOperation(w http.ResponseWriter, r *http.Request) {
	for {
		time.Sleep(time.Second * 1)
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.Handle("/notfound", http.NotFoundHandler())                                 // 404
	http.Handle("/redirect", http.RedirectHandler("/notfound", http.StatusNotFound)) // redirect
	http.Handle("/removethis/hello", http.StripPrefix("/removethis", http.HandlerFunc(hello)))
	http.Handle(
		"/longoperation",
		http.TimeoutHandler(
			http.HandlerFunc(longOperation),
			time.Second*2,
			"too slow",
		),
	)
	http.Handle("/", http.FileServer(http.Dir("static")))
	// another usage:
	// http.ListenAndServe(":8080", http.FileServer(...))
	//
	// or:
	// http.Handle("/fs", http.StripPrefix("/fs", http.FileServer(...)))

	server.ListenAndServe()
}
