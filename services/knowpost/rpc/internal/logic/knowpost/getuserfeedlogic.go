package knowpostlogic

import (
	"context"
	"strconv"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
)

type GetUserFeedLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserFeedLogic {
	return &GetUserFeedLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GetUserFeed 获取用户的个性化Feed流（推拉混合）
func (l *GetUserFeedLogic) GetUserFeed(in *knowpost.GetUserFeedReq) (*knowpost.FeedPage, error) {
	page, size := normalizePage(in.Page, in.Size)

	// 使用FeedReader获取推拉混合的Feed
	postIDs, hasMore, err := l.svcCtx.FeedReader.GetFeed(l.ctx, in.UserId, page, size)
	if err != nil {
		l.Logger.Errorf("FeedReader.GetFeed failed: user=%d, page=%d, size=%d, err=%v",
			in.UserId, page, size, err)
		return nil, err
	}

	// 批量获取帖子详情
	items := make([]*knowpost.FeedItem, 0, len(postIDs))
	for _, postID := range postIDs {
		// 调用GetDetail获取完整信息
		detail, err := l.svcCtx.KnowPostsModel.FindOne(l.ctx, uint64(postID))
		if err != nil {
			l.Logger.Infof("get post detail failed: id=%d, err=%v", postID, err)
			continue
		}

		var imgUrls []string
		if detail.ImgUrls.Valid {
			imgUrls = []string{detail.ImgUrls.String}
		}

		item := &knowpost.FeedItem{
			Id:          strconv.FormatUint(detail.Id, 10),
			CreatorId:   int64(detail.CreatorId),
			Title:       detail.Title.String,
			Description: detail.Description.String,
			ImgUrls:     imgUrls,
		}
		items = append(items, item)
	}

	return &knowpost.FeedPage{
		Items:   items,
		HasMore: hasMore,
		Page:    int32(page),
		Size:    int32(size),
	}, nil
}
