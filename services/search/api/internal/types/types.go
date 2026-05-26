package types

// SearchReq 检索请求。
//
// q       关键词
// size    每页大小 1..50；默认 20
// tags    逗号分隔；可选
// after   Base64URL 编码的 cursor；首页空字符串
type SearchReq struct {
	Q     string `form:"q"`
	Size  int    `form:"size,default=20"`
	Tags  string `form:"tags,optional"`
	After string `form:"after,optional"`
}

// Hit 单条命中（已聚合作者 summary）。
type Hit struct {
	Id             string   `json:"id"`
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	CoverImage     string   `json:"coverImage"`
	Tags           []string `json:"tags"`
	TagJson        string   `json:"tagJson"`
	AuthorId       int64    `json:"authorId"`
	AuthorNickname string   `json:"authorNickname"`
	AuthorAvatar   string   `json:"authorAvatar"`
	LikeCount      int32    `json:"likeCount"`
	FavoriteCount  int32    `json:"favoriteCount"`
	Liked          bool     `json:"liked"`
	Faved          bool     `json:"faved"`
	IsTop          bool     `json:"isTop"`
}

// SearchResp 分页响应。
type SearchResp struct {
	Items     []Hit  `json:"items"`
	NextAfter string `json:"nextAfter"`
	HasMore   bool   `json:"hasMore"`
}

// SuggestReq Completion suggester 请求。
type SuggestReq struct {
	Prefix string `form:"prefix"`
	Size   int    `form:"size,default=10"`
}

// SuggestResp 自动补全候选列表（去重保序）。
type SuggestResp struct {
	Items []string `json:"items"`
}
