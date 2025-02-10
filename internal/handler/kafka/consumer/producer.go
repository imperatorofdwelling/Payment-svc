package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"go.uber.org/zap"
)

const (
	PaymentReqTopic      = "payment"
	PayoutReqTopic       = "payout"
	PaymentResponseTopic = "payment-response"
)

type Producer struct {
	sync sarama.SyncProducer
	log  *zap.SugaredLogger
}

func NewKafkaProducer(log *zap.SugaredLogger) *Producer {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(ServerAddr, config)
	if err != nil {
		log.Errorf("Failed to start Sarama producer: %s", err)
	}

	return &Producer{
		sync: producer,
	}
}

func (p *Producer) Close() error {
	return p.sync.Close()
}

func (p *Producer) SendMessage(topic string, key string, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshalling message: %s", err.Error())
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),
		Value: sarama.ByteEncoder(messageBytes),
	}

	_, _, err = p.sync.SendMessage(msg)

	return err
}
