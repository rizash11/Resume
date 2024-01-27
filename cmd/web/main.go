package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
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

	fmt.Println(app.TemplateCache)

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
