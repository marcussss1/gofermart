package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	Consumer sarama.ConsumerGroup
	topics   []string
	handler  sarama.ConsumerGroupHandler
}

func NewKafkaConsumer(brokers []string, groupID string, topics []string, handler sarama.ConsumerGroupHandler) (*KafkaConsumer, error) {
	config := sarama.NewConfig()
	//config.Consumer.Group.Rebalance.Strategy = sarama.Bala
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("ошибка создания группы потребителей: %w", err)
	}

	return &KafkaConsumer{
		Consumer: consumer,
		topics:   topics,
		handler:  handler,
	}, nil
}

func (c *KafkaConsumer) Consume(ctx context.Context) error {
	for {
		err := c.Consumer.Consume(ctx, c.topics, c.handler)
		if err != nil {
			return fmt.Errorf("ошибка при потреблении: %w", err)
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *KafkaConsumer) Close() error {
	return c.Consumer.Close()
}

// Пример обработчика сообщений
type ExampleHandler struct{}

func (h *ExampleHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ExampleHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (h *ExampleHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Получено сообщение: тема = %s, раздел = %d, смещение = %d, ключ = %s, значение = %s\n",
			message.Topic, message.Partition, message.Offset, string(message.Key), string(message.Value))
		session.MarkMessage(message, "")
	}
	return nil
}
