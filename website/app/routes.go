package app

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/brycekwon/parking-availability-system/website/handlers"
)

func (a *App) loadStatic() error {
	content, err := fs.Sub(a.assets, "static")
	if err != nil {
		return err
	}

	f := http.FileServerFS(content)

	a.router.Handle("GET /favicon.ico", f)
	a.router.Handle("GET /robots.txt", f)
	a.router.Handle("GET /css/", f)
	a.router.Handle("GET /js/", f)

	return nil
}

func (a *App) loadViews() error {
	views, err := template.New("").ParseFS(a.views, "views/*")
	if err != nil {
		return err
	}

	h := handlers.New(a.logger, views, a.db)

	a.router.Handle("GET /{$}", http.HandlerFunc(h.ViewHome))
	a.router.Handle("GET /", http.HandlerFunc(h.PageNotFound))

	a.router.Handle("POST /{lot}/{status}", http.HandlerFunc(h.InsertEvent))

	return nil
}
