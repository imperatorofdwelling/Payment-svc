package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"github.com/imperatorofdwelling/payment-svc/pkg/yookassa"
	"go.uber.org/zap"
	"net/http"
)

type paymentsHandler struct {
	svc         service.IPaymentSvc
	log         *zap.SugaredLogger
	yookassaHdl *yookassa.PaymentsSvc
}

func NewPaymentsHandler(r chi.Router, svc service.IPaymentSvc, yookassaHdl *yookassa.PaymentsSvc, log *zap.SugaredLogger) {
	handler := &paymentsHandler{
		svc:         svc,
		log:         log,
		yookassaHdl: yookassaHdl,
	}

	r.Route("/payments", func(r chi.Router) {
		r.Group(func(r chi.Router) {
			r.Post("/", handler.createPayment)
		})
	})
}

func (h *paymentsHandler) createPayment(w http.ResponseWriter, r *http.Request) {
	const op = "handler.v1.payments.createPayment"
	var payment model.Payment

	idempotenceKey := r.Header.Get("Idempotence-Key")
	if idempotenceKey == "" {
		h.log.Errorf("%s: %v", op, ErrGettingIdempotenceKey)
		json.WriteError(w, http.StatusBadRequest, ErrGettingIdempotenceKey.Error(), json.GettingHeaderDataError)
		return
	}

	err := json.Read(r.Body, &payment)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
		return
	}

	if err := v10.Validate.Struct(payment); err != nil {
		validationErr := err.(validator.ValidationErrors)
		json.WriteError(w, http.StatusBadRequest, validationErr.Error(), json.ValidationError)
		return
	}

	paymentRes, err := h.yookassaHdl.CreatePayment(&payment, idempotenceKey)
	if err != nil {
		h.log.Errorf("%s: %v", op, err.Error())
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.ExternalApiError)
		return
	}

	var newPayment model.Payment

	err = json.Read(paymentRes.Body, &newPayment)
	if err != nil {
		h.log.Errorf("%s: %v", op, err.Error())
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.ExternalApiError)
		return
	}

	if newPayment.Status == "" {
		h.log.Error("invalid response from API", zap.String("op", op), zap.String("description", newPayment.Description))
		json.WriteError(w, http.StatusBadRequest, newPayment.Description, json.ExternalApiError)
		return
	}

	err = h.svc.CreatePayment(r.Context(), &newPayment)
	if err != nil {
		h.log.Errorf("%s: %v", op, zap.Error(err))
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.InternalApiError)
		return
	}

	json.Write(w, http.StatusOK, newPayment)
}
