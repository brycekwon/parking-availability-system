package app

import (
	"net/http"
	"os"

	"github.com/brycekwon/parking-availability-system/website/internal/handlers"
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
	h := handlers.New(a.ctx, a.logger, a.db, a.cache)

	a.router.Handle("POST /", http.HandlerFunc(h.Update))

	// TESTING ONLY //
	a.router.Handle("GET /{lot}/{status}", http.HandlerFunc(h.InsertEvent))
	a.router.Handle("POST /C/{lot}/{status}", http.HandlerFunc(h.PostCache))
	a.router.Handle("GET /C/{lot}/{status}", http.HandlerFunc(h.GetCache))

	return nil
}
