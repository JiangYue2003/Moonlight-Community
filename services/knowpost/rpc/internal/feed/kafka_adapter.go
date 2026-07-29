package feed

import (
	"context"

	"github.com/zeromicro/go-queue/kq"
)

// KafkaProducerAdapter 实现 KafkaProducer 接口，适配 kq
type KafkaProducerAdapter struct {
	pusher interface {
		PushWithKey(ctx context.Context, key, value string) error
	}
}

func NewKafkaProducerAdapter(pusher *kq.Pusher) *KafkaProducerAdapter {
	return &KafkaProducerAdapter{pusher: pusher}
}

func (k *KafkaProducerAdapter) SendMessage(ctx context.Context, _ string, key string, value []byte) error {
	return k.pusher.PushWithKey(ctx, key, string(value))
}
