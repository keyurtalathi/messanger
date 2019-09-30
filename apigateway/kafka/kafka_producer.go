package kafka

import (
	"fmt"

	"github.com/Shopify/sarama"
)

func Push_messages(message string) {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	brokers := []string{"localhost:9092"}

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		//fmt.Println("kafka create producer")
		panic(err)
	}

	defer func() {
		if err := producer.Close(); err != nil {
			//fmt.Println("kafka producer close")
			panic(err)
		}
	}()

	topic := "messages"
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(string(message)),
	}
	partition, offset, err := producer.SendMessage(msg)

	if err != nil {
		//fmt.Println("kafka send msg")
		panic(err)
	}

	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
}
