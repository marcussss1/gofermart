package kafka

import (
	"github.com/IBM/sarama"
)

type ProducerClient struct {
	Producer sarama.SyncProducer
}

func NewProducerClient(brokers []string) (*ProducerClient, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// Инициализация продюсера
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		return nil, err
	}

	return &ProducerClient{
		Producer: producer,
	}, nil
}

// Не забудьте добавить метод для закрытия соединений
func (r *ProducerClient) Close() error {
	if err := r.Producer.Close(); err != nil {
		return err
	}
	return nil
}
