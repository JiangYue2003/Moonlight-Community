package esx

import (
	"context"
	_ "embed"
	"encoding/json"
	"fmt"
	"net/http"
)

//go:embed mapping/content.json
var contentMappingJSON []byte

//go:embed mapping/rag.json
var ragMappingJSON []byte

// ContentMapping 返回知识帖搜索索引 mapping（embed JSON）。
func ContentMapping() json.RawMessage { return contentMappingJSON }

// RagMapping 返回 RAG 向量索引 mapping（embed JSON）。
func RagMapping() json.RawMessage { return ragMappingJSON }

// IndexExists HEAD /{index} 判断索引是否存在。
func (c *Client) IndexExists(ctx context.Context, index string) (bool, error) {
	resp, err := c.do(ctx, http.MethodHead, "/"+escapePath(index), nil)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		return true, nil
	}
	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}
	return false, fmt.Errorf("esx index exists: status=%d", resp.StatusCode)
}

// EnsureIndex 幂等创建索引：存在跳过，不存在则用给定 mapping 创建。
func (c *Client) EnsureIndex(ctx context.Context, index string, mapping json.RawMessage) error {
	exists, err := c.IndexExists(ctx, index)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}
	resp, err := c.do(ctx, http.MethodPut, "/"+escapePath(index), mapping)
	if err != nil {
		return err
	}
	return c.readJSON(resp, nil)
}

// DeleteIndex DELETE /{index}；索引不存在返回 nil（幂等）。
func (c *Client) DeleteIndex(ctx context.Context, index string) error {
	resp, err := c.do(ctx, http.MethodDelete, "/"+escapePath(index), nil)
	if err != nil {
		return err
	}
	if err := c.readJSON(resp, nil); err != nil {
		if IsNotFound(err) {
			return nil
		}
		return err
	}
	return nil
}

// Count 返回索引总文档数；索引不存在返回 0（不报错）。
func (c *Client) Count(ctx context.Context, index string) (int64, error) {
	resp, err := c.do(ctx, http.MethodGet, "/"+escapePath(index)+"/_count", nil)
	if err != nil {
		return 0, err
	}
	var out struct {
		Count int64 `json:"count"`
	}
	if err := c.readJSON(resp, &out); err != nil {
		if IsNotFound(err) {
			return 0, nil
		}
		return 0, err
	}
	return out.Count, nil
}
