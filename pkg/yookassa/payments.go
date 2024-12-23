package yookassa

import (
	"encoding/json"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"net/http"
)

type PaymentsHandler struct {
	*Client
	log *zap.SugaredLogger
}

func NewPaymentsHandler(client *Client, log *zap.SugaredLogger) *PaymentsHandler {
	return &PaymentsHandler{
		client,
		log,
	}
}

func (h *PaymentsHandler) CreatePayment(payment *model.Payment, idempotencyKey string) (*http.Response, error) {
	pJSON, err := json.Marshal(payment)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payment json: %w", err)
	}

	res, err := h.makeRequest(http.MethodPost, PaymentEndpoint, "", pJSON, nil, idempotencyKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}
