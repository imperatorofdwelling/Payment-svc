package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"github.com/rvinnie/yookassa-sdk-go/yookassa"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
	"net/http"
)

type paymentsHandler struct {
	svc         service.IPaymentSvc
	log         *zap.SugaredLogger
	yookassaHdl *yookassa.PaymentHandler
}

func NewPaymentsHandler(r chi.Router, svc service.IPaymentSvc, yookassaHdl *yookassa.PaymentHandler, log *zap.SugaredLogger) {
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
	const op = "hanlder.v1.payments.createPayment"
	var payment yoopayment.Payment

	idempotenceKey := r.Header.Get("Idempotence-Key")
	if idempotenceKey == "" {
		h.log.Errorf("%s: %v", op, ErrGettingIdempotenceKey)
		json.WriteError(w, http.StatusBadRequest, ErrGettingIdempotenceKey.Error(), json.GettingHeaderDataError)
	}

	err := json.Read(r, &payment)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
	}

	newPayment, err := h.yookassaHdl.WithIdempotencyKey(idempotenceKey).CreatePayment(&payment)
	if err != nil {
		json.WriteError(w, http.StatusInternalServerError, err.Error(), json.ExternalApiError)
	}

	//createdPayment := h.svc.CreatePayment(&payment)

	json.Write(w, http.StatusOK, newPayment)

}

//v := h.svc.GetSTest()

//if err := v10.Validate.Struct(tt); err != nil {
//	validationErr := err.(validator.ValidationErrors)
//	json.WriteError(w, http.StatusBadRequest, validationErr.Error(), json.ValidationError)
//	return
//}

//json.Write(w, http.StatusOK, tt)
