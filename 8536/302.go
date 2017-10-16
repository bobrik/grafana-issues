package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Location", "http://example.com/")
		w.WriteHeader(http.StatusFound)
	})

	http.ListenAndServe("127.0.0.1:80", nil)
}
