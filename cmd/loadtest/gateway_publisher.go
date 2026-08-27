package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	knowpostpb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	"google.golang.org/grpc"
)

type gatewayPublisherClient struct {
	baseURL string
	client  *http.Client
	tokens  map[int64]string
}

func newGatewayPublisherClient(baseURL string, client *http.Client, tokens map[int64]string) *gatewayPublisherClient {
	if client == nil {
		client = http.DefaultClient
	}
	return &gatewayPublisherClient{baseURL: strings.TrimRight(baseURL, "/"), client: client, tokens: tokens}
}

func (c *gatewayPublisherClient) CreateDraft(ctx context.Context, in *knowpostpb.CreateDraftReq, _ ...grpc.CallOption) (*knowpostpb.CreateDraftResp, error) {
	var response struct {
		ID string `json:"id"`
	}
	if err := c.doJSON(ctx, http.MethodPost, "/api/v1/knowposts/drafts", in.CreatorId, struct{}{}, &response); err != nil {
		return nil, err
	}
	if response.ID == "" {
		return nil, fmt.Errorf("gateway create draft returned an empty id")
	}
	return &knowpostpb.CreateDraftResp{Id: response.ID}, nil
}

func (c *gatewayPublisherClient) PatchMetadata(ctx context.Context, in *knowpostpb.PatchMetadataReq, _ ...grpc.CallOption) (*knowpostpb.KnowPostDetail, error) {
	body := make(map[string]any)
	if in.TitleSet {
		body["title"] = in.Title
	}
	if in.DescriptionSet {
		body["description"] = in.Description
	}
	if in.TagIdSet {
		body["tagId"] = in.TagId
	}
	if in.TagsSet {
		body["tags"], body["tagsSet"] = in.Tags, true
	}
	if in.ImgUrlsSet {
		body["imgUrls"], body["imgUrlsSet"] = in.ImgUrls, true
	}
	if in.VisibleSet {
		body["visible"] = in.Visible
	}
	if in.IsTopSet {
		body["isTop"] = in.IsTop
	}
	var response struct {
		ID string `json:"id"`
	}
	path := "/api/v1/knowposts/" + strconv.FormatInt(in.Id, 10)
	if err := c.doJSON(ctx, http.MethodPatch, path, in.CreatorId, body, &response); err != nil {
		return nil, err
	}
	return &knowpostpb.KnowPostDetail{Id: response.ID}, nil
}

func (c *gatewayPublisherClient) ConfirmContent(ctx context.Context, in *knowpostpb.ConfirmContentReq, _ ...grpc.CallOption) (*knowpostpb.Empty, error) {
	body := map[string]any{
		"objectKey": in.ObjectKey,
		"etag":      in.Etag,
		"size":      in.Size,
		"sha256":    in.Sha256,
	}
	path := "/api/v1/knowposts/" + strconv.FormatInt(in.Id, 10) + "/content/confirm"
	if err := c.doJSON(ctx, http.MethodPost, path, in.CreatorId, body, nil); err != nil {
		return nil, err
	}
	return &knowpostpb.Empty{}, nil
}

func (c *gatewayPublisherClient) Publish(ctx context.Context, in *knowpostpb.PublishReq, _ ...grpc.CallOption) (*knowpostpb.KnowPostDetail, error) {
	var response struct {
		ID string `json:"id"`
	}
	path := "/api/v1/knowposts/" + strconv.FormatInt(in.Id, 10) + "/publish"
	if err := c.doJSON(ctx, http.MethodPost, path, in.CreatorId, struct{}{}, &response); err != nil {
		return nil, err
	}
	return &knowpostpb.KnowPostDetail{Id: response.ID}, nil
}

func (c *gatewayPublisherClient) doJSON(ctx context.Context, method, path string, authorID int64, body, output any) error {
	token := c.tokens[authorID]
	if token == "" {
		return fmt.Errorf("gateway access token is required for author %d", authorID)
	}
	payload, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("encode gateway request: %w", err)
	}
	req, err := http.NewRequestWithContext(ctx, method, c.baseURL+path, bytes.NewReader(payload))
	if err != nil {
		return fmt.Errorf("create gateway request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("gateway publish workflow %s %s: %w", method, path, err)
	}
	defer resp.Body.Close()
	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusMultipleChoices {
		responseBody, _ := io.ReadAll(io.LimitReader(resp.Body, 4096))
		return fmt.Errorf("gateway publish workflow %s %s: status=%d body=%s", method, path, resp.StatusCode, strings.TrimSpace(string(responseBody)))
	}
	if output == nil || resp.StatusCode == http.StatusNoContent {
		return nil
	}
	if err := json.NewDecoder(resp.Body).Decode(output); err != nil {
		return fmt.Errorf("decode gateway publish workflow %s %s: %w", method, path, err)
	}
	return nil
}
