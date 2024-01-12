package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Manga struct {
	Title  string
	Author string
}

func main() {
	fmt.Println("WIP Go app...")

	// handler function #1 - returns the index.html template, with dummy data
	h1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("src/index.html"))
		manga := map[string][]Manga{
			"Manga": {
				{Title: "Blah", Author: "Mr Blah"},
				{Title: "Blah II", Author: "Mr Blah Blah"},
				{Title: "Dummy Manga", Author: "Dummy Author"},
			},
		}
		tmpl.Execute(w, manga)
	}

	// handler function #2 - returns the template block with the newly added data, as an HTMX response
	h2 := func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1 * time.Second)
		title := r.PostFormValue("title")
		author := r.PostFormValue("author")
		tmpl := template.Must(template.ParseFiles("src/index.html"))
		tmpl.ExecuteTemplate(w, "manga-list-element", Manga{Title: title, Author: author})
	}
	// define handlers
	http.HandleFunc("/", h1)
	http.HandleFunc("/add-manga/", h2)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
