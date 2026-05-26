package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetDetailPublicLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	id     int64
}

func NewGetDetailPublicLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetDetailPublicLogic {
	return &GetDetailPublicLogic{Logger: logx.WithContext(ctx), ctx: ctx, svcCtx: svcCtx}
}
func (l *GetDetailPublicLogic) WithId(id int64) *GetDetailPublicLogic { l.id = id; return l }

// GetDetailPublic 公开详情接口；可选鉴权（未登录也能访问公开帖子）。
// viewer_id 用于私有帖子的可见性判定（私有帖只有作者本人能看）。
func (l *GetDetailPublicLogic) GetDetailPublic() (*types.KnowPostDetail, error) {
	viewerId, _ := ctxdata.GetUserId(l.ctx)
	r, err := l.svcCtx.KnowPostRpc.GetDetail(l.ctx, &pb.GetDetailReq{Id: l.id, ViewerId: viewerId})
	if err != nil {
		return nil, err
	}
	return detailFromPb(l.ctx, l.svcCtx, r), nil
}
