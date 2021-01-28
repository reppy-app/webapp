package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// LoadTemplates put html inside variable templates.
func LoadTemplates() {
	templates = template.Must(template.ParseGlob("views/*.html"))
}

// ExecuteTemplate render html on screen.
func ExecuteTemplate(w http.ResponseWriter, template string, data interface{}) {
	templates.ExecuteTemplate(w, template, data)
}
