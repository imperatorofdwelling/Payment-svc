package kafka

import (
	"context"
	"github.com/IBM/sarama"
	consumer "github.com/imperatorofdwelling/payment-svc/internal/handler/kafka/consumer/payment"
	"log"
)

const (
	ConsumerGroup = "payment-group"
	PaymentTopic  = "payment"
	PayoutTopic   = "payout"
)

type Consumer struct {
	Group sarama.ConsumerGroup
}

var ServerAddr = []string{"localhost:9094", "localhost:9095", "localhost:9096"}

func SetupKafkaConsumers() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		consumerGroup := newConsumerGroup()
		defer consumerGroup.Group.Close()

		paymentConsumer := consumer.NewPaymentConsumer()

		consumerGroup.Setup(ctx, consumerGroup.Group, paymentConsumer)

		defer cancel()
	}()
}

func newConsumerGroup() *Consumer {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(ServerAddr, ConsumerGroup, config)
	if err != nil {
		log.Fatalf("error creating consumer group %v", err)
	}
	return &Consumer{
		Group: consumerGroup,
	}
}

func (c *Consumer) Setup(ctx context.Context, group sarama.ConsumerGroup, hdl sarama.ConsumerGroupHandler) error {
	for {
		err := group.Consume(ctx, []string{PaymentTopic}, hdl)
		if err != nil {
			log.Printf("Error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return nil
		}
	}
}
