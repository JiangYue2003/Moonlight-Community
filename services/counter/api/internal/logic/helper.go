package logic

import (
	"context"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/counter/api/internal/types"
	"github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

// dispatchToggle 共享 like/unlike/fav/unfav 的实现路径。
// metric ∈ {"like","fav"}; add 表示设/取。
func dispatchToggle(ctx context.Context, svcCtx *svc.ServiceContext, req *types.ActionReq, metric string, add bool) (*types.ActionResp, error) {
	uid, _ := ctxdata.GetUserId(ctx)
	resp, err := svcCtx.CounterRpc.Toggle(ctx, &counter.ToggleReq{
		EntityType: req.EntityType,
		EntityId:   req.EntityId,
		Metric:     metric,
		UserId:     uid,
		Add:        add,
	})
	if err != nil {
		return nil, err
	}
	out := &types.ActionResp{Changed: resp.Changed}
	switch metric {
	case "like":
		out.Liked = add
	case "fav":
		out.Faved = add
	}
	return out, nil
}
