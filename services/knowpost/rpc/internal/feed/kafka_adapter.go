package feed

import (
	"context"

	"github.com/zeromicro/go-queue/kq"
)

// KafkaProducerAdapter 实现 KafkaProducer 接口，适配 kq
type KafkaProducerAdapter struct {
	pusher *kq.Pusher
}

func NewKafkaProducerAdapter(pusher *kq.Pusher) *KafkaProducerAdapter {
	return &KafkaProducerAdapter{pusher: pusher}
}

func (k *KafkaProducerAdapter) SendMessage(ctx context.Context, topic string, key string, value []byte) error {
	return k.pusher.Push(ctx, string(value))
}
