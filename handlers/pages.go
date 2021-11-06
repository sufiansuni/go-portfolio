package handlers

import (
	"net/http"
)

func IndexPageHandler(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "index.tmpl.html", nil)
}

func AboutPageHandler(res http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(res, "about.tmpl.html", nil)
}
