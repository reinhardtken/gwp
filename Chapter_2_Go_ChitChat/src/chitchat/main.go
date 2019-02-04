package main

import (
	//"chitchat/sub"
	"net/http"
	"time"
	"chitchat/enter"

)

func init() {
	enter.Test1()
}

func main() {
	enter.P("ChitChat", enter.Version(), "started at", enter.Config.Address)

	// handle static assets
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(enter.Config.Static))
	mux.Handle("/static/", http.StripPrefix("/static/", files))

	//
	// all route patterns matched here
	// route handler functions defined in other files
	//

	// index
	mux.HandleFunc("/", enter.Index)
	// error
	mux.HandleFunc("/err", enter.Err)

	// defined in route_auth.go
	mux.HandleFunc("/login", enter.Login)
	mux.HandleFunc("/logout", enter.Logout)
	mux.HandleFunc("/signup", enter.Signup)
	mux.HandleFunc("/signup_account", enter.SignupAccount)
	mux.HandleFunc("/authenticate", enter.Authenticate)

	// defined in route_thread.go
	mux.HandleFunc("/thread/new", enter.NewThread)
	mux.HandleFunc("/thread/create", enter.CreateThread)
	mux.HandleFunc("/thread/post", enter.PostThread)
	mux.HandleFunc("/thread/read", enter.ReadThread)

	// starting up the server
	server := &http.Server{
		Addr:           enter.Config.Address,
		Handler:        mux,
		ReadTimeout:    time.Duration(enter.Config.ReadTimeout * int64(time.Second)),
		WriteTimeout:   time.Duration(enter.Config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes: 1 << 20,
	}
	server.ListenAndServe()
}
