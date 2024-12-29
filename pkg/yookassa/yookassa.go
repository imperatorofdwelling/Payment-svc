package yookassa

import (
	"bytes"
	"fmt"
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
	client          http.Client
	shopID          int
	secretKey       string
	payoutAgentID   int
	payoutSecretKey string
}

func NewYookassaClient(shopID int, secretKey string, payoutAgentID int, payoutSecretKey string) *Client {
	client := http.Client{}

	return &Client{
		client,
		shopID,
		secretKey,
		payoutAgentID,
		payoutSecretKey,
	}
}

type customRoundTripper struct {
	proxied        http.RoundTripper
	username       string
	password       string
	idempotencyKey string
}

func (lrt customRoundTripper) RoundTrip(req *http.Request) (res *http.Response, e error) {
	req.Header.Set("Content-Type", "application/json")
	if req.Method == http.MethodPost || req.Method == http.MethodDelete {
		req.Header.Set("Idempotence-Key", lrt.idempotencyKey)
	}

	req.SetBasicAuth(lrt.username, lrt.password)

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

	if param == "" {
		uri = fmt.Sprintf("%s/%s", yookassaApiAddr, endpoint)
	} else {
		uri = fmt.Sprintf("%s/%s/%s", yookassaApiAddr, endpoint, param)
	}

	req, err := http.NewRequest(method, uri, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	if query != nil {
		q := req.URL.Query()
		for paramName, paramVal := range query {
			q.Add(paramName, fmt.Sprintf("%v", paramVal))
		}
		req.URL.RawQuery = q.Encode()
	}

	customTripper := customRoundTripper{
		proxied:        http.DefaultTransport,
		idempotencyKey: idempotencyKey,
	}

	if endpoint == PaymentEndpoint {
		customTripper.username = strconv.Itoa(c.shopID)
		customTripper.password = c.secretKey
	} else if endpoint == PayoutEndpoint {
		customTripper.username = strconv.Itoa(c.payoutAgentID)
		customTripper.password = c.payoutSecretKey
	}

	c.client.Transport = customTripper

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
