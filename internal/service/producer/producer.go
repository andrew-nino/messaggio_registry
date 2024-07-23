package service

import (
	"context"
	"encoding/json"
	"os"
	"strconv"
	"strings"

	"github.com/andrew-nino/messaggio/config"
	"github.com/andrew-nino/messaggio/internal/domain/models"

	"github.com/segmentio/kafka-go"
	"github.com/sirupsen/logrus"
)

type producer struct {
	log         *logrus.Logger
	kafkaWriter *kafka.Writer
}

func New(log *logrus.Logger, cfg *config.Kafka) *producer {

	if len(cfg.Brokers) == 0 {
		cfg.Brokers = os.Getenv("KAFKA_BROKERS")
	}
	if len(cfg.Topic) == 0 {
        cfg.Topic = os.Getenv("KAFKA_TOPIC")
    }

	addrs := strings.Split(cfg.Brokers, ",")
	topic := cfg.Topic

	return &producer{
		log: log,
		kafkaWriter: &kafka.Writer{
			Addr:     kafka.TCP(addrs...),
			Topic:    topic,
			Balancer: &kafka.Hash{},
		},
	}
}

func (p *producer) SendToBroker(id int, client models.Client) error {

	msg, err := json.Marshal(client)
	if err != nil {
		panic(err)
	}

	payload := kafka.Message{
		Key:   []byte(strconv.Itoa(id)),
		Value: []byte(msg),
	}

	err = p.kafkaWriter.WriteMessages(context.Background(), payload)
	if err != nil {
		p.log.Fatal("failed to write messages:", err)
	}

	return nil
}
