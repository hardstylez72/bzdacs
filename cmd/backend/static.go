package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	path, err := filepath.Abs(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, path)

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}

func startStaticServer(staticServerHost, pathToStaticFiles string, done chan struct{}) error {
	fs := spaHandler{staticPath: pathToStaticFiles, indexPath: "index.html"}
	http.Handle("/", fs)

	serverUrl, err := url.Parse(staticServerHost)
	if err != nil {
		return err
	}

	log.Println("Listening static server on :", serverUrl.Port())

	err = http.ListenAndServe(":"+serverUrl.Port(), fs)
	if err != nil {
		return err
	}
	done <- struct{}{}
	return nil
}
