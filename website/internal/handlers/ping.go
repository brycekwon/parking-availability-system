package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/brycekwon/parking-availability-system/website/internal/models"
)

// TESTING PURPOSES ONLY //
func (h *Handler) InsertEvent(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("lot")
	if len(name) <= 0 {
		h.logger.Error("invalid lot name provided")
		return
	}

	status, err := strconv.Atoi(r.PathValue("status"))
	if err != nil {
		h.logger.Error("invalid status provided")
		return
	} else if status != 0 && status != 1 {
		h.logger.Error("invalid status provided2")
		return
	}

	err = models.InsertEvent(h.db, name, 2, status)
	if err != nil {
		h.logger.Error("Failed to insert event to database", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) PostCache(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("lot")
	if len(name) <= 0 {
		h.logger.Error("invalid lot name provided")
		return
	}

	status, err := strconv.Atoi(r.PathValue("status"))
	if err != nil {
		h.logger.Error("invalid status provided")
		return
	} else if status != 0 && status != 1 {
		h.logger.Error("invalid status provided2")
		return
	}

	err = h.cache.Set(h.ctx, name, status, 0).Err()
	if err != nil {
		h.logger.Error("failed to insert event into cache", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (h *Handler) GetCache(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("lot")
	if len(name) <= 0 {
		h.logger.Error("invalid lot name provided")
		return
	}

	status, err := strconv.Atoi(r.PathValue("status"))
	if err != nil {
		h.logger.Error("invalid status provided")
		return
	} else if status != 0 && status != 1 {
		h.logger.Error("invalid status provided2")
		return
	}

	val, err := h.cache.Get(h.ctx, name).Result()
	if err != nil {
		h.logger.Error("failed to fetch event from cache", slog.Any("error", err))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	h.logger.Info("fetched", slog.String("val", val))
}
