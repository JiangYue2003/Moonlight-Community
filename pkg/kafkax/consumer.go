package kafkax

import (
	"context"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/logx"
)

// HandlerFunc 处理一条消息；返回 nil 视为成功并提交 offset。
type HandlerFunc func(ctx context.Context, m kafka.Message) error

// ConsumerConfig consumer group 参数。
type ConsumerConfig struct {
	Brokers  []string
	Topic    string
	GroupId  string
	MinBytes int
	MaxBytes int
	// StartOffset 仅在 consumer group 无已提交 offset 时生效：
	// - "latest"   -> kafka.LastOffset
	// - "earliest" -> kafka.FirstOffset
	StartOffset string
	// MaxRetries 超过此次数后把消息写入 DlqTopic 并跳过。0 = 无限重试（向后兼容）。
	MaxRetries int
	// DlqTopic 死信队列 topic。空串 = 只记 ERROR 日志，不写 Kafka。
	DlqTopic string
}

const (
	backoffBase = 100 * time.Millisecond
	backoffMax  = 30 * time.Second
)

// RunConsumer 启动一个阻塞式 consumer，直到 ctx 被取消。
// 内部使用手动提交：handler 成功才提交 offset，保证至少一次语义。
// handler 持续失败时以指数退避重试（100ms → 200ms → ... → 30s），避免 CPU 空转。
// 超过 MaxRetries（>0）后把原始消息写入 DlqTopic 并提交 offset 继续消费。
func RunConsumer(ctx context.Context, cfg ConsumerConfig, h HandlerFunc) error {
	if cfg.MinBytes == 0 {
		cfg.MinBytes = 1
	}
	if cfg.MaxBytes == 0 {
		cfg.MaxBytes = 10 << 20 // 10MB
	}
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        cfg.Brokers,
		Topic:          cfg.Topic,
		GroupID:        cfg.GroupId,
		MinBytes:       cfg.MinBytes,
		MaxBytes:       cfg.MaxBytes,
		CommitInterval: 0, // 手动提交
		StartOffset:    parseStartOffset(cfg.StartOffset),
	})
	defer r.Close()

	var dlqWriter *kafka.Writer
	if cfg.DlqTopic != "" {
		dlqWriter = &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    cfg.DlqTopic,
			Balancer: &kafka.LeastBytes{},
		}
		defer dlqWriter.Close()
	}

	var retries int
	for {
		m, err := r.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		if err := h(ctx, m); err != nil {
			retries++

			if cfg.MaxRetries > 0 && retries >= cfg.MaxRetries {
				logx.Errorf("kafkax: message exceeded MaxRetries=%d, sending to DLQ. topic=%s partition=%d offset=%d err=%v",
					cfg.MaxRetries, m.Topic, m.Partition, m.Offset, err)
				if dlqWriter != nil {
					dlqMsg := kafka.Message{
						Key:   m.Key,
						Value: m.Value,
						Headers: append(m.Headers,
							kafka.Header{Key: "X-Retry-Count", Value: []byte(fmt.Sprintf("%d", retries))},
							kafka.Header{Key: "X-Original-Topic", Value: []byte(m.Topic)},
						),
					}
					if werr := dlqWriter.WriteMessages(ctx, dlqMsg); werr != nil {
						logx.Errorf("kafkax: failed to write DLQ message: %v", werr)
					}
				}
				// 提交 offset，跳过毒丸消息
				retries = 0
				if cerr := r.CommitMessages(ctx, m); cerr != nil {
					if errors.Is(cerr, context.Canceled) {
						return nil
					}
					return cerr
				}
				continue
			}

			delay := backoffBase * (1 << min(retries-1, 8)) // 最多左移 8 位 = 25.6s
			if delay > backoffMax {
				delay = backoffMax
			}
			logx.Errorf("kafkax handler error (retry=%d, backoff=%s): %v", retries, delay, err)
			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil
			}
			continue
		}

		retries = 0 // 成功后重置退避计数
		if err := r.CommitMessages(ctx, m); err != nil {
			if errors.Is(err, context.Canceled) {
				return nil
			}
			return err
		}
	}
}

func parseStartOffset(v string) int64 {
	switch v {
	case "earliest", "first":
		return kafka.FirstOffset
	default:
		return kafka.LastOffset
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
