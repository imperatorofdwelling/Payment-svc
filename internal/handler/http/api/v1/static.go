package v1

import (
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type staticHandler struct {
	log *zap.SugaredLogger
}

func NewStaticHandler(r chi.Router, logger *zap.SugaredLogger) {
	handler := &staticHandler{
		log: logger,
	}

	r.Route("/static", func(r chi.Router) {
		r.Route("/deals", func(r chi.Router) {
			r.Get("/save-card", handler.ServeCardSavePage)
		})
	})
}

func (h *staticHandler) ServeCardSavePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/pages/get-payout-data.html")
}
