package handlers

import "net/http"

func (h *Handler) ViewHome(w http.ResponseWriter, r *http.Request) {
	h.views.ExecuteTemplate(w, "index.html", nil)
}
