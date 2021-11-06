package handlers

import "html/template"

var tpl *template.Template

func LoadTemplates() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}
