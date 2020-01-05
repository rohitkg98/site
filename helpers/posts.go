package helpers

import (
	"net/http"
	"io/ioutil"
	"html/template"
)

// Global templates for caching
var templates = template.Must(template.ParseGlob(
	"templates/*.html",
))

// // Global markdown pages cached
// var content = ioutil.ReadFile()

// How a Page will be stored in memory
type Page struct {
	Title string
	Body []string // type implies that it is a slice
}

// Load page from storage to a struct page
func loadPage(title string) (*Page, error) {
	filename := "./content/" + title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := BreakAndSanitize(body)
	return &Page{Title: title, Body: content}, nil
}

// A generic template renderer
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    err := templates.ExecuteTemplate(w, tmpl + ".html", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func PostHandle(w http.ResponseWriter, r *http.Request) {
	post_name := r.URL.Path[len("/view/"):]
	to_render, err := loadPage(post_name)
	if err != nil {
		panic(err)
	}
	RenderTemplate(w, "post", to_render)
}
