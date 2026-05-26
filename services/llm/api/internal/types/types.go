package types

// DescribeReq 描述生成请求体。
type DescribeReq struct {
	Body    string `json:"body,optional"`
	Content string `json:"content,optional"`
}

// DescribeResp 描述生成响应。
type DescribeResp struct {
	Description string `json:"description"`
}

// QaReq 流式问答请求（GET，参数走 query；EventSource 不支持 POST）。
type QaReq struct {
	PostId    int64  `form:"postId"`
	Question  string `form:"question"`
	TopK      int    `form:"topK,default=5"`
	MaxTokens int    `form:"maxTokens,default=1024"`
}
