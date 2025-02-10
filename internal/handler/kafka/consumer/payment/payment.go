package consumer

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/eclipsemode/go-yookassa-sdk/yookassa"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-playground/validator/v10"
	kafka "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"go.uber.org/zap"
)

type PaymentConsumer struct {
	log                *zap.SugaredLogger
	yookassaPaymentSvc *yookassa.PaymentsSvc
	paymentSvc         service.IPaymentSvc
	kafkaProducer      *kafka.Producer
}

func (*PaymentConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*PaymentConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c *PaymentConsumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		requestID := string(msg.Key)

		var payment yoomodel.Payment

		err := json.Unmarshal(msg.Value, &payment)
		if err != nil {
			return fmt.Errorf("error unmarshalling message: %w", err)
		}

		if err := v10.Validate.Struct(payment); err != nil {
			validationErr := err.(validator.ValidationErrors)
			return fmt.Errorf("error validating payment fields: %v", validationErr)
		}

		err = c.kafkaProducer.SendMessage(kafka.PaymentResponseTopic, requestID, "GUTEN TAG BACK")
		if err != nil {
			return fmt.Errorf("error sending message: %w", err)
		}

		sess.MarkMessage(msg, "")
	}
	return nil
}

func NewPaymentConsumer(log *zap.SugaredLogger, yookassaPaymentSvc *yookassa.PaymentsSvc, paymentSvc service.IPaymentSvc, kafkaProducer *kafka.Producer) *PaymentConsumer {
	return &PaymentConsumer{
		log, yookassaPaymentSvc, paymentSvc, kafkaProducer,
	}
}
