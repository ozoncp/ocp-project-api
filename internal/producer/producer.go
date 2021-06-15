package producer

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	IsAvailable() bool
	SendMessage(msg EventMessage) error
}

const (
	capacity = 256
	topic    = "events"
)

var brokers = []string{"127.0.0.1:9094"}

func NewProducer(ctx context.Context) (Producer, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewHashPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	saramaClient, err := sarama.NewClient(brokers, config)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(saramaClient)
	if err != nil {
		return nil, err
	}

	messages := make(chan *sarama.ProducerMessage, capacity)
	pr := &logProducer{producer: producer, topic: topic, messages: messages, client: saramaClient, isAvailable: true}
	go pr.handleMessages(ctx)

	return pr, nil
}

type logProducer struct {
	producer    sarama.SyncProducer
	topic       string
	messages    chan *sarama.ProducerMessage
	client      sarama.Client
	mutex       sync.Mutex
	isAvailable bool
}

func (pr *logProducer) SendMessage(msg EventMessage) error {
	if !pr.IsAvailable() {
		return errors.New("Kafka is not available")
	}

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
				pr.messages <- msg
				log.Error().Msgf("producer.SendMessage(...) returns error: %v", err)
			}
		case <-ctx.Done():
			close(pr.messages)
			pr.producer.Close()
			return
		}
	}
}

func (pr *logProducer) IsAvailable() bool {
	pr.mutex.Lock()
	defer pr.mutex.Unlock()

	msg := &sarama.ProducerMessage{
		Topic:     "ping",
		Partition: -1,
		Value:     sarama.StringEncoder("{}"),
	}

	_, _, err := pr.producer.SendMessage(msg)
	if err != nil {
		pr.isAvailable = false
	} else {
		pr.isAvailable = true
	}

	return pr.isAvailable
}
