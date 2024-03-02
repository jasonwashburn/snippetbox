package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	fileserver := http.FileServer(http.Dir("./ui/static/"))

	router := httprouter.New()
	router.NotFound = http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		app.notFound(w)
	})
	router.Handler(http.MethodGet, "/static", http.StripPrefix("/static/*filepath", fileserver))
	router.HandlerFunc(http.MethodGet, "/", app.home)
	router.HandlerFunc(http.MethodGet, "/snippet/view/:id", app.snippetView)
	router.HandlerFunc(http.MethodGet, "/snippet/create", app.snippetCreate)
	router.HandlerFunc(http.MethodPost, "/snippet/create", app.snippetCreatePost)

	standard := alice.New(app.recoverPanic, app.logRequest, secureHeaders)
	return standard.Then(router)
}
