package main

import "net/http"

func (app *application) RegisterRoutes() {
	SrvMux := http.NewServeMux()
	SrvMux.HandleFunc("/", app.Home)

	SrvMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(NeuteredFileSystem{http.Dir("./ui/static/")})))

	app.SrvMux = SrvMux
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/":
		app.notFound(w)
		return
	}

	app.render(w, r, "home.page.html")
}
