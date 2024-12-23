package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"go.uber.org/zap"
	"net/http"
)

type payoutsHandler struct {
	svc     service.IPayoutsSvc
	cardSvc service.ICardsSvc
	log     *zap.SugaredLogger
}

func NewPayoutsHandler(r chi.Router, svc service.IPayoutsSvc, cardSvc service.ICardsSvc, log *zap.SugaredLogger) {
	handler := &payoutsHandler{svc, cardSvc, log}

	r.Route("/payouts", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			r.Post("/create", handler.createCard)
			r.Delete("/{cardId}", handler.deleteCardByID)
		})
	})
}

func (h *payoutsHandler) createCard(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payouts.CreateCard"

	var newCard model.Card

	err := json.Read(r, &newCard)
	if err != nil {
		h.log.Errorf("%s: %v", op, ErrUnmarshallingBody)
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	err = h.cardSvc.CreateBankCard(r.Context(), newCard)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.InternalApiError)
		return
	}

	json.Write(w, http.StatusCreated, newCard)
}

func (h *payoutsHandler) deleteCardByID(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payouts.DeleteCardByID"

	userId := chi.URLParam(r, "cardId")

	cardIdUUID, err := uuid.Parse(userId)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.ParseError)
		return
	}

	err = h.cardSvc.DeleteCardByID(r.Context(), cardIdUUID)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.InternalApiError)
		return
	}

	json.Write(w, http.StatusNoContent, "successfully deleted card")
}
