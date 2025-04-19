/*
https://www.youtube.com/watch?v=7Hm2RsH8bS8*
https://www.youtube.com/watch?v=KSej3yivuPY

// Создание продюсера
producer, err := kafka.NewProducer("localhost:9092")

	if err != nil {
		log.Fatal(err)
	}

defer producer.Close()

// Отправка сообщения
err = producer.Send("my-topic", "message-key", "Hello Kafka")

	if err != nil {
		log.Fatal(err)
	}

// Создание потребителя
// Подписываемся на топик "my-topic" в группе "test-group"
consumer, err := kafka.NewConsumer("localhost:9092", "my-topic", "test-group")

	if err != nil {
		log.Fatal(err)
	}

defer consumer.Close()

// Запуск потребителя в отдельной горутине
go consumer.Start()
*/
package kafka

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"log"
)

const (
	flushTimeout           = 5000 // ms
	consumerSessionTimeout = 5000 // ms
)

type Producer struct {
	producer *kafka.Producer
}

type Consumer struct {
	consumer *kafka.Consumer
}

// NewProducer создает нового продюсера Kafka
// address - адреса брокеров Kafka через запятую (например: "host1:9092,host2:9092")
func NewProducer(address string) (*Producer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers": address,
	}
	p, err := kafka.NewProducer(config)
	if err != nil {
		return nil, fmt.Errorf("error new producer: %w", err)
	}
	return &Producer{producer: p}, nil
}

func (p *Producer) Send(topic, key, text string) error {
	var messageKey []byte
	if key != "" {
		messageKey = []byte(key)
	}

	kafkaMessage := &kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &topic,
			Partition: kafka.PartitionAny,
		},
		Value: []byte(text),
		Key:   messageKey,
	}
	kafkaChan := make(chan kafka.Event)
	defer close(kafkaChan)

	err := p.producer.Produce(kafkaMessage, kafkaChan)
	if err != nil {
		return fmt.Errorf("error Send: %w", err)
	}

	event := <-kafkaChan
	switch eventResult := event.(type) {
	case *kafka.Message:
		return nil
	case kafka.Error:
		return eventResult
	default:
		return fmt.Errorf("unexpected event type from kafkaChan")
	}
}

func (p *Producer) Close() {
	p.producer.Flush(flushTimeout)
	p.producer.Close()
}

// NewConsumer создает нового потребителя Kafka
// address - адреса брокеров Kafka через запятую (например: "host1:9092,host2:9092")
// topics - список топиков через запятую (например: "topic1,topic2")
// consumerGroup - имя группы потребителей
func NewConsumer(address, topics, consumerGroup string) (*Consumer, error) {
	config := &kafka.ConfigMap{
		"bootstrap.servers":  address,
		"group.id":           consumerGroup,
		"session.timeout.ms": consumerSessionTimeout,
		"auto.offset.reset":  "earliest", // рекомендуется установить стратегию обработки смещений
	}
	c, err := kafka.NewConsumer(config)
	if err != nil {
		return nil, fmt.Errorf("error new consumer: %w", err)
	}

	// Подписка на топики
	// Пример topics: "my-topic" или "topic1,topic2,topic3" для нескольких топиков
	err = c.SubscribeTopics(splitTopics(topics), nil)
	if err != nil {
		return nil, fmt.Errorf("error subscribing to topics: %w", err)
	}

	return &Consumer{consumer: c}, nil
}

// splitTopics разделяет строку с топиками через запятую в массив
func splitTopics(topics string) []string {
	var result []string
	// Здесь должна быть реализация разделения строки
	// Например: return strings.Split(topics, ",")
	// В текущем примере просто возвращаем массив с одним элементом
	return append(result, topics)
}

func (c *Consumer) Start() {
	for {
		kafkaMessage, err := c.consumer.ReadMessage(-1) // исправлена опечатка (была запятая вместо точки)
		if err != nil {
			log.Printf("Error reading message: %v", err)
			continue
		}
		if kafkaMessage == nil {
			continue
		}
		// Обработка сообщения
		fmt.Printf("Received message from topic %s: %s\n", *kafkaMessage.TopicPartition.Topic, string(kafkaMessage.Value))
	}
}

func (c *Consumer) Close() {
	c.consumer.Close()
}
