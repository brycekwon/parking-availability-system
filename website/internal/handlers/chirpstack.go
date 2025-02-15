package handlers

import (
	"encoding/hex"
	"io"
	"log/slog"
	"net/http"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"

	"github.com/chirpstack/chirpstack/api/go/v4/integration"
)

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		h.logger.Error("failed to ready request body", slog.Any("error", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	h.logger.Info("Dump", slog.Any("json", r.URL.Query()))

	event := r.URL.Query().Get("event")
	switch event {
	case "up":
		err = h.uplink_event(b)
	case "join":
		err = h.join_event(b)
	default:
		h.logger.Error("handler for event is not implemented", slog.String("event", event))
		return
	}

	if err != nil {
		h.logger.Error("handled event", slog.String("event", event))
	}
}

func (h *Handler) uplink_event(b []byte) error {
	var up integration.UplinkEvent
	if err := h.unmarshal(b, &up); err != nil {
		return err
	}
	h.logger.Info("Uplink message received", slog.String("device", up.GetDeviceInfo().DevEui), slog.String("payload", hex.EncodeToString(up.Data)))
	return nil
}

func (h *Handler) join_event(b []byte) error {
	var join integration.JoinEvent
	if err := h.unmarshal(b, &join); err != nil {
		return err
	}
	h.logger.Info("Join message received", slog.String("device", join.GetDeviceInfo().DevEui), slog.String("address", join.DevAddr))
	return nil
}

func (h *Handler) unmarshal(b []byte, v proto.Message) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: true,
		AllowPartial:   true,
	}.Unmarshal(b, v)
}
