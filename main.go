package main

import (
	"os"
	"log"
	"net/http"
	"path/filepath"
	"site/helpers"
)

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

	helpers.RenderTemplate(w, "home", files)
}

func main() {
	http.HandleFunc("/view/", helpers.PostHandle)
	http.HandleFunc("/", homeHandle)
	log.Fatal(http.ListenAndServe(":8000", nil))
}