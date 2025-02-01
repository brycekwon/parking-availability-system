package handlers

import "net/http"

func (h *Handler) PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	h.views.ExecuteTemplate(w, "error.html", nil)
}
