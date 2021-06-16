package producer

import (
	"context"
	"encoding/json"
	"errors"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-project-api/internal/config"
	"github.com/rs/zerolog/log"
)

type Producer interface {
	IsAvailable() bool
	SendMessage(msg EventMessage) error
}

func NewProducer(ctx context.Context) (Producer, error) {
	saramaConfig := sarama.NewConfig()
	saramaConfig.Producer.Partitioner = sarama.NewHashPartitioner
	saramaConfig.Producer.RequiredAcks = sarama.WaitForAll
	saramaConfig.Producer.Return.Successes = true

	saramaClient, err := sarama.NewClient(config.Global.Producer.Brokers, saramaConfig)
	if err != nil {
		return nil, err
	}

	producer, err := sarama.NewSyncProducerFromClient(saramaClient)
	if err != nil {
		return nil, err
	}

	messages := make(chan *sarama.ProducerMessage, config.Global.Producer.Capacity)
	pr := &logProducer{
		producer:    producer,
		topic:       config.Global.Producer.EventsTopic,
		messages:    messages,
		client:      saramaClient,
		isAvailable: true,
	}
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
		Topic:     config.Global.Producer.PingTopic,
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
