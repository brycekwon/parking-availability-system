package handlers

import "net/http"

func (h *Handler) InternalServerError(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	h.views.ExecuteTemplate(w, "error.html", nil)
}
