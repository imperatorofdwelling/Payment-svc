package model

type (
	// PaymentReq define payment request.
	//
	// Description: The description of the transaction (no more than 128 characters), which you will see in the personal account of the user, and the user — when paying. For example: "Payment for order No. 72 for user@yoomoney.ru ".
	PaymentReq struct {
		Amount            `json:"amount" validate:"required"`
		Capture           bool `json:"capture"`
		Confirmation      `json:"confirmation"`
		PaymentMethodData `json:"payment_method_data,omitempty"`
		Description       string `json:"description,omitempty"`
	}

	// PaymentMethodData is data for payment by a specific method (payment_method). You don't have to pass this object in the request. In this case, the user will choose the payment method on the YUKASSA side.
	//
	// Type: payment method code
	//
	// Phone: the phone number from which the payment is made. It is specified in the ITU-T E.164 format, for example 79000000000. Leave empty if type is not mobile_balance
	PaymentMethodData struct {
		Type     PaymentMethodType `json:"type"`
		Phone    string            `json:"phone,omitempty"`
		BankCard `json:"card,omitempty"`
	}

	// BankCard Bank card data (required if you collect user card data on your side).
	//
	// Number: Bank card number
	BankCard struct {
		Number     string `json:"number"`
		ExpiryYear string `json:"expiry_year"`
	}

	// Amount specifies payment amount. Sometimes YouKassa partners charge an extra commission, which is not included in this amount.
	//
	// Value: The amount in the selected currency. Always a fractional value. The decimal separator is a dot, the thousand separator is missing. The number of digits after the dot depends on the selected currency. Example: 1000.00.
	//
	// Currency: Three-letter currency code in ISO-4217 format. Example: RUB. Must match the subaccount currency (recipient.gateway_id), if you share the payment flows, and the account currency (shopId in your personal account), if you do not share.
	Amount struct {
		Value    string `json:"value" validate:"required"`
		Currency `json:"currency"`
	}

	Confirmation struct {
		Type      string `json:"type"`
		ReturnURL string `json:"return_url"`
	}
)

type Currency string

var (
	RUB Currency = "RUB"
)

type PaymentMethodType string

var (
	SberLoan      PaymentMethodType = "sber_loan"
	MobileBalance PaymentMethodType = "mobile_balance"
)

type (
	PaymentStatus string

	// PaymentRes define payment result from YouKassa.
	//
	// ID specifies payment id.
	//
	// Status specifies payment status, possible values: pending, waiting_for_capture, succeeded, canceled.
	//
	// Paid specifies payment indication.
	//
	//
	PaymentRes struct {
		ID     string        `json:"id" validate:"required"`
		Status PaymentStatus `json:"status"`
		Paid   bool          `json:"paid"`
	}
)

var (
	Pending           PaymentStatus = "pending"
	WaitingForCapture PaymentStatus = "waiting_for_capture"
	Succeeded         PaymentStatus = "succeeded"
	Canceled          PaymentStatus = "canceled"
)
