package yookassa

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"github.com/imperatorofdwelling/payment-svc/internal/domain/model"
	"go.uber.org/zap"
	"net/http"
)

type paymentsHandler struct {
	log    *zap.SugaredLogger
	client *http.Client
	cfg    config.PayApi
}

func newPaymentsHandler(cfgPayApi config.PayApi, logger *zap.SugaredLogger) *paymentsHandler {
	return &paymentsHandler{
		log: logger,
		client: &http.Client{
			Transport: LoggingRoundTripper{
				Proxied: http.DefaultTransport,
				cfg:     cfgPayApi,
			},
		},
		cfg: cfgPayApi,
	}
}

func (h *paymentsHandler) CreatePayment(payment model.PaymentReq) (*model.PaymentRes, error) {
	pJSON, err := json.Marshal(payment)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payment json: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, h.cfg.Addr, bytes.NewBuffer(pJSON))
	if err != nil {
		return nil, fmt.Errorf("create payment request: %w", err)
	}

	res, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("create payment request: %w", err)
	}

	defer res.Body.Close()

	var paymentRes model.PaymentRes

	err = json.NewDecoder(res.Body).Decode(&paymentRes)
	if err != nil {
		return nil, fmt.Errorf("create payment response: %w", err)
	}

	return &paymentRes, nil
}

type LoggingRoundTripper struct {
	Proxied http.RoundTripper
	cfg     config.PayApi
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%d:%s", lrt.cfg.ShopID, lrt.cfg.SecretKey))

	res, e = lrt.Proxied.RoundTrip(req)

	if e != nil {
		fmt.Printf("Error: %v", e)
	} else {
		fmt.Printf("Received %v response\n", res.Status)
	}

	return
}
