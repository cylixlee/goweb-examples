package main

import (
	"fmt"
	"io"
	"net/http"
)

// URL: scheme://[userinfo@]host/path[?query][#fragment]
//   or scheme:opaque[?query][#fragment]

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/url", func(w http.ResponseWriter, r *http.Request) {
		// fragment
		fmt.Fprintln(w, r.URL.Fragment) // probably filtered by browser
		// header
		fmt.Fprintln(w, r.Header)                        // map[string][]string
		fmt.Fprintln(w, r.Header["Accept-Encoding"])     // []string
		fmt.Fprintln(w, r.Header.Get("Accept-Encoding")) // string
	})

	http.HandleFunc("/post", func(w http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		defer r.Body.Close() // NOTE: use defer to avoid memory leak
		fmt.Fprintln(w, "read: ", string(data))
	})

	http.HandleFunc("/query", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.URL.RawQuery)
		query := r.URL.Query() // map[string][]string because query can carry multiple values.
		id := query["id"]
		fmt.Fprintln(w, id)                // the []string
		fmt.Fprintln(w, query.Get("name")) // the first string in that slice
	})

	server.ListenAndServe()
}
