package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.Home)
	mux.HandleFunc("/show", app.showZametka)
	mux.HandleFunc("/del", app.DelZametka)
	mux.HandleFunc("/create", app.Create)

	fileServer := http.FileServer(http.Dir("ui/html/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return mux
}
