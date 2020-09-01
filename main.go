package main

import (
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Materials []Material
}

type Material struct {
	Title, Url string
}

var page = Page{
	Materials: []Material{
		{"http package", "https://golang.org/pkg/net/http/"},
		{"Writing Web Applications", "https://golang.org/doc/articles/wiki/"},
		{"Go by Example: HTTP Servers", "https://gobyexample.com/http-servers"},
		{"Hello world HTTP server example", "https://yourbasic.org/golang/http-server-example/"},
		{"How I write Go HTTP services after seven years", "https://medium.com/@matryer/how-i-write-go-http-services-after-seven-years-37c208122831"},
		{"Gorilla web toolkit", "https://www.gorillatoolkit.org/pkg/mux"},
	},
}

var templates = template.Must(template.ParseFiles("templates/materials.tpl.html"))

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	materialsRoute := "/materials"
	switch path := r.URL.Path; path {
	case "/":
		http.Redirect(w, r, materialsRoute, http.StatusPermanentRedirect)
	case materialsRoute:
		renderTemplate(w, "materials.tpl", &page)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
