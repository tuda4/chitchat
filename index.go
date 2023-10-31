package main

import (
	"chitchat/data"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err == nil {
		public_tmpl_files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}
		private_tmpl_files := []string{
			"templates/layout.html",
			"templates/navbar.html",
			"templates/index.html",
		}

		var templates *template.Template
		_, err := session(w, r)
		if err != nil {
			templates = template.Must(template.ParseFiles(private_tmpl_files...))
		} else {
			templates = template.Must(template.ParseFiles(public_tmpl_files...))
		}

		templates.ExecuteTemplate(w, "layout", threads)
	}
}
