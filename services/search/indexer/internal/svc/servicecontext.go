package svc

import (
	"context"
	"log"
	"net/http"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/zhiguang/zhiguang-go/pkg/esx"
	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"github.com/zhiguang/zhiguang-go/services/search/indexer/internal/config"
)

type ServiceContext struct {
	Config config.Config

	Es          *esx.Client
	KnowPostRpc knowpostpb.KnowPostClient
	Redis       goredis.UniversalClient
	HttpClient  *http.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	es, err := esx.New(esx.Config{
		Addrs:    c.Es.Addrs,
		Username: c.Es.Username,
		Password: c.Es.Password,
		Timeout:  time.Duration(c.Es.TimeoutMs) * time.Millisecond,
	})
	if err != nil {
		log.Fatalf("search-indexer: init esx: %v", err)
	}
	if err := es.EnsureIndex(context.Background(), c.ContentIndex, esx.ContentMapping()); err != nil {
		log.Fatalf("search-indexer: ensure index: %v", err)
	}
	rdb := goredis.NewUniversalClient(&goredis.UniversalOptions{
		Addrs:    []string{c.Redis.Host},
		Password: c.Redis.Pass,
	})
	return &ServiceContext{
		Config:      c,
		Es:          es,
		KnowPostRpc: knowpostpb.NewKnowPostClient(zrpc.MustNewClient(c.KnowPostRpc).Conn()),
		Redis:       rdb,
		HttpClient:  &http.Client{Timeout: time.Duration(c.HttpFetchTimeoutMs) * time.Millisecond},
	}
}

func (s *ServiceContext) Close() {
	if s.Redis != nil {
		_ = s.Redis.Close()
	}
}
