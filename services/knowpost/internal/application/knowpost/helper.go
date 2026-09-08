package knowpostlogic

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	pb "github.com/zhiguang/zhiguang-go/services/knowpost/rpc/knowpost"
	model "github.com/zhiguang/zhiguang-go/services/knowpost/shared/model"
)

// rowToDetail 将 DB 行映射为 RPC KnowPostDetail。
func rowToDetail(row *model.KnowPosts) *pb.KnowPostDetail {
	if row == nil {
		return nil
	}
	d := &pb.KnowPostDetail{
		Id:               formatId(row.Id),
		CreatorId:        int64(row.CreatorId),
		Title:            nsToStr(row.Title),
		Description:      nsToStr(row.Description),
		ContentUrl:       nsToStr(row.ContentUrl),
		ContentObjectKey: nsToStr(row.ContentObjectKey),
		ContentEtag:      nsToStr(row.ContentEtag),
		ContentSize:      niToInt(row.ContentSize),
		ContentSha256:    nsToStr(row.ContentSha256),
		Visible:          row.Visible,
		Status:           row.Status,
		Type:             row.Type,
		IsTop:            row.IsTop != 0,
		CreateTime:       row.CreateTime.UnixMilli(),
		UpdateTime:       row.UpdateTime.UnixMilli(),
	}
	if row.TagId.Valid {
		d.TagId = row.TagId.Int64
	}
	if row.PublishTime.Valid {
		d.PublishTime = row.PublishTime.Time.UnixMilli()
	}
	d.Tags = parseStringList(row.Tags.String)
	d.ImgUrls = parseStringList(row.ImgUrls.String)
	return d
}

// rowToFeedItem feed 流单条（剥离 etag/sha256/size 等无关字段）。
func rowToFeedItem(row *model.KnowPosts) *pb.FeedItem {
	if row == nil {
		return nil
	}
	it := &pb.FeedItem{
		Id:          formatId(row.Id),
		CreatorId:   int64(row.CreatorId),
		Title:       nsToStr(row.Title),
		Description: nsToStr(row.Description),
		ContentUrl:  nsToStr(row.ContentUrl),
		Visible:     row.Visible,
		IsTop:       row.IsTop != 0,
	}
	if row.PublishTime.Valid {
		it.PublishTime = row.PublishTime.Time.UnixMilli()
	}
	it.Tags = parseStringList(row.Tags.String)
	it.ImgUrls = parseStringList(row.ImgUrls.String)
	return it
}

func parseStringList(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	var out []string
	if err := json.Unmarshal([]byte(s), &out); err != nil {
		return nil
	}
	return out
}

func encodeStringList(items []string) (string, error) {
	if items == nil {
		return "", nil
	}
	b, err := json.Marshal(items)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func formatId(id uint64) string { return strconv.FormatUint(id, 10) }

func parseId(s string) (uint64, error) {
	if s == "" {
		return 0, fmt.Errorf("empty id")
	}
	return strconv.ParseUint(s, 10, 64)
}

func nsToStr(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}

func niToInt(v sql.NullInt64) int64 {
	if v.Valid {
		return v.Int64
	}
	return 0
}

func setNS(s string) sql.NullString {
	if s == "" {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}
