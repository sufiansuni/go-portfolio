package handlers

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func LoadTemplates() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func IndexPageHandler(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.tmpl.html", nil)
}
