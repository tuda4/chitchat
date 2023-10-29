package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("/"))
	mux.Handle("/static", http.StripPrefix("/static", files))

	mux.HandleFunc("/", index)

	server := &http.Server{
		Addr:    "http://localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Errorf("server is not connect %v", err)
	}

	fmt.Println("server is connected")
}
