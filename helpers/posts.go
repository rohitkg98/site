package helpers

import (
	"html/template"
	"io/ioutil"
	"net/http"
)

// Global templates for caching
var base = template.Must(
	template.ParseFiles("./templates/base.html"))

var tmpls = make(map[string]*template.Template)

func init() {
	tmpls["home"] = template.Must(template.Must(base.Clone()).ParseFiles("./templates/home.html"))
	tmpls["post"] = template.Must(template.Must(base.Clone()).ParseFiles("./templates/post.html"))
	tmpls["err"] = template.Must(template.Must(base.Clone()).ParseFiles("./templates/err.html"))
}

// Global markdown pages cached
// var content = ioutil.ReadFile()

// Page will be stored in memory
type Page struct {
	Title string
	Body  []string // type implies that it is a slice
}

// Load page from storage to a struct page
func loadMarkdown(title string) (*Page, error) {
	filename := "./content/" + title + ".md"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	content := BreakAndSanitize(body)
	return &Page{Title: title, Body: content}, nil
}

// RenderTemplate : A generic template renderer
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {

	templates := tmpls[tmpl]

	err := templates.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// PostHandle for rendering posts
func PostHandle(w http.ResponseWriter, r *http.Request) {
	postName := r.URL.Path[len("/view/"):]
	toRender, err := loadMarkdown(postName)
	if err != nil {
		toRender, err = loadMarkdown("error")
	}
	RenderTemplate(w, "post", toRender)
}
