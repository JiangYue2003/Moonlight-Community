package esx

import (
	"context"
	"encoding/json"
	"net/http"
)

// Index PUT /{index}/_doc/{id}?refresh=wait_for（UPSERT 语义）。
func (c *Client) Index(ctx context.Context, index, id string, doc any) error {
	path := "/" + escapePath(index) + "/_doc/" + escapePath(id) + "?refresh=wait_for"
	resp, err := c.do(ctx, http.MethodPut, path, doc)
	if err != nil {
		return err
	}
	return c.readJSON(resp, nil)
}

// Update POST /{index}/_update/{id}?refresh=wait_for
// partial 形如 { "doc": {...} } 或 { "script": {...} }
func (c *Client) Update(ctx context.Context, index, id string, partial any) error {
	path := "/" + escapePath(index) + "/_update/" + escapePath(id) + "?refresh=wait_for"
	resp, err := c.do(ctx, http.MethodPost, path, partial)
	if err != nil {
		return err
	}
	return c.readJSON(resp, nil)
}

// Delete DELETE /{index}/_doc/{id}?refresh=wait_for；找不到返回 nil（幂等）。
func (c *Client) Delete(ctx context.Context, index, id string) error {
	path := "/" + escapePath(index) + "/_doc/" + escapePath(id) + "?refresh=wait_for"
	resp, err := c.do(ctx, http.MethodDelete, path, nil)
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

// Get GET /{index}/_doc/{id}；找不到返回 (nil, nil)。
func (c *Client) Get(ctx context.Context, index, id string) (json.RawMessage, error) {
	path := "/" + escapePath(index) + "/_doc/" + escapePath(id)
	resp, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}
	var out struct {
		Found  bool            `json:"found"`
		Source json.RawMessage `json:"_source"`
	}
	if err := c.readJSON(resp, &out); err != nil {
		if IsNotFound(err) {
			return nil, nil
		}
		return nil, err
	}
	if !out.Found {
		return nil, nil
	}
	return out.Source, nil
}

// DeleteByQuery POST /{index}/_delete_by_query?refresh=true
func (c *Client) DeleteByQuery(ctx context.Context, index string, query json.RawMessage) (int64, error) {
	path := "/" + escapePath(index) + "/_delete_by_query?refresh=true&conflicts=proceed"
	resp, err := c.do(ctx, http.MethodPost, path, query)
	if err != nil {
		return 0, err
	}
	var out struct {
		Deleted int64 `json:"deleted"`
	}
	if err := c.readJSON(resp, &out); err != nil {
		return 0, err
	}
	return out.Deleted, nil
}

// Bulk POST /_bulk?refresh=wait_for；body 必须是 NDJSON（每行 \n 结尾）。
func (c *Client) Bulk(ctx context.Context, ndjson []byte) error {
	resp, err := c.do(ctx, http.MethodPost, "/_bulk?refresh=wait_for", ndjson)
	if err != nil {
		return err
	}
	var out struct {
		Errors bool            `json:"errors"`
		Items  json.RawMessage `json:"items"`
	}
	if err := c.readJSON(resp, &out); err != nil {
		return err
	}
	if out.Errors {
		return &Error{Status: http.StatusOK, Body: "bulk has item errors: " + string(out.Items)}
	}
	return nil
}
