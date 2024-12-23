package yookassa

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

const yookassaApiAddr = "https://api.yookassa.ru/v3"

type YookassaEndpoint string

const (
	PaymentEndpoint YookassaEndpoint = "payments"
	PayoutEndpoint  YookassaEndpoint = "payouts"
	CaptureEndpoint YookassaEndpoint = "capture"
	CancelEndpoint  YookassaEndpoint = "cancel"
)

type Client struct {
	client    http.Client
	shopID    int
	secretKey string
}

func NewYookassaClient(shopID int, secretKey string) *Client {
	client := http.Client{
		Transport: loggingRoundTripper{
			proxied:   http.DefaultTransport,
			shopID:    strconv.Itoa(shopID),
			secretKey: secretKey,
		},
	}

	return &Client{
		client,
		shopID,
		secretKey,
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
	endpoint YookassaEndpoint,
	param string,
	body []byte,
	query map[string]interface{},
	idempotencyKey string,
) (*http.Response, error) {
	var uri string

	if param != "" {
		uri = fmt.Sprintf("%s/%s", yookassaApiAddr, endpoint)
	} else {
		uri = fmt.Sprintf("%s/%s/%s", yookassaApiAddr, endpoint, param)
	}

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

	if query != nil {
		q := req.URL.Query()
		for paramName, paramVal := range query {
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
