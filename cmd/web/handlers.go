package main

import "net/http"

func (app *application) RegisterRoutes() {
	SrvMux := http.NewServeMux()
	SrvMux.HandleFunc("/", app.Home)
	app.SrvMux = SrvMux
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/":
		return
	}
}
