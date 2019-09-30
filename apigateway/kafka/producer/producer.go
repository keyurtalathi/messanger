package main

import (
	

	"github.com/Shopify/sarama"
)

func main() {

	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}
	// brokers := []string{"192.168.59.103:9092"}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		// Should not reach here
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			// Should not reach here
			panic(err)
		}
	}()

	topic := "important"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("sending new message 1"),
	}
	msg1 := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("sending new message 2"),
	}
	msg2 := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder("sending new message 3"),
	}
	partition, offset, err := producer.SendMessage(msg)
	partition, offset, err = producer.SendMessage(msg1)
	partition, offset, err = producer.SendMessage(msg2)

	if err != nil {
		panic(err)
	}

	//fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
