package main

import (
	"html/template"
	"log"
	"net/http"
)

type application struct {
	InfoLog       *log.Logger
	ErrorLog      *log.Logger
	SrvMux        *http.ServeMux
	TemplateCache map[string]*template.Template
}
