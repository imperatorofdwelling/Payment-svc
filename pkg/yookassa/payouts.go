package yookassa

import (
	"encoding/json"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"net/http"
)

type PayoutsHandler struct {
	*Client
	log *zap.SugaredLogger
}

func NewPayoutsHandler(client *Client, log *zap.SugaredLogger) *PayoutsHandler {
	return &PayoutsHandler{
		client, log,
	}
}

func (h *PayoutsHandler) MakePayout(payout *model.Payout, idempotencyKey string) (*http.Response, error) {
	jsonData, err := json.Marshal(payout)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payout: %s", err)
	}

	res, err := h.makeRequest(http.MethodPost, PayoutEndpoint, "", jsonData, nil, idempotencyKey)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *PayoutsHandler) GetPayoutInfo(id string) (*http.Response, error) {
	res, err := h.makeRequest(http.MethodGet, PayoutEndpoint, id, nil, nil, "")
	if err != nil {
		return nil, err
	}

	return res, nil
}
