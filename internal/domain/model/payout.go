package model

import (
	"github.com/google/uuid"
	"time"
)

type (
	// Payout object contains all the information about the payout that is relevant at the current time. It is generated when creating a payout and comes in response to any request related to payments.
	//
	// ID: payment id.
	//
	// PayoutToken: Tokenized payment data. For example, a synonym for a bank card. Required parameter if payout_destination_data or payment_method_id is not passed.
	//
	// Description: Description of the transaction (maximum 128 characters). For example: "Payment under contract 37".
	//
	// Metadata: Any additional data that you need to work with (for example, your internal order ID). They are transmitted as a set of key-value pairs and returned in a response from Yandex. Restrictions: a maximum of 16 keys, the key name is no more than 32 characters, the key value is no more than 512 characters, and the data type is a string in UTF-8 format.
	//
	// PayoutDestination: the seller's means of payment to which Kassa transfers the payment.
	//
	// CreatedAt: The time when the payout was created. It is specified in UTC and transmitted in the ISO 8601 format. Example: 2017-11-03T11:52:31.827Z
	//
	// Test: Indicates a test operation.
	Payout struct {
		ID                uuid.UUID `json:"id,omitempty" validate:"omitempty,uuid"`
		Amount            `json:"amount"`
		PayoutToken       string             `json:"payout_token,omitempty"`
		Description       string             `json:"description,omitempty"`
		Metadata          any                `json:"metadata,omitempty"`
		Status            *TransactionStatus `json:"status,omitempty" validate:"omitempty"`
		PayoutDestination *PayoutDestination `json:"payout_destination,omitempty" validate:"omitempty"`
		CreatedAt         *time.Time         `json:"created_at,omitempty" validate:"omitempty,datetime"`
		Test              bool               `json:"test,omitempty"`
	}

	// PayoutDestination defines the seller's means of payment to which Kassa transfers the payment.
	//
	// Card: Bank card details.
	//
	// BankID: The ID of the member of the SBP bank or payment service connected to the service.
	//
	// Phone: The phone number to which the payee's account is linked in the joint venture participant's system. It is specified in the ITU-T E.164 format, for example 79000000000.
	//
	// RecipientChecked: Verification of the payee : true — the payment was made with verification of the payee, false — the payment was made without verification of the payee.
	//
	// AccountNumber: Webmoney wallet number, for example 41001614575714. The length is from 11 to 33 digits.
	PayoutDestination struct {
		Type             string       `json:"type" validate:"required,oneof=bank_card sbp yoo_money"`
		Card             BankCardData `json:"card,omitempty" validate:"omitempty,required_if=Type bank_card"`
		BankID           string       `json:"bank_id,omitempty" validate:"omitempty,required_if=Type sbp"`
		Phone            string       `json:"phone,omitempty" validate:"omitempty,required_if=Type sbp,e164"`
		RecipientChecked bool         `json:"recipient_checked,omitempty" validate:"omitempty,required_if=Type sbp"`
		AccountNumber    string       `json:"account_number,omitempty" validate:"omitempty,required_if=Type yoo_money,min=11,max=33,numeric"`
	}
)
