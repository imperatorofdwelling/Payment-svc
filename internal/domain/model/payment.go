package model

import (
	"time"
)

type (
	// Payment object contains all currently relevant information about the payment.
	//
	// ID: payment id.
	//
	// Status specifies payment status, possible values: pending, waiting_for_capture, succeeded, canceled.
	//
	// Amount: payment amount. Sometimes YouKassa partners charge an extra commission, which is not included in this amount.
	//
	// IncomeAmount: amount of payment to be received by the store: the amount value minus the YooMoney commission.
	//
	// Capture: Automatic acceptance of the received payment. Possible values: true — the payment is debited immediately (payment in one stage); false — the payment is held and debited upon your request (payment in two stages ).
	//
	// Description: The description of the transaction (no more than 128 characters), which you will see in the personal account of the user, and the user — when paying. For example: "Payment for order No. 72 for user@yoomoney.ru ".
	//
	// Receipt: Data for the formation of the receipt. It is necessary to transmit in these cases: you are a company or sole proprietor and you use Checks from YUKASSA to pay in compliance with the requirements of Federal Law No. 54; you are a company or sole proprietor, for payment in compliance with the requirements of Federal Law No. 54, you use a third-party online sales register and send data for receipts according to one of the scenarios: Payment and receipt at the same time, or First the receipt, then the payment ; you are self-employed and use the UCassa solution for auto-sending checks.
	//
	// Recipient: payment recipient. Required for separating payment flows within one account or making payments to other accounts.
	//
	// PaymentMethodData: data for payment by a specific method (payment_method). You don't have to pass this object in the request. In this case, the user will choose the payment method on the YUKASSA side.
	//
	// CapturedAt: Time of order creation, based on UTC and specified in the ISO 8601 format.
	//
	// CreatedAt: Time of order creation, based on UTC and specified in the ISO 8601 format.
	//
	// ExpiresAt: The period during which you can cancel or capture a payment for free.
	//
	// Confirmation: The data required to initiate the selected payment confirmation scenario by the user.
	//
	// Test: A sign of a test operation.
	//
	// RefundedAmount: The amount refunded to the user. Specified if the payment has successful refunds.
	//
	// Paid specifies payment indication.
	//
	// Refundable: The ability to make a refund via the API.
	//
	// ReceiptRegistration: status of receipt delivery.
	//
	// Metadata: Any additional data that you need to work with (for example, your internal order ID). They are transmitted as a set of key-value pairs and are returned in a response from UCassa. Restrictions: maximum of 16 keys, the key name is no more than 32 characters, the key value is no more than 512 characters, the data type is a string in UTF-8 format.
	//
	// PaymentToken: One-time payment token generated with Checkout.js or mobile SDK .
	//
	// PaymentMethodID: Saved payment method 's ID.
	//
	// SavePaymentMethod: Saving payment data for making autopayments. Possible values: true — save the payment method (save the payment data); false — make a payment without saving the payment method. Available only after consultation with the Kassa manager.
	//
	// ClientIP: User’s IPv4 or IPv6 address. If not specified, the TCP connection’s IP address is used.
	//
	// Metadata: Any additional data you might require for processing payments (for example, your internal order ID), specified as a “key-value” pair and returned in response from YooMoney. Limitations: no more than 16 keys, no more than 32 characters in the key’s title, no more than 512 characters in the key’s value, data type is a string in the UTF-8 format.
	//
	// CancellationDetails: Commentary to the canceled status: who and why canceled the payment.
	//
	// AuthorizationDetails: Information about the authorization of the payment when paying with a bank card. Only available for these payment methods: bank card, MirPay, SberPay, T-Pay.
	//
	// Transfers: Information about money distribution: the amounts of transfers and the stores to be transferred to. Specified if you use Split payments .
	//
	// Deal: The deal within which the payment is being carried out. Specified if you use Safe deal .
	//
	// MerchantCustomerID: The identifier of the customer in your system, such as email address or phone number. No more than 200 characters. Specified if you want to save a bank card and offer it for a recurring payment in the YooMoney payment widget .
	Payment struct {
		ID                   string            `json:"id,omitempty"`
		Status               TransactionStatus `json:"status,omitempty" validate:"omitempty,oneof=pending waiting_for_capture succeeded canceled"`
		Amount               *Amount           `json:"amount,omitempty"`
		IncomeAmount         *Amount           `json:"income_amount,omitempty"`
		Capture              bool              `json:"capture,omitempty"`
		Description          string            `json:"description,omitempty" validate:"omitempty,max=128"`
		Receipt              *Receipt          `json:"receipt,omitempty" validate:"omitempty"`
		Recipient            *Recipient        `json:"recipient,omitempty" validate:"omitempty"`
		PaymentMethodData    `json:"payment_method,omitempty" validate:"omitempty"`
		CapturedAt           *time.Time `json:"captured_at,omitempty" validate:"omitempty,datetime"`
		CreatedAt            *time.Time `json:"created_at,omitempty" validate:"omitempty,datetime"`
		ExpiresAt            *time.Time `json:"expires_at,omitempty" validate:"omitempty,datetime"`
		Confirmation         `json:"confirmation,omitempty" validate:"omitempty"`
		Test                 bool                  `json:"test,omitempty"`
		RefundedAmount       *Amount               `json:"refunded_amount,omitempty"`
		Paid                 bool                  `json:"paid,omitempty"`
		Refundable           bool                  `json:"refundable,omitempty"`
		ReceiptRegistration  TransactionStatus     `json:"receipt_registration,omitempty"`
		Metadata             interface{}           `json:"metadata" validate:"omitempty"`
		CancellationDetails  *CancellationDetails  `json:"cancellation_details,omitempty"`
		AuthorizationDetails *AuthorizationDetails `json:"authorization_details" validate:"omitempty"`
		Transfers            *Transfers            `json:"transfers,omitempty" validate:"omitempty"`
		Deal                 *Deal                 `json:"deal,omitempty" validate:"omitempty"`
		MerchantCustomerID   string                `json:"merchant_customer_id,omitempty" validate:"omitempty,max=200"`
	}

	// CancellationDetails define commentary to the canceled status: who and why canceled the payment..
	//
	// Party: The participant of the payment process that made the decision to cancel the payment. Possible values are yoo_money, payment_network, and merchant.
	//
	// Reason: reason behind the cancelation.
	CancellationDetails struct {
		Party  string `json:"party,omitempty"`
		Reason string `json:"reason,omitempty"`
	}

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
	//
	// Receipt: Data for the formation of the receipt. It is necessary to transmit in these cases: you are a company or sole proprietor and you use Checks from YUKASSA to pay in compliance with the requirements of Federal Law No. 54; you are a company or sole proprietor, for payment in compliance with the requirements of Federal Law No. 54, you use a third-party online sales register and send data for receipts according to one of the scenarios: Payment and receipt at the same time, or First the receipt, then the payment ; you are self-employed and use the UCassa solution for auto-sending checks.
	//
	// Recipient: Payment recipient. Required for separating payment flows within one account or making payments to other accounts.
	//
	// PaymentToken: One-time payment token generated with Checkout.js or mobile SDK .
	//
	// PaymentMethodID: Saved payment method 's ID.
	//
	// SavePaymentMethod: Saving payment data for making autopayments. Possible values: true — save the payment method (save the payment data); false — make a payment without saving the payment method. Available only after consultation with the Kassa manager.
	//
	// ClientIP: User’s IPv4 or IPv6 address. If not specified, the TCP connection’s IP address is used.
	//
	// Metadata: Any additional data you might require for processing payments (for example, your internal order ID), specified as a “key-value” pair and returned in response from YooMoney. Limitations: no more than 16 keys, no more than 32 characters in the key’s title, no more than 512 characters in the key’s value, data type is a string in the UTF-8 format.
	//
	// Airline: Object containing the data for selling airline tickets. Used only for bank card payments.
	//
	// Transfers: Information about money distribution: the amounts of transfers and the stores to be transferred to. Specified if you use Split payments .
	//
	// Deal: The deal within which the payment is being carried out. Specified if you use Safe deal .
	//
	// MerchantCustomerID: The identifier of the customer in your system, such as email address or phone number. No more than 200 characters. Specified if you want to save a bank card and offer it for a recurring payment in the YooMoney payment widget .
	//
	// Receiver: Payment receiver's details specified when you want to add money to e-wallet, bank account, or phone balance .
	PaymentReq struct {
		Amount  `json:"amount" validate:"required"`
		Capture bool `json:"capture,omitempty"`
		//PaymentMethodData `json:"payment_method_data,omitempty" validate:"omitempty"`
		Confirmation `json:"confirmation,omitempty" validate:"omitempty"`
		Description  string `json:"description,omitempty" validate:"omitempty,max=128"`
		//Receipt            `json:"receipt,omitempty" validate:"omitempty"`
		//Recipient          `json:"recipient,omitempty" validate:"omitempty"`
		//PaymentToken       string `json:"payment_token,omitempty"`
		//PaymentMethodID    string `json:"payment_method_id,omitempty" validate:"omitempty,uuid"`
		//SavePaymentMethod  bool   `json:"save_payment_method,omitempty"`
		//ClientIP           string `json:"client_ip,omitempty" validate:"omitempty,ip4_addr|ip6_addr"`
		//Metadata           any    `json:"metadata,omitempty" validate:"omitempty"`
		//Airline            `json:"airline,omitempty" validate:"omitempty"`
		//Transfers          `json:"transfers,omitempty" validate:"omitempty"`
		//Deal               `json:"deal,omitempty" validate:"omitempty"`
		//MerchantCustomerID string `json:"merchant_customer_id,omitempty" validate:"omitempty,max=200"`
		//Receiver           `json:"receiver,omitempty" validate:"omitempty"`
	}

	// Receipt specifies data for the formation of the receipt. It is necessary to transmit in these cases: you are a company or sole proprietor and you use Checks from YUKASSA to pay in compliance with the requirements of Federal Law No. 54; you are a company or sole proprietor, for payment in compliance with the requirements of Federal Law No. 54, you use a third-party online sales register and send data for receipts according to one of the scenarios: Payment and receipt at the same time, or First the receipt, then the payment ; you are self-employed and use the UCassa solution for auto-sending checks.
	//
	// Customer: User details. You should specify at least the basic contact information: email address (customer.email) or phone number (customer.phone).
	//
	// Items: List of products in an order. Receipts sent in accordance with 54-FZ can contain up to 100 items. Receipts for the self-employed can contain up to six items.
	//
	// TaxSystemCode: Store's tax system (tag 1055 in 54-FZ). The parameter is required if you use the ATOL online sales register updated to FFD 1.2, or if you use several tax systems. Otherwise, the parameter is not specified.
	//
	// ReceiptIndustryDetails: Industry attribute of the receipt (tag 1261 in 54-FZ). Must be specified if FFD 1.2 is used.
	//
	// ReceiptOperationalDetails: Transaction attribute of the receipt (tag 1270 in 54-FZ). Must be specified if FFD 1.2 is used.
	Receipt struct {
		Customer                  `json:"customer,omitempty" validate:"omitempty"`
		Items                     []ReceiptItem           `json:"items" validate:"required"`
		TaxSystemCode             int                     `json:"tax_system_code,omitempty" validate:"omitempty,gte=1,lte=6"`
		ReceiptIndustryDetails    []ReceiptIndustryDetail `json:"receipt_industry_details,omitempty" validate:"omitempty"`
		ReceiptOperationalDetails `json:"receipt_operational_details,omitempty" validate:"omitempty"`
	}

	// ReceiptItem specifies list of products in an order. Receipts sent in accordance with 54-FZ can contain up to 100 items. Receipts for the self-employed can contain up to six items.
	//
	// Description: Product name (maximum 128 characters). Tag 1030 in 54-FZ.
	//
	// VatCode: VAT rate (tag 1199 in 54-FZ). In receipts sent in accordance with 54-FZ, possible value is a number from 1 to 6.
	//
	// Quantity: Product quantity (tag 1023 in 54-FZ). In receipts sent in accordance with 54-FZ, the maximum possible value depends on the model of your online sales register. Receipts for the self-employed can only specify positive integers (without separator and fractional part). Example: 1.
	//
	// Measure: Unit of measurement of product quantity: for example, items or grams. Tag 2108 in 54-FZ. This parameter must be specified starting from FFD 1.2. List of possible measures: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/other-services/parameters-values#measure
	//
	// MarkQuantity: Fraction of a marked product (tag 1291 in 54-FZ). Must be specified if all of the following applies: FFD version 1.2 is used; payment is made for a marked product; the measure field has the piece value. Example: you're selling pencils by the piece. They're supplied in packages 100 pencils each with one marking code. To sell one pencil, enter 1 in numerator and 100 in denominator.
	//
	// PaymentSubject: Payment subject attribute (tag 1212 in 54-FZ): what the payment is made for, for example, a product or service. List of possible subjects: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/other-services/parameters-values#payment-subject
	//
	// PaymentMode: Payment method attribute (tag 1214 in 54-FZ): contains information about the payment method and shows whether the product has been handed over to the customer. Example: a customer makes a full payment for a product and immediately receives it. In this case, the full_payment (full payment) value must be specified. List of possible values: https://yookassa.ru/developers/payment-acceptance/receipts/54fz/other-services/parameters-values#payment-mode.
	//
	// CountryOfOriginCode: Country of origin code according to the Russian classifier of world countries (OK (MK (ISO 3166) 004-97) 025-2001). Tag 1230 in 54-FZ. Example: RU. Online sales register that support this parameter: Orange Data, Kit Invest.
	//
	// CustomsDeclarationNumber: Customs declaration number (1 to 32 characters). Tag 1231 in 54-FZ. Online sales register that support this parameter: Orange Data, Kit Invest.
	//
	// Excise: Amount of excise tax on products including kopeks. Tag 1229 in 54-FZ. Decimal number with 2 digits after the period. Online sales register that support this parameter: Orange Data, Kit Invest. Example:20.00.
	//
	// ProductCode: Product code is a unique number assigned to a unit of product during marking process. Tag 1162 in 54-FZ. Format: hexadecimal number with spaces. Maximum length is 32 bytes. Example: 00 00 00 01 00 21 FA 41 00 23 05 41 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 12 00 AB 00. Required parameter for marked products.
	//
	// MarkCodeInfo: Product code (tag 1163 in 54-FZ). Must be specified if the FFD 1.2 protocol is used and if the product must be marked. At least one of the fields must be filled in.
	//
	// MarkMode: Method of processing the marking code (tag 2102 in 54-FZ). Must be specified if all of the following applies: FFD version 1.2 is used; payment is made for a marked product; an online sales register from ATOL Online or BusinessRu is used. Must get a value equal to "0". Pattern:^[0]{1}$
	//
	// PaymentSubjectIndustryDetails: Industry attribute of the payment subject (tag 1260 in 54-FZ). Must be specified if FFD 1.2 is used.
	ReceiptItem struct {
		Description              string `json:"description" validate:"required,max=128"`
		Amount                   `json:"amount" validate:"required"`
		VatCode                  string `json:"vat_code" validate:"required,gte=1,lte=6"`
		Quantity                 int    `json:"quantity" validate:"required,number"`
		Measure                  string `json:"measure,omitempty" validate:"omitempty"`
		MarkQuantity             `json:"mark_quantity,omitempty" validate:"omitempty"`
		PaymentSubject           string `json:"payment_subject,omitempty" validate:"omitempty"`
		PaymentMode              string `json:"payment_mode,omitempty" validate:"omitempty"`
		CountryOfOriginCode      string `json:"country_of_origin_code,omitempty" validate:"omitempty"`
		CustomsDeclarationNumber string `json:"customs_declaration_number,omitempty" validate:"omitempty,min=1,max=32"`
		Excise                   string `json:"excise,omitempty" validate:"omitempty"`
		ProductCode              string `json:"product_code,omitempty" validate:"omitempty,len=2"`
		// TODO create MarcCodeInfo struct
		MarkCodeInfo                  any                   `json:"mark_code_info,omitempty" validate:"omitempty"`
		MarkMode                      string                `json:"mark_mode,omitempty" validate:"omitempty"`
		PaymentSubjectIndustryDetails ReceiptIndustryDetail `json:"payment_subject_industry_details,omitempty" validate:"omitempty"`
	}

	// MarkQuantity specifies fraction of a marked product (tag 1291 in 54-FZ). Must be specified if all of the following applies: FFD version 1.2 is used; payment is made for a marked product; the measure field has the piece value. Example: you're selling pencils by the piece. They're supplied in packages 100 pencils each with one marking code. To sell one pencil, enter 1 in numerator and 100 in denominator.
	//
	// Numerator: The number of products sold from one customer package (tag 1293 in 54-FZ). Cannot exceed the denominator.
	//
	// Denominator: The total number of products in the customer package (tag 1294 in 54-FZ).
	MarkQuantity struct {
		Numerator   int `json:"numerator" validate:"required,gte=1"`
		Denominator int `json:"denominator" validate:"required,gte=1"`
	}

	// ReceiptIndustryDetail specifies industry attribute of the receipt (tag 1261 in 54-FZ). Must be specified if FFD 1.2 is used.
	//
	// FederalID: ID of the federal executive authority (tag 1262 in 54-FZ). Pattern:(^00[1-9]{1}$)|(^0[1-6]{1}[0-9]{1}$)|(^07[0-3]{1}$)
	//
	// DocumentDate: Date of the incorporation document. Tag 1263 in 54-FZ. Specified in the ISO 8601 format.
	//
	// DocumentNumber: Number of the regulation issued by the federal executive authority prescribing how the "Industry attribute value" attribute must be filled in. Tag 1264 in 54-FZ.
	//
	// Value: Industry attribute value (tag 1265 in 54-FZ). Example:123/43
	ReceiptIndustryDetail struct {
		FederalID      string `json:"federal_id" validate:"required"`
		DocumentDate   string `json:"document_date" validate:"required,datetime"`
		DocumentNumber string `json:"document_number" validate:"required,max=32"`
		Value          string `json:"value" validate:"required,max=256"`
	}

	// Customer specifies user details. You should specify at least the basic contact information: email address (customer.email) or phone number (customer.phone).
	//
	// FullName: Name of the organization for companies, full name for sole proprietors and individuals. If the individual doesn't have a Tax Identification Number (INN), specify their passport information in this parameter. Maximum 256 characters. Online sales register that support this parameter: Orange Data, ATOL Online.
	//
	// INN: User's Tax Identification Number (INN) (10 or 12 digits). If the individual doesn't have an INN, specify their passport information in the full_name parameter. Online sales register that support this parameter: Orange Data, ATOL Online.
	//
	// Email: User's email address for sending the receipt. Required parameter if phone isn't specified.
	//
	// Phone: User's phone number for sending the receipt. Specified in the ITU-T E.164 format, for example, 79000000000. Required parameter if email isn't specified.
	Customer struct {
		FullName string `json:"full_name,omitempty" validate:"omitempty,max=256"`
		INN      string `json:"inn,omitempty" validate:"omitempty,numeric,len=10|len=12"`
		Email    string `json:"email,omitempty" validate:"omitempty,email"`
		// TODO check phone format e164 with + before numbers
		Phone string `json:"phone,omitempty" validate:"omitempty,e164"`
	}

	// ReceiptOperationalDetails specifies transaction attribute of the receipt (tag 1270 in 54-FZ). Must be specified if FFD 1.2 is used.
	//
	// OperationID: Transaction ID (tag 1271 in 54-FZ). From 0 to 255 characters.
	//
	// Value: Transaction details (tag 1272 in 54-FZ).
	//
	// CreatedAt: Time when the transaction was initiated (tag 1273 in 54-FZ). Formatted in accordance with UTC standart and specified in the ISO 8601. Example: 2017-11-03T11:52:31.827Z
	ReceiptOperationalDetails struct {
		OperationID int       `json:"operation_id" validate:"required,gte=0,lte=255"`
		Value       string    `json:"value" validate:"required,max=64"`
		CreatedAt   time.Time `json:"created_at" validate:"required"`
	}

	// Receiver specifies payment receiver's details specified when you want to add money to e-wallet, bank account, or phone balance .
	//
	// Type: Value: mobile_balance. Payment receiver code.
	//
	// AccountNumber: Bank account number. Format: 20 characters.
	//
	// BIC: Bank Identification Code (BIC) of the bank where the account is created. Format: 9 characters.
	//
	// Phone: Phone number where money should be added. Maximum 15 characters. Specified in the format ITU-T E.164. Example: 79000000000.
	Receiver struct {
		Type          string `json:"type" validate:"required,oneof=bank_account digital_wallet mobile_balance"`
		AccountNumber string `json:"account_number,omitempty" validate:"required_if=Type bank_account|required_if=Type digital_wallet"`
		BIC           string `json:"bic,omitempty" validate:"required_if=Type bank_account"`
		Phone         string `json:"phone,omitempty" validate:"required_if=Type mobile_balance"`
	}

	// Deal specifies the deal within which the payment is being carried out. Specified if you use Safe deal .
	//
	// ID: Deal ID.
	//
	// Settlements: Information about money distribution.
	Deal struct {
		ID          string           `json:"id" validate:"required,uuid"`
		Settlements []DealSettlement `json:"settlements" validate:"required"`
	}

	// DealSettlement specifies information about money distribution.
	//
	// Type: Transaction type. Fixed value: payout — payout to seller.
	//
	// Amount: Amount of seller’s remuneration.
	DealSettlement struct {
		Type   string `json:"type" validate:"required"`
		Amount `json:"amount" validate:"required"`
	}

	// Transfers specifies information about money distribution: the amounts of transfers and the stores to be transferred to. Specified if you use Split payments .
	//
	// AccountID: ID of the store in favor of which you're accepting the receipt. Provided by YooMoney, displayed in the Sellers section of your Merchant Profile (shopId column).
	//
	// Amount: Amount to be transferred to the store.
	//
	// PlatformFeeAmount: Commission for sold products or services charged in your favor.
	//
	// Description: Transaction description (up to 128 characters), which the seller will see in the YooMoney Merchant Profile. Example: "Marketplace order No. 72".
	//
	// Metadata: Any additional data you might require for processing payments (for example, your internal order ID), specified as a “key-value” pair and returned in response from YooMoney. Limitations: no more than 16 keys, no more than 32 characters in the key’s title, no more than 512 characters in the key’s value, data type is a string in the UTF-8 format.
	Transfers struct {
		AccountID         string `json:"account_id" validate:"required"`
		Amount            `json:"amount" validate:"required"`
		PlatformFeeAmount Amount `json:"platform_fee_amount,omitempty" validate:"omitempty"`
		Description       string `json:"description,omitempty" validate:"omitempty,max=128"`
		Metadata          any    `json:"metadata,omitempty" validate:"omitempty"`
	}

	// Airline specifies object containing the data for selling airline tickets. Used only for bank card payments.
	//
	// TicketNumber: Unique ticket number. If you already know the ticket number during payment creation, ticket_number is a required parameter. If you don't, specify booking_reference instead of ticket_number.
	//
	// BookingReference: Booking reference number, required if ticket_number is not specified.
	//
	// Passengers: List of passengers.
	//
	// Legs: List of flight legs.
	Airline struct {
		TicketNumber     string        `json:"ticket_number,omitempty" validate:"omitempty,min=1,max=150,numeric"`
		BookingReference string        `json:"booking_reference,omitempty" validate:"omitempty,min=1,max=20"`
		Passengers       []Passengers  `json:"passengers,omitempty" validate:"omitempty"`
		Legs             []AirlineLegs `json:"legs,omitempty" validate:"omitempty"`
	}

	// AirlineLegs specifies list of flight legs.
	//
	// DepartureAirport: Code of the departure airport according to IATA, for example, LED.
	//
	// DestinationAirport: Code of the arrival airport according to IATA, for example, AMS.
	//
	// DepartureDate: Departure date in the YYYY-MM-DD ISO 8601:2004 format.
	//
	// CarrierCode: Airline code according to IATA.
	AirlineLegs struct {
		DepartureAirport   string `json:"departure_airport" validate:"required,len=3,alpha"`
		DestinationAirport string `json:"destination_airport" validate:"required,len=3,alpha"`
		DepartureDate      string `json:"departure_date" validate:"required,datetime"`
		CarrierCode        string `json:"carrier_code,omitempty" validate:"omitempty,min=2,max=3"`
	}

	// Passengers specifies list of passengers.
	//
	// FirstName: Passenger's first name. Only use Latin characters, for example, SERGEI.
	//
	// LastName: Passenger's last name. Only use Latin characters, for example, IVANOV.
	Passengers struct {
		FirstName string `json:"first_name" validate:"required,min=1,max=64"`
		LastName  string `json:"last_name" validate:"required,min=1,max=64"`
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
		Currency `json:"currency" validate:"required,iso4217"`
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
		Type              ConfirmationType `json:"type" validate:"required"`
		ReturnURL         string           `json:"return_url,omitempty" validate:"required_if=Type redirect,required_if=Type mobile_application,omitempty,url"`
		ConfirmationURL   string           `json:"confirmation_url,omitempty" validate:"omitempty,required_if=Type redirect,omitempty,url"`
		ConfirmationToken string           `json:"confirmation_token,omitempty" validate:"omitempty,required_if=Type embedded"`
		Locale            string           `json:"locale,omitempty" validate:"omitempty,oneof=ru_RU en_US"`
		Enforce           bool             `json:"enforce,omitempty" validate:"omitempty"`
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

const (
	Untaxed    VatDataType = "untaxed"
	Calculated VatDataType = "calculated"
	Mixed      VatDataType = "mixed"
)

type ConfirmationType string

const (
	Embedded          ConfirmationType = "embedded"
	External          ConfirmationType = "external"
	MobileApplication ConfirmationType = "mobile_application"
	QR                ConfirmationType = "qr"
	Redirect          ConfirmationType = "redirect"
)

type Currency string

const (
	RUB Currency = "RUB"
	USD Currency = "USD"
)

type PaymentMethodType string

const (
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
		ID                   string            `json:"id" validate:"required,len=36,uuid"`
		Status               TransactionStatus `json:"status" validate:"required, oneof=pending waiting_for_capture succeeded canceled"`
		Paid                 bool              `json:"paid" validate:"required"`
		Amount               `json:"amount" validate:"required"`
		Confirmation         `json:"confirmation,omitempty" validate:"omitempty"`
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
		AccountID string `json:"account_id,omitempty" validate:"omitempty"`
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
