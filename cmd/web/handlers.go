package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"zametki/pkg/models"

	_ "github.com/go-sql-driver/mysql"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	s, err := app.zametkis.Get()
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}
	// Используем помощника render() для отображения шаблона.
	app.render(w, r, "home.page.tmpl", s)

}

func (app *application) showZametka(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	s, err := app.zametkis.GetOne(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	// Используем помощника render() для отображения шаблона.
	app.renderOne(w, r, "show.page.tmpl", &templateData{
		Zametki: s,
	})
}

func (app *application) DelZametka(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.notFound(w)
		return
	}
	app.zametkis.Del(id)
	http.Redirect(w, r, "/", 301)
}

func (app *application) Create(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/create" {
		app.notFound(w)
		return
	}

	app.renderPage(w, r, "create.page.tmpl")

	Title := r.FormValue("Title")
	Content := r.FormValue("Content")
	fmt.Println(Title, Content)
	if Title == "" || Content == "" {
		fmt.Printf("Заполните все поля")
	} else {
		app.zametkis.Insert(Title, Content)
	}

}
