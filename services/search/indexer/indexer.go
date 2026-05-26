// search-indexer：消费 canal-outbox（aggregate_type=knowpost），UPSERT zhiguang_content_index。
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
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/processor"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/svc"
)

var configFile = flag.String("f", "etc/indexer.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	sc := svc.NewServiceContext(c)
	defer sc.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// 启动 backfill（异步；不阻塞 Kafka 消费）
	go func() {
		if err := processor.NewBackfiller(sc).Run(ctx); err != nil {
			logx.Errorf("backfill: %v", err)
		}
	}()

	proc := processor.New(sc)
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
			logx.Errorf("indexer consumer exit: %v", err)
			cancel()
		}
	}()

	logx.Info("search-indexer started")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	logx.Info("search-indexer shutting down")
	cancel()
}
