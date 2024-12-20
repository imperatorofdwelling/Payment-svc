package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"go.uber.org/zap"
	"net/http"
)

type cardsHandler struct {
	svc service.ICardsSvc
	log *zap.SugaredLogger
}

func NewCardsHandler(r chi.Router, svc service.ICardsSvc, log *zap.SugaredLogger) {
	handler := &cardsHandler{svc: svc, log: log}

	r.Route("/cards", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/create", handler.CreateCard)
		})
	})
}

func (h *cardsHandler) CreateCard(w http.ResponseWriter, r *http.Request) {
	const op = "handler.cards.CreateCard"

	var newCard model.Card

	err := json.Read(r, &newCard)
	if err != nil {
		h.log.Errorf("%s: %v", op, ErrUnmarshallingBody)
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	err = h.svc.CreateBankCard(r.Context(), newCard)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.InternalApiError)
		return
	}

	json.Write(w, http.StatusCreated, "successfully created card")
}
