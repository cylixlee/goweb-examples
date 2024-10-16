package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.HandleFunc("/write", func(w http.ResponseWriter, r *http.Request) {
		// implicitly calls WriteHeader, and first 512 bytes of data is used to decide the
		// Content-Type.
		w.Write([]byte("<html><body><b>HelloWorld</b></body></html>"))
	})
	http.HandleFunc("/writeHeader", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented) // implicitly called when first Write.
		fmt.Fprintln(w, "No such service, try next door.")
	})
	http.HandleFunc("/redirect", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Location", "https://go.dev") // the response header can be modified manually
		w.WriteHeader(http.StatusFound)              // before writing the header.
	})
	server.ListenAndServe()
}
