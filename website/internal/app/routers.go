package app

import (
	"html/template"
	"io/fs"
	"net/http"

	"github.com/brycekwon/parking-availability-system/website/internal/handlers"
)

func (a *App) loadAssets() error {
	content, err := fs.Sub(a.assets, "dist")
	if err != nil {
		return err
	}
	a.router.Handle("GET /dist/", http.StripPrefix("/dist/", http.FileServer(http.FS(content))))

	return nil
}

func (a *App) loadPages() error {
	pages := template.Must(template.New("").ParseFS(a.pages, "templates/*"))

	h := handlers.New(a.logger, pages)

	a.router.Handle("GET /{$}", http.HandlerFunc(h.Home))
	a.router.Handle("POST /update", http.HandlerFunc(h.Update))

	return nil
}
