package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type readerIdentity struct {
	UserID      int64  `json:"user_id"`
	AccessToken string `json:"access_token,omitempty"`
}

type feedItemResponse struct {
	ID        string `json:"id"`
	CreatorID int64  `json:"creatorId"`
	Title     string `json:"title"`
}

type feedPageResponse struct {
	Items      []feedItemResponse `json:"items"`
	HasMore    bool               `json:"hasMore"`
	Page       int32              `json:"page"`
	Size       int32              `json:"size"`
	NextCursor string             `json:"nextCursor"`
}

type gatewayFeedClient struct {
	baseURL string
	client  *http.Client
}

func newGatewayFeedClient(baseURL string, client *http.Client) *gatewayFeedClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &gatewayFeedClient{baseURL: strings.TrimRight(baseURL, "/"), client: client}
}

func (c *gatewayFeedClient) GetUserFeed(
	ctx context.Context,
	reader readerIdentity,
	read feedReadRequest,
) (*feedPageResponse, error) {
	if reader.AccessToken == "" {
		return nil, fmt.Errorf("gateway access token is required for user %d", reader.UserID)
	}

	endpoint, err := url.Parse(c.baseURL + "/api/v1/knowposts/following-feed")
	if err != nil {
		return nil, fmt.Errorf("parse gateway URL: %w", err)
	}
	query := endpoint.Query()
	if read.Cursor == "" {
		query.Set("page", strconv.FormatInt(int64(read.Page), 10))
	} else {
		query.Set("cursor", read.Cursor)
	}
	query.Set("size", strconv.FormatInt(int64(read.Size), 10))
	endpoint.RawQuery = query.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, endpoint.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("create gateway request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+reader.AccessToken)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("gateway following feed: %w", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return nil, fmt.Errorf("gateway following feed: status=%d body=%s", resp.StatusCode, strings.TrimSpace(string(body)))
	}

	var pageResp feedPageResponse
	if err := json.NewDecoder(resp.Body).Decode(&pageResp); err != nil {
		return nil, fmt.Errorf("decode gateway following feed: %w", err)
	}
	return &pageResp, nil
}
