// rag-indexer：消费 canal-outbox（aggregate_type=knowpost），切块/嵌入/写 ES dense_vector。
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
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/processor"
	"github.com/zhiguang/zhiguang-go/services/llm/ragindexer/internal/svc"
)

var configFile = flag.String("f", "etc/ragindexer.yaml", "config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)

	sc := svc.NewServiceContext(c)
	defer sc.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	proc := processor.New(sc)

	go func() {
		if err := processor.NewBackfiller(sc).Run(ctx, proc); err != nil {
			logx.Errorf("backfill: %v", err)
		}
	}()

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
			logx.Errorf("rag-indexer consumer exit: %v", err)
			cancel()
		}
	}()

	logx.Info("rag-indexer started")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	cancel()
	logx.Info("rag-indexer shutting down")
}
