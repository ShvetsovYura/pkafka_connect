package consumer

import (
	"context"
	"encoding/json"
	"github/ShvetsovYura/pkafka_connect/internal/logger"
	"github/ShvetsovYura/pkafka_connect/internal/types"
	"log"
	"log/slog"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

// const sessionTimeout = 6000
// const fetchMinBytes = 1024

type KConsumer struct {
	consumer *kafka.Consumer

	opts types.ConsumerOpts
	// topics           []string
	// group            string
	// servers          string
	// timeout          time.Duration
	// enableAutoCommit bool
}

func NewKafkaConsumer(opts types.ConsumerOpts) *KConsumer {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":  opts.BootstrapServers,
		"group.id":           opts.Group,
		"session.timeout.ms": opts.SessionTimeoutMs,
		"enable.auto.commit": opts.EnableAutoCommit,
		"auto.offset.reset":  opts.AutoOffsetReset,
		"fetch.min.bytes":    opts.FetchMinBytes,
	})

	if err != nil {
		log.Fatalf("Невозможно создать консьюмер: %s\n", err)
	}

	logger.Log.Info("Консьюмер создан")

	return &KConsumer{
		consumer: c,
		opts:     opts,
	}
}
func (c *KConsumer) Run(ctx context.Context, queue chan types.Metric) {
	defer func() {
		c.consumer.Close()
	}()

	err := c.consumer.SubscribeTopics(c.opts.Topics, nil)

	if err != nil {
		log.Fatalf("Невозможно подписаться на топик: %s\n", err)
	}

	run := true
	for run {
		select {
		case <-ctx.Done():
			logger.Log.Info("Останавливается консьюмер")
			run = false
		default:
			msg, err := c.consumer.ReadMessage(c.opts.PollTimeout)
			if err == nil {
				logger.Log.Info("message", slog.String("partition", msg.TopicPartition.String()), slog.String("value", string(msg.Value)))
				var metrics map[string]types.Metric
				json.Unmarshal(msg.Value, &metrics)
				for _, v := range metrics {
					queue <- v
				}

			} else if !err.(kafka.Error).IsTimeout() {
				logger.Log.Error("Consumer error", slog.Any("err", err), slog.Any("msg", msg))
			}
		}
	}
}
