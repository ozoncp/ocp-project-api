package producer

import (
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	SendMessage(msg EventMessage) error
}

const (
	capacity = 256
)

var brokers = []string{"127.0.0.1:9094"}

func NewProducer(ctx context.Context, topic string) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	messages := make(chan *sarama.ProducerMessage, capacity)
	pr := &logProducer{producer: producer, topic: topic, messages: messages}
	go pr.handleMessages(ctx)

	return pr, nil
}

type logProducer struct {
	producer sarama.SyncProducer
	topic    string
	messages chan *sarama.ProducerMessage
}

func (pr *logProducer) SendMessage(msg EventMessage) error {
	b, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	pr.messages <- &sarama.ProducerMessage{
		Topic:     pr.topic,
		Partition: -1,
		Key:       sarama.StringEncoder(pr.topic),
		Value:     sarama.StringEncoder(b),
	}

	return nil
}

func (pr *logProducer) handleMessages(ctx context.Context) {
	for {
		select {
		case msg := <-pr.messages:
			_, _, err := pr.producer.SendMessage(msg)
			if err != nil {
				log.Error().Msgf("producer.SendMessage(...) returns error: %v", err)
			}
		case <-ctx.Done():
			close(pr.messages)
			return
		}
	}
}
