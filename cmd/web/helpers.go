package main

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"
	"runtime/debug"
)

func (app *application) newTemplateCache(dir string) (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	pages, err := filepath.Glob(filepath.Join(dir, "*.page.html"))
	if err != nil {
		return nil, err
	}

	for _, page := range pages {
		ts, err := template.ParseFiles(page)
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*layout.html"))
		if err != nil {
			return nil, err
		}

		ts, err = ts.ParseGlob(filepath.Join(dir, "*partial.html"))
		if err != nil {
			return nil, err
		}

		name := filepath.Base(page)
		templateCache[name] = ts
	}

	return templateCache, nil
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string) {
	tmpl, ok := app.TemplateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("the %s page doesn't exist", name))
		return
	}

	err := tmpl.Execute(w, nil)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) notFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
