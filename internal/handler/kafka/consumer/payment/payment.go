package consumer

import (
	jsonDefault "encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/eclipsemode/go-yookassa-sdk/yookassa"
	yoomodel "github.com/eclipsemode/go-yookassa-sdk/yookassa/model"
	"github.com/go-playground/validator/v10"
	kafka "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer"
	v10 "github.com/imperatorofdwelling/payment-svc/internal/lib/validator"
	"github.com/imperatorofdwelling/payment-svc/internal/service"
	"github.com/imperatorofdwelling/payment-svc/pkg/json"
	"go.uber.org/zap"
)

type PaymentConsumer struct {
	// TODO fix bug with using log (crushing the app)
	log                *zap.SugaredLogger
	yookassaPaymentSvc *yookassa.PaymentsSvc
	paymentSvc         service.IPaymentSvc
	kafkaProducer      *kafka.Producer
}

func (*PaymentConsumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*PaymentConsumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (c *PaymentConsumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	const op = "kafka.consumer.payment.ConsumeClaim"

	for msg := range claim.Messages() {
		requestID := string(msg.Key)

		var payment yoomodel.Payment

		err := jsonDefault.Unmarshal(msg.Value, &payment)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}

		if err := v10.Validate.Struct(payment); err != nil {
			validationErr := err.(validator.ValidationErrors)
			return fmt.Errorf("%s: %w", "error validating payment fields", validationErr)
		}

		paymentRes, err := c.yookassaPaymentSvc.CreatePayment(&payment, requestID)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}

		var newPayment yoomodel.Payment

		err = json.Read(paymentRes.Body, &newPayment)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}

		if newPayment.Status == "" {
			return fmt.Errorf("%s: %s", op, "error creating payment")
		}

		err = c.paymentSvc.CreatePayment(sess.Context(), &newPayment)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
		}

		err = c.kafkaProducer.SendMessage(kafka.PaymentResponseTopic, requestID, newPayment)
		if err != nil {
			return fmt.Errorf("%s: %s", op, err.Error())
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
