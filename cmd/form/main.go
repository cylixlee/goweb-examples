package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()           // parse from URL & form (application/x-www-form-urlencoded)
		fmt.Fprintln(w, r.Form) // map[string][]string, as always.

		// use r.PostForm to ignore data form URL (only access to the form's data).

		// r.ParseMultipartForm() is used to parse multipart MIME form data
		// (multipart/form-data).
		//
		// often used in uploading files.

		// use r.FormValue() / r.PostFormValue() to get the first element of a specific
		// key.
		//
		// it automatically calls r.ParseForm().
	})

	server.ListenAndServe()
}
