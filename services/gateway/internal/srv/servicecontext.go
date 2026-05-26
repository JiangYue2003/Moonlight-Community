package srv

import (
	"log"

	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/config"
	counterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/counter"
	usercounterclient "github.com/zhiguang/zhiguang-go/services/counter/rpc/client/usercounter"
	knowpostclient "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/client/knowpost"
	llmclient "github.com/zhiguang/zhiguang-go/services/llm/rpc/client/llm"
	relationclient "github.com/zhiguang/zhiguang-go/services/relation/rpc/client/relation"
	searchclient "github.com/zhiguang/zhiguang-go/services/search/rpc/client/search"
	storageclient "github.com/zhiguang/zhiguang-go/services/storage/rpc/client/storage"
	authclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/auth"
	userclient "github.com/zhiguang/zhiguang-go/services/user/rpc/client/user"
)

type ServiceContext struct {
	Config      config.Config
	Oss         *ossx.Client
	AuthRpc     authclient.Auth
	UserRpc     userclient.User
	StorageRpc  storageclient.Storage
	KnowPostRpc knowpostclient.KnowPost
	RelationRpc relationclient.Relation
	CounterRpc  counterclient.Counter
	UserCounterRpc usercounterclient.UserCounter
	SearchRpc   searchclient.Search
	LlmRpc      llmclient.Llm
}

func NewServiceContext(c config.Config) *ServiceContext {
	ossCli, err := ossx.New(c.Oss)
	if err != nil {
		log.Fatalf("gateway: oss init: %v", err)
	}
	authCli := zrpc.MustNewClient(c.AuthRpc)
	userCli := zrpc.MustNewClient(c.UserRpc)
	return &ServiceContext{
		Config:         c,
		Oss:            ossCli,
		AuthRpc:        authclient.NewAuth(authCli),
		UserRpc:        userclient.NewUser(userCli),
		StorageRpc:     storageclient.NewStorage(zrpc.MustNewClient(c.StorageRpc)),
		KnowPostRpc:    knowpostclient.NewKnowPost(zrpc.MustNewClient(c.KnowPostRpc)),
		RelationRpc:    relationclient.NewRelation(zrpc.MustNewClient(c.RelationRpc)),
		CounterRpc:     counterclient.NewCounter(zrpc.MustNewClient(c.CounterRpc)),
		UserCounterRpc: usercounterclient.NewUserCounter(zrpc.MustNewClient(c.UserCounterRpc)),
		SearchRpc:      searchclient.NewSearch(zrpc.MustNewClient(c.SearchRpc)),
		LlmRpc:         llmclient.NewLlm(zrpc.MustNewClient(c.LlmRpc)),
	}
}
