package main

import (
	"encoding/json"
	"fmt"

	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-project-api/internal/config"
	"github.com/ozoncp/ocp-project-api/internal/producer"
	"github.com/rs/zerolog/log"
)

func subscribe(topic string, consumer sarama.Consumer) error {
	partitionList, err := consumer.Partitions(topic) //get all partitions on the given topic
	if err != nil {
		return err
	}
	initialOffset := sarama.OffsetOldest //get offset for the oldest message on the topic

	for _, partition := range partitionList {
		pc, err := consumer.ConsumePartition(topic, partition, initialOffset)
		if err != nil {
			return err
		}

		for message := range pc.Messages() {
			messageReceived(message)
		}
	}

	return nil
}

func messageReceived(message *sarama.ConsumerMessage) {
	fmt.Printf("Analyzing message: %s\n", string(message.Value))
	var msg producer.EventMessage
	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		fmt.Printf("Error unmarshalling message: %s\n", err)
	}

	log.Info().Msgf("Got message: %v", msg.Body)
}

func main() {
	consumer, err := sarama.NewConsumer(config.Global.Producer.Brokers, nil)
	if err != nil {
		log.Fatal().Msgf("NewConsumer error: %v", err)
	}

	err = subscribe(config.Global.Producer.EventsTopic, consumer)
	if err != nil {
		log.Fatal().Msgf("Subscribe failed: %v", err)
	}
}
