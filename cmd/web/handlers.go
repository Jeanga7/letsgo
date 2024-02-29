package main

import (
	"errors"
	"fmt"
	"letsgo/internal/models"
	"net/http"
	"strconv"
	"text/template"
)

var files = []string{
	"/home/student/letsgo/ui/html/base.tmpl.html",
	"/home/student/letsgo/ui/html/partials/nav.tmpl.html",
	"/home/student/letsgo/ui/html/pages/home.tmpl.html",
}

func (app *Application) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	/* ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		app.ServerError(w, err)
	} */
}

func (app *Application) View(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	snippet, err := app.Snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.ServerError(w, err)
		}
		return
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	data := &TemplateData{
		Snippet: snippet,
	}
	
	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.ServerError(w, err)
	}
}

func (app *Application) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		app.ClientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n- Kobayashi Issa"
	expires := 7

	id, err := app.Snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}
