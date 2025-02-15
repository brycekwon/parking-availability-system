package app

import (
	"net/http"
	"os"

	"github.com/brycekwon/parking-availability-system/website/internal/handler"
)

func (a *App) loadPages() error {
	pages := http.FileServer(http.Dir("./public"))

	a.router.Handle("GET /", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		filePath := "./public" + r.URL.Path
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			http.ServeFile(w, r, "./public/404.html")
		} else {
			pages.ServeHTTP(w, r)
		}
	}))

	return nil
}

func (a *App) loadRoutes() error {
	h := handler.New(a.logger, a.db)

	a.router.Handle("POST /", http.HandlerFunc(h.Update))
	// a.router.Handle("POST /{lot}/{status}", http.HandlerFunc(h.InsertEvent))

	return nil
}
