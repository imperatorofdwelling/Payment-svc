package yookassa

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"github.com/imperatorofdwelling/payment-svc/internal/config"
	"net/http"
	"strconv"
)

const (
	PaymentEndpoint = "payments"
	CaptureEndpoint = "capture"
	CancelEndpoint  = "cancel"
)

type Client struct {
	client      http.Client
	cfgYookassa config.PayApi
}

func NewYookassaClient(cfgPayApi config.PayApi) *Client {
	client := http.Client{
		Transport: loggingRoundTripper{
			proxied:   http.DefaultTransport,
			shopID:    strconv.Itoa(cfgPayApi.ShopID),
			secretKey: cfgPayApi.SecretKey,
		},
	}

	return &Client{
		client,
		cfgPayApi,
	}
}

type loggingRoundTripper struct {
	proxied   http.RoundTripper
	shopID    string
	secretKey string
}

func (lrt loggingRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(lrt.shopID, lrt.secretKey)

	res, e = lrt.proxied.RoundTrip(req)

	if e != nil {
		fmt.Printf("Error: %v", e)
	} else {
		fmt.Printf("Received %v response\n", res.Status)
	}

	return
}

func (c *Client) makeRequest(
	method string,
	endpoint string,
	body []byte,
	params map[string]interface{},
	idempotencyKey string,
) (*http.Response, error) {
	uri := fmt.Sprintf("%s/%s", c.cfgYookassa.Addr, endpoint)

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if idempotencyKey == "" {
		idempotencyKey = uuid.NewString()
	}

	if method == http.MethodPost || method == http.MethodDelete {
		req.Header.Set("Idempotence-Key", idempotencyKey)
	}

	if params != nil {
		q := req.URL.Query()
		for paramName, paramVal := range params {
			q.Add(paramName, fmt.Sprintf("%v", paramVal))
		}
		req.URL.RawQuery = q.Encode()
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
