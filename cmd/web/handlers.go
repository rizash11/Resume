package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
)

func (app *application) RegisterRoutes() {
	SrvMux := http.NewServeMux()
	SrvMux.HandleFunc("/", app.home)
	SrvMux.HandleFunc("/download/", app.download)

	SrvMux.Handle("/static/", http.StripPrefix("/static", http.FileServer(NeuteredFileSystem{http.Dir("./ui/static/")})))

	app.SrvMux = SrvMux
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path != "/":
		app.notFound(w)
		return
	}

	app.render(w, r, "home.page.html")
}

func (app *application) download(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		app.serverError(w, err)
		return
	}

	f, err := os.Open("./ui/static/img/Resume_template.pdf")
	if err != nil {
		app.serverError(w, err)
		return
	}
	defer f.Close()

	contentDisposition := fmt.Sprintf("attachment; filename=%s", path.Base(f.Name()))
	w.Header().Set("Content-Disposition", contentDisposition)

	contentType := "text/plain; charset=utf-8"
	w.Header().Set("Content-Type", contentType)

	finfo, err := f.Stat()
	if err != nil {
		app.serverError(w, err)
	}

	contentLength := strconv.Itoa(int(finfo.Size()))
	w.Header().Set("Content-Length", contentLength)

	if _, err := io.Copy(w, f); err != nil {
		app.serverError(w, err)
		return
	}
}
