package main

import (
	"os"
	"log"
	"net/http"
	"io/ioutil"
	"html/template"
	"path/filepath"
)

// Global templates for caching
var templates = template.Must(template.ParseFiles(
	"templates/home.html",
))

// How a Page will be stored in memory
type Page struct {
	Title string
	Body []byte // type implies that it is a slice
}

// Load page from storage to a struct page
func loadPage(title string) (*Page, error) {
	filename := "./content/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body:body}, nil
}

// A generic template renderer
func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func homeHandle(w http.ResponseWriter, r *http.Request) {
	var files []string

	err := filepath.Walk("./content", func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			files = append(files, filepath.Base(path[:len(path)-3]))
		}
		return nil
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	renderTemplate(w, "home", files)
}

func main() {
	http.HandleFunc("/", homeHandle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}