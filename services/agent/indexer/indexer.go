package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/pkg/kafkax"
	"github.com/zhiguang/zhiguang-go/services/agent/indexer/internal/config"
	"github.com/zhiguang/zhiguang-go/services/agent/indexer/internal/processor"
	"github.com/zhiguang/zhiguang-go/services/agent/indexer/internal/svc"
)

var configFile = flag.String("f", "etc/agent-indexer.yaml", "config file")

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
	go processor.RunCompensation(ctx, proc, time.Duration(c.CompensateMinutes)*time.Minute)

	go func() {
		err := kafkax.RunConsumer(ctx, kafkax.ConsumerConfig{
			Brokers: c.Kafka.Brokers,
			Topic:   c.Kafka.Topic,
			GroupId: c.Kafka.GroupId,
		}, func(ctx context.Context, m kafka.Message) error {
			return proc.Handle(ctx, m)
		})
		if err != nil {
			logx.Errorf("agent-indexer consumer exit: %v", err)
			cancel()
		}
	}()

	logx.Info("agent-indexer started")
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	cancel()
	logx.Info("agent-indexer shutting down")
}
