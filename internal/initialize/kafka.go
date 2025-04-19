package initialize

import (
	"vtuanjs/my-project/global"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

func InitKafka() {
	global.KafkaProducer = &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "test-my-project",
		Balancer: &kafka.LeastBytes{},
	}

	global.Logger.Info("Init Kafka success")
}

func CloseKafka() {
	if err := global.KafkaProducer.Close(); err != nil {
		global.Logger.Error("Close Kafka failed", zap.Error(err))
	}
}
