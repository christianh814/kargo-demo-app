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

// Greet Struct
type Greet struct {
	Msg string
}

// GreetHtml is the template for the /greet route
const GreetHtml string = "html/greet.tmpl"

// App configuration file location
const GreetConfigFile string = "config/greet.json"

/*
// appGreet is the route that returns the pod information
func appGreet(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(GreetHtml))
	msg, err := utils.LoadConfig(GreetConfigFile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Internal Error occurred")
		return
	}

	config := Greet{
		Msg: msg,
	}

	// Display index page from template
	tmpl.Execute(w, config)
}
*/
