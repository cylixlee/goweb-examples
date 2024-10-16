package main

import (
	"fmt"
	"io"
	"log"
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
		// use r.FormValue() / r.PostFormValue() to get the first element of a specific
		// key.
		//
		// it automatically calls r.ParseForm().
	})

	http.HandleFunc("/processMultipart", func(w http.ResponseWriter, r *http.Request) {
		r.ParseMultipartForm(1024) // need to specify max memory

		// // map[string][]*multipart.FileHeader
		// //
		// // because a <input type="file"> can upload multiple files.
		// header := r.MultipartForm.File["uploaded"][0]
		// file, err := header.Open()
		// if err != nil {
		// 	log.Fatalln(err.Error())
		// }
		// defer file.Close()

		// shorter usage (not that short actually)
		file, _, err := r.FormFile("uploaded") // the first file
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			log.Fatalln(err.Error())
		}
		fmt.Fprintln(w, string(data))
	})

	server.ListenAndServe()
}
