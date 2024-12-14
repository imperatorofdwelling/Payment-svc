package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	yoopayment "github.com/rvinnie/yookassa-sdk-go/yookassa/payment"
	"go.uber.org/zap"
	"net/http"
)

type paymentsHandler struct {
	svc service.IPaymentSvc
	log *zap.SugaredLogger
}

func NewPaymentsHandler(r chi.Router, svc service.IPaymentSvc, log *zap.SugaredLogger) {
	handler := &paymentsHandler{
		svc: svc,
		log: log,
	}

	r.Route("/payments", func(r chi.Router) {
		r.Post("/", handler.createPayment)
	})

}

func (h *paymentsHandler) createPayment(w http.ResponseWriter, r *http.Request) {
	var payment yoopayment.Payment

	err := json.Read(r, &payment)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err.Error(), json.DecodeBodyError)
	}

	json.Write(w, http.StatusOK, payment)

}

//v := h.svc.GetSTest()

//if err := v10.Validate.Struct(tt); err != nil {
//	validationErr := err.(validator.ValidationErrors)
//	json.WriteError(w, http.StatusBadRequest, validationErr.Error(), json.ValidationError)
//	return
//}

//json.Write(w, http.StatusOK, tt)
