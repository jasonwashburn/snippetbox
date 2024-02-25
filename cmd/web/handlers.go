package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jasonwashburn/snippetbox/internal/models"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	snippets, err := app.snipppets.Latest()
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.render(w, r, http.StatusOK, "home.tmpl.html", templateData{Snippets: snippets})
}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snipppets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	app.render(w, r, http.StatusOK, "view.tmpl.html", templateData{Snippet: snippet})
}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snipppets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, r, err)
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
