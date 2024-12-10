package v1

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"go.uber.org/zap"
	"net/http"
)

type paymentHandler struct {
	svc service.IPaymentSvc
	log *zap.SugaredLogger
}

var errs error

type TestStruct struct {
	Value string `json:"value" validate:"required,currency"`
}

func NewPaymentHandler(r chi.Router, svc service.IPaymentSvc, log *zap.SugaredLogger) {
	handler := &paymentHandler{
		svc: svc,
		log: log,
	}

	r.Route("/payment", func(r chi.Router) {
		r.Get("/", handler.getTest)
	})

}

func (h *paymentHandler) getTest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//v := h.svc.GetSTest()

	tt := &TestStruct{
		Value: "13f.92",
	}

	if err := v10.Validate.Struct(tt); err != nil {
		validationErr := err.(validator.ValidationErrors)
		json.WriteError(w, http.StatusBadRequest, validationErr.Error(), json.ValidationError)
		return
	}

	json.Write(w, http.StatusOK, tt)
}
