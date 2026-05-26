package logic

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/query"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/svc"
	"github.com/zhiguang/zhiguang-go/services/search/api/internal/types"
	counterpb "github.com/zhiguang/zhiguang-go/services/counter/rpc/counter"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Search 调 ES 全文检索，拼装 highlight snippet 与游标。
func (l *SearchLogic) Search(req *types.SearchReq) (*types.SearchResp, error) {
	q := strings.TrimSpace(req.Q)
	if q == "" {
		return &types.SearchResp{Items: []types.Hit{}, HasMore: false}, nil
	}
	size := req.Size
	if size <= 0 || size > 50 {
		size = 20
	}

	after, err := query.DecodeCursor(req.After)
	if err != nil {
		// cursor 损坏 → 退化为首页
		after = nil
	}
	body := query.BuildSearchBody(q, query.ParseTags(req.Tags), size, after)

	res, err := l.svcCtx.Es.Search(l.ctx, l.svcCtx.Config.ContentIndex, body)
	if err != nil {
		return nil, err
	}

	items := make([]types.Hit, 0, len(res.Hits.Hits))
	viewerID, _ := ctxdata.GetUserId(l.ctx)
	for _, h := range res.Hits.Hits {
		var src docSource
		if err := json.Unmarshal(h.Source, &src); err != nil {
			continue
		}
		likeCount := src.LikeCount
		favCount := src.FavoriteCount
		if l.svcCtx.CounterRpc != nil && src.ContentId != "" {
			if counts, err := l.svcCtx.CounterRpc.GetCounts(l.ctx, &counterpb.GetCountsReq{
				EntityType: "knowpost",
				EntityId:   src.ContentId,
			}); err == nil && counts != nil {
				likeCount = int32(counts.Counts["like"])
				favCount = int32(counts.Counts["fav"])
			}
		}
		liked := false
		faved := false
		if viewerID > 0 && l.svcCtx.CounterRpc != nil && src.ContentId != "" {
			if resp, err := l.svcCtx.CounterRpc.IsMarked(l.ctx, &counterpb.IsMarkedReq{
				EntityType: "knowpost",
				EntityId:   src.ContentId,
				Metric:     "like",
				UserId:     viewerID,
			}); err == nil && resp != nil {
				liked = resp.Marked
			}
			if resp, err := l.svcCtx.CounterRpc.IsMarked(l.ctx, &counterpb.IsMarkedReq{
				EntityType: "knowpost",
				EntityId:   src.ContentId,
				Metric:     "fav",
				UserId:     viewerID,
			}); err == nil && resp != nil {
				faved = resp.Marked
			}
		}
		items = append(items, types.Hit{
			Id:             src.ContentId,
			Title:          src.Title,
			Description:    snippet(h.Highlight, src.Description),
			CoverImage:     firstOrEmpty(src.ImgUrls),
			Tags:           src.Tags,
			TagJson:        "",
			AuthorId:       src.AuthorId,
			AuthorNickname: src.AuthorNickname,
			AuthorAvatar:   src.AuthorAvatar,
			LikeCount:      likeCount,
			FavoriteCount:  favCount,
			Liked:          liked,
			Faved:          faved,
			IsTop:          src.IsTop,
		})
	}

	resp := &types.SearchResp{Items: items}
	if len(items) == size && len(res.Hits.Hits) > 0 {
		last := res.Hits.Hits[len(res.Hits.Hits)-1]
		resp.NextAfter = query.EncodeCursor(last.Sort)
		resp.HasMore = true
	}
	return resp, nil
}

func firstOrEmpty(items []string) string {
	if len(items) == 0 {
		return ""
	}
	return items[0]
}

// snippet 优先用 ES 高亮，缺失则取 description 前 160 字。
func snippet(hl map[string][]string, desc string) string {
	var parts []string
	if v := hl["title"]; len(v) > 0 {
		parts = append(parts, v[0])
	}
	if v := hl["body"]; len(v) > 0 {
		parts = append(parts, v[0])
	}
	if len(parts) > 0 {
		return strings.Join(parts, " … ")
	}
	r := []rune(desc)
	if len(r) > 160 {
		return string(r[:160])
	}
	return desc
}

// docSource 命中文档 _source 子集。
type docSource struct {
	ContentId      string   `json:"content_id"`
	ContentType    string   `json:"content_type"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	Tags           []string `json:"tags"`
	AuthorId       int64    `json:"author_id"`
	AuthorNickname string   `json:"author_nickname"`
	AuthorAvatar   string   `json:"author_avatar"`
	PublishTime    int64    `json:"publish_time"`
	LikeCount      int32    `json:"like_count"`
	FavoriteCount  int32    `json:"favorite_count"`
	ViewCount      int32    `json:"view_count"`
	ImgUrls        []string `json:"img_urls"`
	IsTop          bool     `json:"is_top"`
}
