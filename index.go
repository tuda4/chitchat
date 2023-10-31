package main

import (
	"chitchat/data"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		return
	}

	_, err = session(w, r)
	if err != nil {
		generateHTML(w, threads, "layout", "public.navbar", "index")
		return
	}
	generateHTML(w, threads, "layout", "private.navbar", "index")
}
