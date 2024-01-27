package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func main() {

	app := application{
		InfoLog:  log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile),
	}

	var err error
	app.TemplateCache, err = app.newTemplateCache("ui/html")
	if err != nil {
		app.ErrorLog.Fatalln(err)
	}

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		app.InfoLog.Printf("defaulting to port %s", port)
	}

	app.RegisterRoutes()
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: app.SrvMux,
	}

	app.InfoLog.Println("Starting a server at http://localhost:" + port)
	app.ErrorLog.Fatalln(srv.ListenAndServe())

}

type NeuteredFileSystem struct {
	fs http.FileSystem
}

func (nfs NeuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if err != nil {
		return nil, err
	}

	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}

	return f, nil
}
