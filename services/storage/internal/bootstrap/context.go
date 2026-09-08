package bootstrap

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zhiguang/zhiguang-go/pkg/ossx"
	knowmodel "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
	"github.com/zhiguang/zhiguang-go/services/storage/internal/application"
)

type ServiceContext struct {
	Storage *application.Service
}

func NewServiceContext(c Config) *ServiceContext {
	ossClient, err := ossx.New(c.Oss)
	if err != nil {
		panic(err)
	}
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	knowPosts := knowmodel.NewKnowPostsModel(conn, c.CacheRedis)
	return &ServiceContext{Storage: application.NewService(knowPosts, ossClient)}
}
