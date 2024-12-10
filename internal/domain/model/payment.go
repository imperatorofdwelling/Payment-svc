package model

type (
	// PaymentReq define payment request.
	//
	// Amount: payment amount. Sometimes YouKassa partners charge an extra commission, which is not included in this amount.
	//
	// Capture: Automatic acceptance of the received payment. Possible values: true — the payment is debited immediately (payment in one stage); false — the payment is held and debited upon your request (payment in two stages ).
	//
	// PaymentMethodData: data for payment by a specific method (payment_method). You don't have to pass this object in the request. In this case, the user will choose the payment method on the YUKASSA side.
	//
	// Confirmation: The data required to initiate the selected payment confirmation scenario by the user.
	//
	// Description: The description of the transaction (no more than 128 characters), which you will see in the personal account of the user, and the user — when paying. For example: "Payment for order No. 72 for user@yoomoney.ru ".
	PaymentReq struct {
		Amount            `json:"amount" validate:"required"`
		Capture           bool `json:"capture,omitempty"`
		PaymentMethodData `json:"payment_method_data,omitempty" validate:"omitempty"`
		Confirmation      `json:"confirmation,omitempty" validate:"omitempty"`
		Description       string `json:"description,omitempty" validate:"omitempty,lte=128"`
	}

	// PaymentMethodData is data for payment by a specific method (payment_method). You don't have to pass this object in the request. In this case, the user will choose the payment method on the YUKASSA side.
	//
	// Type: payment method code.
	//
	// Phone: the phone number from which the payment is made. It is specified in the ITU-T E.164 format, for example 79000000000. Leave empty if type is not mobile_balance.
	//
	// BankCard: bank card data (required if you collect user card data on your side).
	//
	// PaymentPurpose: The purpose of the payment (no more than 210 characters).
	//
	// VatData: data on value added tax (VAT). The payment may or may not be subject to VAT. Goods can be taxed at the same VAT rate or at different rates.
	//
	// Articles: Shopping cart (in NSPK terms) is a list of goods that can be paid for using a certificate. It is necessary to transfer it only when paying on the ready-made Yandex.Checkout page.
	//
	// ElectronicCertificate: data from the FES NSPK for payment by electronic certificate. It is necessary to transfer only when paying with data collection on your side.
	PaymentMethodData struct {
		Type           PaymentMethodType `json:"type" validate:"required,should_exist_field=mobile_balance Phone b2b_sberbank PaymentPurpose b2b_sberbank VatData"`
		Phone          string            `json:"phone,omitempty" validate:"omitempty,e164,required_if=Type mobile_balance"`
		BankCard       `json:"card,omitempty" validate:"omitempty"`
		PaymentPurpose string `json:"payment_purpose,omitempty" validate:"omitempty,max=210,required_if=Type b2b_sberbank"`
		VatData        `json:"vat_data" validate:"omitempty,required_if=Type b2b_sberbank"`
		// TODO Create Articles struct
		Articles              []any `json:"articles" validate:"omitempty"`
		ElectronicCertificate `json:"electronic_certificate,omitempty" validate:"omitempty"`
	}

	// ElectronicCertificate specifies data from the FES NSPK for payment by electronic certificate. It is necessary to transfer only when paying with data collection on your side.
	//
	// Amount: The amount to be used for an electronic certificate is the totalCertAmount value that you received from the FES NSPK in the request for preliminary approval of the use of the certificate (Pre-Auth). The amount must not exceed the total amount of the payment (amount).
	//
	// BasketID: The identifier of the shopping cart generated in the NSPK is the purchaseBasketId value that you received from the NSPK FES in the request for preliminary approval of the use of the certificate (Pre-Auth).
	ElectronicCertificate struct {
		Amount   `json:"amount" validate:"required"`
		BasketID string `json:"basket_id" validate:"required"`
	}

	// BankCard Bank card data (required if you collect user card data on your side).
	//
	// Number: Bank card number.
	//
	// ExpiryYear: Validity period, year, YYYY.
	//
	// ExpiryMonth: Validity period, month, MM.
	//
	// CardHolder: The name of the cardholder.
	//
	// Csc: The CVC2 or CVV2 code, 3 or 4 characters, is printed on the back of the card.
	BankCard struct {
		Number      string `json:"number" validate:"required,credit_card"`
		ExpiryYear  string `json:"expiry_year" validate:"required,numeric,max=4"`
		ExpiryMonth string `json:"expiry_month" validate:"required,numeric,max=2"`
		CardHolder  string `json:"cardholder,omitempty" validate:"omitempty,max=26"`
		Csc         string `json:"csc,omitempty" validate:"omitempty,min=3,max=4"`
	}

	// Amount specifies payment amount. Sometimes YouKassa partners charge an extra commission, which is not included in this amount.
	//
	// Value: The amount in the selected currency. Always a fractional value. The decimal separator is a dot, the thousand separator is missing. The number of digits after the dot depends on the selected currency. Example: 1000.00.
	//
	// Currency: Three-letter currency code in ISO-4217 format. Example: RUB. Must match the subaccount currency (recipient.gateway_id), if you share the payment flows, and the account currency (shopId in your personal account), if you do not share.
	Amount struct {
		Value    string `json:"value" validate:"required,money"`
		Currency `json:"currency" validate:"required,currency"`
	}

	// Confirmation specifies the data required to initiate the selected payment confirmation scenario by the user.
	//
	// Type: the code of the confirmation script. Could be on of those: redirect, embedded, external, mobile _application, qr.
	//
	// ReturnURL: The URL or deep link to which the user will return after confirming or canceling the payment in the application. If the payment was made from the mobile version of the site, pass the URL, if from the mobile application — a deep link. Maximum of 2048 characters.
	//
	// Locale: The language of the interface, emails and SMS that the user will see or receive. The format complies with ISO/IEC 15897. Possible values: ru_RU, en_US. The case is important.
	//
	// Enforce: Request to make a payment with 3-D Secure authentication. It will work if you accept payment by bank card by default without confirmation of payment by the user. In all other cases, the 3-D Secure authentication will be managed by UCassa. If you want to accept payments without additional confirmation by the user, write to your UCassa manager.
	Confirmation struct {
		Type      ConfirmationType `json:"type" validate:"required,should_exist_field=redirect ReturnURL mobile_application ReturnURL"`
		ReturnURL string           `json:"return_url,omitempty" validate:"omitempty,url,required_if=Type redirect Type mobile_application"`
		Locale    string           `json:"locale,omitempty" validate:"omitempty,oneof=ru_RU en_US"`
		Enforce   bool             `json:"enforce,omitempty" validate:"omitempty"`
	}

	// VatData specifies data on value added tax (VAT). The payment may or may not be subject to VAT. Goods can be taxed at the same VAT rate or at different rates.
	//
	// Amount: The amount of VAT.
	//
	// Rate: The tax rate (as a percentage). Possible values are 7, 10, 18 and 20.
	VatData struct {
		Type   string `json:"type" validate:"required,vat"`
		Amount `json:"amount" validate:"omitempty,required_if=Type calculated Type mixed"`
		Rate   string `json:"rate" validate:"omitempty,numeric,required_if=Type:calculated"`
	}
)

type VatDataType string

var (
	Untaxed    VatDataType = "untaxed"
	Calculated VatDataType = "calculated"
	Mixed      VatDataType = "mixed"
)

type ConfirmationType string

var (
	Embedded          ConfirmationType = "embedded"
	External          ConfirmationType = "external"
	MobileApplication ConfirmationType = "mobile_application"
	QR                ConfirmationType = "qr"
	Redirect          ConfirmationType = "redirect"
)

type Currency string

var (
	RUB Currency = "RUB"

	ValidCurrencies = map[string]struct{}{
		string(RUB): {},
	}
)

type PaymentMethodType string

var (
	SberLoanType       PaymentMethodType = "sber_loan"
	MobileBalanceType  PaymentMethodType = "mobile_balance"
	BankCardType       PaymentMethodType = "bank_card"
	CashType           PaymentMethodType = "cash"
	SBPType            PaymentMethodType = "sbp"
	B2BSberType        PaymentMethodType = "b2b_sberbank"
	ElectronicCertType PaymentMethodType = "electronic_certificate"
	YooMoneyType       PaymentMethodType = "yoo_money"
	SberPayType        PaymentMethodType = "sberbank"
	TinkoffBankType    PaymentMethodType = "tinkoff_bank"
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
