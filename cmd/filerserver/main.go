package main

import (
	"log"
	"net/http"
	"strings"
)

func main() {
	fs := http.FileServer(http.Dir("/home/hs/go/src/github.com/hardstylez72/bblog/ad/ui/dist"))
	http.Handle("/", fs)

	go func() {
		log.Println("Listening on :3333...")
		err := http.ListenAndServe(":3333", fs)
		if err != nil {
			log.Fatal(err)
		}
	}()

	f := http.NewServeMux()
	f.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/api") {
			http.Redirect(w, r, "http://localhost:3000"+r.URL.Path, http.StatusFound)
		} else {
			http.Redirect(w, r, "http://localhost:3333"+r.URL.Path, http.StatusFound)
		}
	})

	log.Println("Listening on :8080...")
	err := http.ListenAndServe(":8080", f)
	if err != nil {
		log.Fatal(err)
	}
}
