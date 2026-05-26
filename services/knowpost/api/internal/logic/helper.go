package logic

import (
	"context"
	"strconv"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/api/internal/types"
	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	userpb "github.com/zhiguang/zhiguang-go/services/user/rpc/user"
)

func detailFromPb(ctx context.Context, svcCtx *svc.ServiceContext, d *pb.KnowPostDetail) *types.KnowPostDetail {
	if d == nil {
		return &types.KnowPostDetail{}
	}
	author := getUserBrief(ctx, svcCtx, d.CreatorId)
	like, fav := getCounts(ctx, svcCtx, d.Id)
	liked, faved := getMarked(ctx, svcCtx, d.Id)
	return &types.KnowPostDetail{
		Id:             d.Id,
		Title:          d.Title,
		Description:    d.Description,
		ContentUrl:     d.ContentUrl,
		Images:         d.ImgUrls,
		Tags:           d.Tags,
		AuthorId:       d.CreatorId,
		AuthorAvatar:   author.Avatar,
		AuthorNickname: author.Nickname,
		AuthorTagJson:  author.TagsJson,
		LikeCount:      like,
		FavoriteCount:  fav,
		Liked:          liked,
		Faved:          faved,
		IsTop:          d.IsTop,
		Visible:        d.Visible,
		Type:           d.Type,
		PublishTime:    d.PublishTime,
	}
}

func feedPageFromPb(ctx context.Context, svcCtx *svc.ServiceContext, p *pb.FeedPage) *types.FeedPage {
	if p == nil {
		return &types.FeedPage{}
	}
	items := make([]types.FeedItem, 0, len(p.Items))
	for _, it := range p.Items {
		author := getUserBrief(ctx, svcCtx, it.CreatorId)
		cover := ""
		if len(it.ImgUrls) > 0 {
			cover = it.ImgUrls[0]
		}
		like, fav := getCounts(ctx, svcCtx, it.Id)
		liked, faved := getMarked(ctx, svcCtx, it.Id)
		items = append(items, types.FeedItem{
			Id:             it.Id,
			Title:          it.Title,
			Description:    it.Description,
			CoverImage:     cover,
			Tags:           it.Tags,
			TagJson:        author.TagsJson,
			AuthorAvatar:   author.Avatar,
			AuthorNickname: author.Nickname,
			LikeCount:      like,
			FavoriteCount:  fav,
			Liked:          liked,
			Faved:          faved,
			IsTop:          it.IsTop,
			Visible:        it.Visible,
			PublishTime:    it.PublishTime,
		})
	}
	return &types.FeedPage{Items: items, HasMore: p.HasMore, Size: p.Size, Page: p.Page}
}

type userBrief struct {
	Nickname string
	Avatar   string
	TagsJson string
}

func getUserBrief(ctx context.Context, svcCtx *svc.ServiceContext, uid int64) userBrief {
	if uid <= 0 {
		return userBrief{}
	}
	resp, err := svcCtx.UserRpc.GetById(ctx, &userpb.GetByIdReq{Id: uid})
	if err != nil || resp == nil || resp.User == nil {
		return userBrief{}
	}
	return userBrief{
		Nickname: resp.User.Nickname,
		Avatar:   resp.User.Avatar,
		TagsJson: resp.User.TagsJson,
	}
}

func getCounts(ctx context.Context, svcCtx *svc.ServiceContext, entityId string) (int64, int64) {
	resp, err := svcCtx.CounterRpc.GetCounts(ctx, &counterpb.GetCountsReq{
		EntityType: "knowpost",
		EntityId:   entityId,
	})
	if err != nil || resp == nil || resp.Counts == nil {
		return 0, 0
	}
	return resp.Counts["like"], resp.Counts["fav"]
}

func getMarked(ctx context.Context, svcCtx *svc.ServiceContext, entityId string) (bool, bool) {
	uid, ok := ctxdata.GetUserId(ctx)
	if !ok || uid <= 0 {
		return false, false
	}
	likedResp, errLike := svcCtx.CounterRpc.IsMarked(ctx, &counterpb.IsMarkedReq{
		EntityType: "knowpost",
		EntityId:   entityId,
		Metric:     "like",
		UserId:     uid,
	})
	favedResp, errFav := svcCtx.CounterRpc.IsMarked(ctx, &counterpb.IsMarkedReq{
		EntityType: "knowpost",
		EntityId:   entityId,
		Metric:     "fav",
		UserId:     uid,
	})
	liked := errLike == nil && likedResp != nil && likedResp.Marked
	faved := errFav == nil && favedResp != nil && favedResp.Marked
	return liked, faved
}

func parseInt64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}
