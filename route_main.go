package main

import (
	"chitchat/data"
	"net/http"

	"github.com/golang-jwt/jwt/v4/request"
)

func err(w http.ResponseWriter, r *http.Request) {
	vals := r.URL.Query()
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
	} else {
		generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(w, r, msg)
	} else {
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, vals.Get("msg"), "layout", "public.navbar", "error")
		} else {
			generateHTML(w, vals.Get("msg"), "layout", "private.navbar", "error")
		}
	}
}
