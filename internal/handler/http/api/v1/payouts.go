package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"github.com/imperatorofdwelling/payment-svc/pkg/yookassa"
	"go.uber.org/zap"
	"net/http"
)

type payoutsHandler struct {
	svc               service.IPayoutsSvc
	cardSvc           service.ICardsSvc
	yookassaPayoutHdl *yookassa.PayoutsHandler
	logsSvc           service.ILogsSvc
	log               *zap.SugaredLogger
}

func NewPayoutsHandler(r chi.Router, svc service.IPayoutsSvc, cardSvc service.ICardsSvc, yookassaPayoutHdl *yookassa.PayoutsHandler, logsSvc service.ILogsSvc, log *zap.SugaredLogger) {
	handler := &payoutsHandler{svc, cardSvc, yookassaPayoutHdl, logsSvc, log}

	r.Route("/payouts", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			r.Post("/create", handler.createCard)
			r.Delete("/{cardId}", handler.deleteCardByID)
		})

		r.Post("/new", handler.newPayout)
		r.Get("/{payoutId}", handler.getPayoutInfo)
	})
}

func (h *payoutsHandler) createCard(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payouts.createCard"

	var newCard model.Card

	err := json.Read(r.Body, &newCard)
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
	const op = "handler.v1.payouts.deleteCardByID"

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

func (h *payoutsHandler) getPayoutInfo(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payouts.getPayoutInfo"

	payoutID := chi.URLParam(r, "payoutId")

	res, err := h.yookassaPayoutHdl.GetPayoutInfo(payoutID)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.ExternalApiError)
		return
	}

	var payoutInfo model.Payout

	err = json.Read(res.Body, &payoutInfo)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	json.Write(w, http.StatusOK, payoutInfo)

}

func (h *payoutsHandler) newPayout(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payouts.newPayout"

	idempotenceKey := r.Header.Get("Idempotence-Key")
	if idempotenceKey == "" {
		h.log.Errorf("%s: %v", op, ErrGettingIdempotenceKey)
		json.WriteError(w, http.StatusBadRequest, ErrGettingIdempotenceKey.Error(), json.GettingHeaderDataError)
		return
	}

	var newPayout model.Payout

	err := json.Read(r.Body, &newPayout)
	if err != nil {
		h.log.Errorf("%s: %v", op, ErrUnmarshallingBody)
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	if err := v10.Validate.Struct(newPayout); err != nil {
		validationErr := err.(validator.ValidationErrors)
		json.WriteError(w, http.StatusBadRequest, validationErr.Error(), json.ValidationError)
		return
	}

	payoutRes, err := h.yookassaPayoutHdl.MakePayout(&newPayout, idempotenceKey)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.ExternalApiError)
		return
	}

	var createdPayout model.Payout

	err = json.Read(payoutRes.Body, &createdPayout)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.DecodeBodyError)
		return
	}

	newLog := &model.Log{
		TransactionID:   createdPayout.ID,
		TransactionType: model.PayoutType,
		Status:          *createdPayout.Status,
		Value:           createdPayout.Amount.Value,
		Currency:        createdPayout.Amount.Currency,
	}

	err = h.logsSvc.InsertLog(r.Context(), newLog)
	if err != nil {
		h.log.Errorf("%s: %v", op, err)
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.InternalApiError)
		return
	}

	json.Write(w, http.StatusOK, newPayout)
}
