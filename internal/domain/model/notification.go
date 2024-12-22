package model

type Notification struct {
	Type   string   `json:"type"`
	Event  string   `json:"event"`
	Object *Payment `json:"object"`
}

type TransactionType string

const (
	PaymentType TransactionType = "payment"
	RefundType  TransactionType = "refund"
	PayoutType  TransactionType = "payout"
	DealType    TransactionType = "deal"
)

type TransactionStatus string

const (
	Pending           TransactionStatus = "pending"
	WaitingForCapture TransactionStatus = "waiting_for_capture"
	Succeeded         TransactionStatus = "succeeded"
	Canceled          TransactionStatus = "canceled"
)
