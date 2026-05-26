package ossx

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

// PresignReq 描述一次预签名 PUT 请求。
type PresignReq struct {
	ObjectKey   string
	ContentType string
	// ExpiresInSec 0 时回退 cfg.PresignExpiresSec。
	ExpiresInSec int64
}

// PresignResp 返回签名 URL 与前端 PUT 时必须复用的 header。
type PresignResp struct {
	Url        string
	ObjectKey  string
	ExpiresIn  int64
	Headers    map[string]string
	ContentUrl string
}

// Presign 生成 PUT 预签名 URL。
// 前端 PUT 时 header Content-Type 必须与 req.ContentType 完全一致，否则 OSS 鉴权失败。
func (c *Client) Presign(ctx context.Context, req PresignReq) (*PresignResp, error) {
	if req.ObjectKey == "" {
		return nil, errors.New("ossx: ObjectKey required")
	}
	if req.ContentType == "" {
		return nil, errors.New("ossx: ContentType required")
	}
	exp := req.ExpiresInSec
	if exp <= 0 {
		exp = c.cfg.PresignExpiresSec
	}
	if exp <= 0 {
		exp = 600
	}
	url, err := c.bucket.SignURL(req.ObjectKey, http.MethodPut, exp,
		oss.ContentType(req.ContentType))
	if err != nil {
		return nil, err
	}
	return &PresignResp{
		Url:        url,
		ObjectKey:  req.ObjectKey,
		ExpiresIn:  exp,
		Headers:    map[string]string{"Content-Type": req.ContentType},
		ContentUrl: c.BuildContentUrl(req.ObjectKey),
	}, nil
}

// PutObject 后端中转上传（用于头像等小文件）。返回 OSS 计算的 ETag（去除引号）。
func (c *Client) PutObject(ctx context.Context, key string, body io.Reader, contentType string) (string, error) {
	if err := c.bucket.PutObject(key, body, oss.ContentType(contentType)); err != nil {
		return "", err
	}
	meta, err := c.bucket.GetObjectDetailedMeta(key)
	if err != nil {
		return "", err
	}
	etag := meta.Get("ETag")
	return strings.Trim(etag, `"`), nil
}

// BuildContentUrl 根据 PublicDomain 或 endpoint 拼接对象的公开访问 URL。
func (c *Client) BuildContentUrl(objectKey string) string {
	if c.cfg.PublicDomain != "" {
		domain := strings.TrimRight(c.cfg.PublicDomain, "/")
		return domain + "/" + objectKey
	}
	endpoint := strings.TrimPrefix(strings.TrimPrefix(c.cfg.Endpoint, "https://"), "http://")
	return "https://" + c.cfg.Bucket + "." + endpoint + "/" + objectKey
}
