package app

import (
	"html/template"
	"net/http"
)

// AppConfig type
type AppConfig struct {
	Status string
}

// IndexHtml is the template for the / route
const IndexHtml string = "html/index.tmpl"

// appRoot is the HTTP root of the application
func appRoot(w http.ResponseWriter, r *http.Request) {
	// Parse the config file (it's empty/static for the / route)
	tmpl := template.Must(template.ParseFiles(IndexHtml))
	config := AppConfig{}

	// Display index page from template - static for now
	tmpl.Execute(w, config)
}
