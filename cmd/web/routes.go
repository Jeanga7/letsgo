package main

import "net/http"

func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/snippet/view", app.View)
	mux.HandleFunc("/snippet/create", app.Create)
	return mux
}
