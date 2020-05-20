package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"site/helpers"
	"site/middleware"
	"time"
)

func renderHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/not_found", http.StatusSeeOther)
		return
	}

	var files []string

	err := filepath.Walk("./content", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, filepath.Base(path[:len(path)-3]))
		}
		return nil
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	helpers.RenderTemplate(w, "home", files)
}

func nonExistingRoutes(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/not_found", http.StatusSeeOther)
	return
}

func notFound(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "err", "")
}

func home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	mux.HandleFunc("/not_found", notFound)
	mux.HandleFunc("/view/", helpers.PostHandle)
	mux.HandleFunc("/home/", home)
	mux.HandleFunc("/", renderHome)
	mux.HandleFunc("*", nonExistingRoutes)

	var handler http.Handler = mux
	handler = middleware.LogRequestHandler(handler)

	srv := &http.Server{
		Addr:         ":3000",
		ReadTimeout:  120 * time.Second,
		WriteTimeout: 120 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      handler,
	}
	return srv
}

func main() {
	log.Fatal(makeHTTPServer().ListenAndServe())
}
