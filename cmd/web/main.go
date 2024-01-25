package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	app := application{
		InfoLog:  log.New(os.Stdout, "INFO \t", log.Ldate|log.Ltime),
		ErrorLog: log.New(os.Stderr, "ERROR: \t", log.Ldate|log.Ltime|log.Lshortfile),
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
