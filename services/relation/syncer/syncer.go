// relation-syncer：消费 canal-outbox topic，把 outbox 事件落地到 follower 反查表 + ZSet + usercounter。
package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	kafka "github.com/segmentio/kafka-go"

	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/processor"
	"github.com/zhiguang/zhiguang-go/services/relation/syncer/internal/svc"
)

var configFile = flag.String("f", "etc/syncer.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	sc := svc.NewServiceContext(c)
	defer sc.Close()

	proc := processor.New(sc)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		err := kafkax.RunConsumer(ctx, kafkax.ConsumerConfig{
			Brokers: c.Kafka.Brokers,
			Topic:   c.Kafka.Topic,
			GroupId: c.Kafka.GroupId,
			StartOffset: c.Kafka.AutoOffsetReset,
		}, func(ctx context.Context, m kafka.Message) error {
			return proc.Handle(ctx, m.Value)
		})
		if err != nil {
			logx.Errorf("syncer consumer exit: %v", err)
			cancel()
		}
	}()

	logx.Info("relation-syncer started")

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	<-sigCh
	logx.Info("relation-syncer shutting down")
	cancel()
}
