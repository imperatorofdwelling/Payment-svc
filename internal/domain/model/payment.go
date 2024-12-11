package model

import "time"

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
		Type           PaymentMethodType `json:"type" validate:"required"`
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
		Type      ConfirmationType `json:"type" validate:"required"`
		ReturnURL string           `json:"return_url,omitempty" validate:"required_if=Type redirect,required_if=Type mobile_application,omitempty,url"`
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
		Amount `json:"amount" validate:"omitempty,required_if=Type calculated,required_if=Type mixed"`
		Rate   string `json:"rate" validate:"omitempty,numeric,required_if=Type calculated"`
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

	AlfabankType     PaymentMethodType = "alfabank"
	InstallmentsType PaymentMethodType = "installments"
	ApplePayType     PaymentMethodType = "apple_pay"
	GooglePayType    PaymentMethodType = "google_pay"
	QiwiType         PaymentMethodType = "qiwi"
	WechatType       PaymentMethodType = "wechat"
	WebmoneyType     PaymentMethodType = "webmoney"
)

type (

	// PaymentRes define payment result from YouKassa.
	//
	// ID specifies payment id.
	//
	// Status specifies payment status, possible values: pending, waiting_for_capture, succeeded, canceled.
	//
	// Paid specifies payment indication.
	//
	// Amount: The amount of the payment. Sometimes YUKASSA partners charge the user an additional commission, which is not included in this amount.
	//
	// AuthorizationDetails: Information about the authorization of the payment when paying with a bank card. Only available for these payment methods: bank card, MirPay, SberPay, T-Pay.
	//
	// CreatedAt: The time of the order creation. It is specified in UTC and transmitted in ISO 8601 format. Example: 2017-11-03T11:52:31.827Z
	//
	// Description: The description of the transaction (no more than 128 characters), which you will see in the personal account of the user, and the user — when paying.
	//
	// ExpiresAt: The time before which you can cancel or confirm the payment for free. At the specified time, the payment with the waiting_for_capture status will be automatically canceled. It is specified in UTC and transmitted in ISO 8601 format. Example: 2017-11-03T11:52:31.827Z
	//
	// Metadata: Any additional data that you need to work with (for example, your internal order ID). They are transmitted as a set of key-value pairs and are returned in a response from UCassa. Restrictions: maximum of 16 keys, the key name is no more than 32 characters, the key value is no more than 512 characters, the data type is a string in UTF-8 format.
	//
	// PaymentMethod: The payment method that was used for this payment.
	//
	// Recipient: The recipient of the payment.
	//
	// Refundable: The ability to make a refund via the API.
	//
	// Test: A sign of a test operation.
	//
	// IncomeAmount: The payment amount that the store will receive is the amount value minus the Yandex commission. If you are a partner and use an OAuth token to authenticate requests, ask the store for the right to receive information about commissions for payments.
	PaymentRes struct {
		ID                   string        `json:"id" validate:"required,len=36,uuid"`
		Status               PaymentStatus `json:"status" validate:"required, oneof=pending waiting_for_capture succeeded canceled"`
		Paid                 bool          `json:"paid" validate:"required"`
		Amount               `json:"amount" validate:"required"`
		AuthorizationDetails `json:"authorization_details" validate:"omitempty"`
		CreatedAt            time.Time `json:"created_at" validate:"required,datetime"`
		Description          string    `json:"description" validate:"omitempty,lte=128"`
		ExpiresAt            time.Time `json:"expires_at" validate:"omitempty,datetime"`
		Metadata             any       `json:"metadata" validate:"omitempty"`
		PaymentMethod        `json:"payment_method" validate:"omitempty"`
		Recipient            `json:"recipient" validate:"required"`
		Refundable           bool   `json:"refundable" validate:"required"`
		Test                 bool   `json:"test" validate:"required"`
		IncomeAmount         Amount `json:"income_amount" validate:"omitempty"`
	}

	// Recipient specifies recipient of the payment.
	//
	// AccountID: Store ID in Kassa.
	//
	// GatewayID: The account ID. It is used to separate payment flows within the same account.
	Recipient struct {
		AccountID string `json:"account_id" validate:"required"`
		GatewayID string `json:"gateway_id" validate:"required"`
	}

	// AuthorizationDetails specifies information about the authorization of the payment when paying with a bank card. Only available for these payment methods: bank card, MirPay, SberPay, T-Pay.
	//
	// RRN: Retrieval Reference Number — уникальный идентификатор транзакции в системе эмитента.
	//
	// AuthCode: The authorization code. It is issued by the issuer and confirms the authorization.
	//
	// ThreeDSecure: Data about the user's authentication using 3‑D Secure to confirm the payment.
	AuthorizationDetails struct {
		RRN          string `json:"rrn,omitempty" validate:"omitempty,numeric"`
		AuthCode     string `json:"auth_code,omitempty" validate:"omitempty,numeric"`
		ThreeDSecure `json:"three_d_secure" validate:"required"`
	}

	// ThreeDSecure specifies data about the user's authentication using 3‑D Secure to confirm the payment.
	//
	// Applied: Displaying a form to the user for authentication using 3‑D Secure. Possible values: true — Yandex displayed a form to the user so that he could authenticate using 3‑D Secure; false — the payment was made without authentication using 3‑D Secure.
	ThreeDSecure struct {
		Applied bool `json:"applied" validate:"required"`
	}

	// PaymentMethod specifies method that was used for this payment.
	//
	// ID: The ID of the payment method.
	//
	// Saved: You can use the saved payment method to make non-acceptance debits.
	//
	// Title: The name of the payment method.
	//
	// DiscountAmount: The amount of the discount for installments. Present for payments with the waiting_for_capture and succeeded status if the user has selected an installment plan.
	//
	// LoanOption: The credit rate that the user selected when paying. Possible values: loan — loan; installments_XX — installment plan, where XX is the number of months to pay the installment plan. For example, installments_3 is a 3—month installment plan. Present for payments with the waiting_for_capture and succeeded status.
	//
	// Login: The user's login in Alpha Click (linked phone or additional login).
	//
	// BankCard: Bank card details.
	//
	// PayerBankDetails: Details of the account that was used for payment. Required parameter for payments with the succeeded status. In other cases, it may be missing.
	PaymentMethod struct {
		Type             PaymentMethodType `json:"type" validate:"required"`
		ID               string            `json:"id" validate:"required,uuid"`
		Saved            bool              `json:"saved" validate:"required"`
		Title            string            `json:"title,omitempty" validate:"omitempty"`
		DiscountAmount   Amount            `json:"discount_amount,omitempty" validate:"omit_with=Type sber_loan"`
		LoanOption       string            `json:"loan_option,omitempty" validate:"omit_with=Type sber_loan"`
		Login            string            `json:"login,omitempty" validate:"omit_with=Type alfabank"`
		BankCard         BankCardRes       `json:"card,omitempty" validate:"omit_with=Type bank_card,omitempty"`
		PayerBankDetails `json:"payer_bank_details,omitempty" validate:"omit_with=Type sbp|omit_with=Type b2b_sberbank,omitempty"`
		SBPOperationID   string `json:"sbp_operation_id,omitempty" validate:"omit_with=Type sbp"`
	}

	// BankCardRes specifies bank card details.
	//
	// First6: The first 6 numbers of the card (trash can). Please note that the card contained in cash registers and other vaults may not be responsible for the last 4 years, the expired year, the expired month.
	//
	// Last4: The last 4 digits of the card number.
	//
	// ExpiryYear: Validity period, year, YYYY.
	//
	// ExpiryMonth: Validity period, month, MM.
	//
	// CardType: The type of bank card. Possible values: MasterCard (for Mastercard and Maestro cards), Visa (for Visa and Visa Electron cards), Mir, UnionPay, JCB, AmericanExpress, DinersClub, DiscoverCard, InstaPayment, InstaPaymentTM, Laser, Dankort, Solo, Switch and Unknown.
	//
	// CardProduct: The card product of the payment system with which the bank card is associated. For example, the card products of the Mir payment system: Mir Classic, Mir Classic Credit, MIR Privilege Plus and others.
	//
	// IssuerCountry: The name of the bank that issued the card.
	//
	// Source: The source of the bank card data. Possible values: mir_pay, apple_pay, google_pay. It is assigned if the user selected a card saved in Mir Pay, Apple Pay or Google Pay when paying.
	BankCardRes struct {
		First6        string `json:"first6,omitempty" validate:"omitempty,numeric,len=6"`
		Last4         string `json:"last4" validate:"required,numeric,len=4"`
		ExpiryYear    string `json:"expiry_year" validate:"required,numeric,len=4"`
		ExpiryMonth   string `json:"expiry_month" validate:"required,numeric,max=2"`
		CardType      string `json:"card_type" validate:"required,oneof=MasterCard Visa Mir UnionPay JCB AmericanExpress DinersClub DiscoverCard InstaPayment InstaPaymentTM Laser Dankort Solo Switch Unknown"`
		CardProduct   `json:"card_product,omitempty" validate:"omitempty"`
		IssuerCountry string `json:"issuer_country,omitempty" validate:"omitempty,iso3166_1_alpha2"`
		Source        string `json:"source,omitempty" validate:"omitempty,oneof=mir_pay apple_pay google_pay"`
	}

	// CardProduct specifies card product of the payment system with which the bank card is associated. For example, the card products of the Mir payment system: Mir Classic, Mir Classic Credit, MIR Privilege Plus and others.
	//
	// Code: The code of the card product. Example: MCP
	//
	// Name: The name of the card product. Example: MIR Privilege
	CardProduct struct {
		Code string `json:"code" validate:"required"`
		Name string `json:"name,omitempty" validate:"omitempty"`
	}

	// PayerBankDetails specifies details of the account that was used for payment. Required parameter for payments with the succeeded status. In other cases, it may be missing.
	//
	// PayerBankDetails: The ID of the bank or payment service in the SBP (NSPK).
	//
	// BIC: The bank identification code (BIC) of the bank or payment service.
	//
	// FullName: The full name of the organization.
	//
	// ShortName: The abbreviated name of the organization.
	//
	// Address: The address of the organization.
	//
	// INN: The individual tax number (INN) of the organization.
	//
	// BankName: The name of the organization's bank.
	//
	// BankBranch: A branch of the organization's bank.
	//
	// BankBIC: The bank identification code (BIC) of the organization's bank.
	//
	// Account: The account number of the organization.
	//
	// KPP: The code of the reason for registration (KPP) of the organization.
	PayerBankDetails struct {
		BankID     string `json:"bank_id,omitempty" validate:"omitempty,numeric,max=12"`
		BIC        string `json:"bic" validate:"required,bic"`
		FullName   string `json:"full_name" validate:"omitempty,max=800"`
		ShortName  string `json:"short_name,omitempty" validate:"omitempty,max=160"`
		Address    string `json:"address,omitempty" validate:"omitempty,max=500"`
		INN        string `json:"inn,omitempty" validate:"omitempty,len=10|len=12"`
		BankName   string `json:"bank_name,omitempty" validate:"omitempty,max=350,min=1"`
		BankBranch string `json:"bank_branch,omitempty" validate:"omitempty,max=140,min=1"`
		BankBIC    string `json:"bank_bik,omitempty" validate:"omitempty,bic"`
		Account    string `json:"account,omitempty" validate:"omitempty,len=20"`
		KPP        string `json:"kpp,omitempty" validate:"omitempty,len=9"`
	}
)

type PaymentStatus string

var (
	Pending           PaymentStatus = "pending"
	WaitingForCapture PaymentStatus = "waiting_for_capture"
	Succeeded         PaymentStatus = "succeeded"
	Canceled          PaymentStatus = "canceled"
)
