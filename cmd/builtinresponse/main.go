package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/notfound", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	http.HandleFunc("/serveFile", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/serveContent", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("static/index.html")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer file.Close()

		http.ServeContent(w, r, "index.html", time.Now(), file)
	})

	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "https://go.dev", http.StatusFound)
	})

	server.ListenAndServe()
}
