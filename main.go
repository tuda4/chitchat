package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Statis))
	mux.Handle("/static", http.StripPrefix("/static", files))
	fmt.Println(config)
	mux.HandleFunc("/", index)

	mux.HandleFunc("/err", err)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/authenticate", authenticate)

	server := &http.Server{
		Addr:           config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Errorf("server is not connect %v", err)
	}

	fmt.Println("server is connected")

}
