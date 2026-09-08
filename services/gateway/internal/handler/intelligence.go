package handler

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zhiguang/zhiguang-go/common/ctxdata"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
	gh "github.com/zhiguang/zhiguang-go/services/gateway/internal/httpx"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/middleware"
	"github.com/zhiguang/zhiguang-go/services/gateway/internal/srv"
	llmclient "github.com/zhiguang/zhiguang-go/services/llm/rpc/client/llm"
)

func registerIntelligenceRoutes(r *gin.Engine, sc *srv.ServiceContext) {
	llmPrivate := r.Group("/api/v1/llm", middleware.RequiredAuth(sc.AuthRpc))
	llmPrivate.POST("/describe", llmDescribe(sc))

	llmCompat := r.Group("/api/v1/llm", middleware.OptionalAuth(sc.AuthRpc))
	llmCompat.GET("/qa/stream", llmQaStream(sc))
}

func llmDescribe(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		uid, _ := ctxdata.GetUserId(c.Request.Context())
		var req struct {
			Body    string `json:"body"`
			Content string `json:"content"`
		}
		if err := c.ShouldBindJSON(&req); err != nil {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, err.Error()))
			return
		}
		resp, err := sc.LlmRpc.Describe(c.Request.Context(), &llmclient.DescribeReq{UserId: uid, Body: req.Body, Content: req.Content})
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		c.JSON(http.StatusOK, gin.H{"description": resp.Description})
	}
}

func suggestDescription(sc *srv.ServiceContext) gin.HandlerFunc { return llmDescribe(sc) }

func llmQaStream(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := int64(queryInt(c, "postId", 0))
		if postID <= 0 {
			gh.WriteError(c, errorx.New(errorx.CodeBadRequest, "postId required"))
			return
		}
		streamQa(sc, c, postID)
	}
}

func qaCompatStream(sc *srv.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		postID, err := parsePathInt64(c, "id")
		if err != nil {
			gh.WriteError(c, err)
			return
		}
		streamQa(sc, c, postID)
	}
}

func streamQa(sc *srv.ServiceContext, c *gin.Context, postID int64) {
	userID, _ := ctxdata.GetUserId(c.Request.Context())
	topK := int32(queryInt(c, "topK", 5))
	maxTokens := int32(queryInt(c, "maxTokens", 1024))
	stream, err := sc.LlmRpc.QaStream(c.Request.Context(), &llmclient.QaStreamReq{
		UserId: userID, PostId: postID, Question: c.Query("question"), TopK: topK, MaxTokens: maxTokens,
	})
	if err != nil {
		gh.WriteError(c, err)
		return
	}
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")
	c.Status(http.StatusOK)
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		gh.WriteError(c, errorx.New(errorx.CodeInternalError, "streaming unsupported"))
		return
	}
	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			return
		}
		if _, err := fmt.Fprintf(c.Writer, "data: %s\n\n", chunk.Data); err != nil {
			return
		}
		flusher.Flush()
		if chunk.Done {
			return
		}
	}
}
