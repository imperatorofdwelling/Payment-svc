package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/eclipsemode/go-yookassa-sdk/yookassa"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-playground/validator/v10"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"go.uber.org/zap"
)

type PaymentConsumer struct {
	log                *zap.SugaredLogger
	yookassaPaymentSvc *yookassa.PaymentsSvc
	paymentSvc         service.IPaymentSvc
}

func (*PaymentConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*PaymentConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *PaymentConsumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		fmt.Println("Idempotency Key", string(msg.Key))
		fmt.Println("Hello WOOOOORLD: ", string(msg.Value))

		var payment yoomodel.Payment

		err := json.Unmarshal(msg.Value, &payment)
		if err != nil {
			return fmt.Errorf("error unmarshalling message: %w", err)
		}

		if err := v10.Validate.Struct(payment); err != nil {
			validationErr := err.(validator.ValidationErrors)
			return fmt.Errorf("error validating payment fields: %v", validationErr)
		}

		if payment.Status == "" {
			consumer.log.Error("invalid response from api")
			return fmt.Errorf("invalid response from api")
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewPaymentConsumer(log *zap.SugaredLogger, yookassaPaymentSvc *yookassa.PaymentsSvc, paymentSvc service.IPaymentSvc) *PaymentConsumer {
	return &PaymentConsumer{
		log, yookassaPaymentSvc, paymentSvc,
	}
}
