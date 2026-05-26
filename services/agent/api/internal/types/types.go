package types

type CreateSessionReq struct {
	Title string `json:"title,optional"`
}

type CreateSessionResp struct {
	SessionID string `json:"sessionId"`
}

type ChatStreamReq struct {
	SessionID string `form:"sessionId"`
	Question  string `form:"question"`
	TopK      int    `form:"topK,optional"`
}

type HistoryReq struct {
	SessionID string `form:"sessionId"`
	Limit     int    `form:"limit,optional"`
}

type MessageItem struct {
	Role      string `json:"role"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
}

type HistoryResp struct {
	Items []MessageItem `json:"items"`
}

type FeedbackReq struct {
	SessionID string `json:"sessionId"`
	TraceID   string `json:"traceId"`
	Score     int    `json:"score"`
	Comment   string `json:"comment,optional"`
}

type FeedbackResp struct {
	Accepted bool `json:"accepted"`
}

type MemoryPinReq struct {
	SessionID string `json:"sessionId"`
	Content   string `json:"content"`
	Tag       string `json:"tag,optional"`
}

type MemoryPinResp struct {
	Pinned bool `json:"pinned"`
}
