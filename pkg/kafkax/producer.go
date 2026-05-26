// Package kafkax 封装 segmentio/kafka-go 的 Producer / Consumer。
package kafkax

import (
	"context"
	"time"

	"github.com/segmentio/kafka-go"
)

// Producer 简单同步写入封装；上层可在外部加 channel + worker 实现异步。
type Producer struct {
	w *kafka.Writer
}

type ProducerOptions struct {
	LingerMs int
}

// NewProducer 构造异步批量 Writer；topic 在 WriteMessages 时按消息 Topic 字段决定。
func NewProducer(brokers []string) *Producer {
	return NewProducerWithOptions(brokers, ProducerOptions{})
}

// NewProducerWithOptions 支持配置批量等待时间等生产参数。
func NewProducerWithOptions(brokers []string, opts ProducerOptions) *Producer {
	linger := time.Duration(opts.LingerMs) * time.Millisecond
	if linger <= 0 {
		linger = 10 * time.Millisecond
	}
	return &Producer{
		w: &kafka.Writer{
			Addr:                   kafka.TCP(brokers...),
			Balancer:               &kafka.Hash{},
			BatchTimeout:           linger,
			RequiredAcks:           kafka.RequireAll,
			AllowAutoTopicCreation: true,
		},
	}
}

// Publish 发送一条消息；key 决定分区。
func (p *Producer) Publish(ctx context.Context, topic, key string, value []byte) error {
	return p.w.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Key:   []byte(key),
		Value: value,
	})
}

// Close 优雅关闭。
func (p *Producer) Close() error { return p.w.Close() }
