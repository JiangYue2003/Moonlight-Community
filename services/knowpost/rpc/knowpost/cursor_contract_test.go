package knowpost

import (
	"testing"

	"google.golang.org/protobuf/proto"
)

func TestCursorFieldsSurviveProtobufRoundTrip(t *testing.T) {
	request := &GetUserFeedReq{
		UserId: 42,
		Page:   9,
		Size:   20,
		Cursor: "cursor-v1",
	}
	requestRaw, err := proto.Marshal(request)
	if err != nil {
		t.Fatalf("marshal request: %v", err)
	}
	var decodedRequest GetUserFeedReq
	if err := proto.Unmarshal(requestRaw, &decodedRequest); err != nil {
		t.Fatalf("unmarshal request: %v", err)
	}
	if decodedRequest.GetCursor() != request.Cursor {
		t.Fatalf("request cursor = %q, want %q", decodedRequest.GetCursor(), request.Cursor)
	}

	response := &FeedPage{
		Items:      []*FeedItem{{Id: "99"}},
		HasMore:    true,
		Size:       20,
		Page:       0,
		NextCursor: "cursor-v2",
	}
	responseRaw, err := proto.Marshal(response)
	if err != nil {
		t.Fatalf("marshal response: %v", err)
	}
	var decodedResponse FeedPage
	if err := proto.Unmarshal(responseRaw, &decodedResponse); err != nil {
		t.Fatalf("unmarshal response: %v", err)
	}
	if decodedResponse.GetNextCursor() != response.NextCursor {
		t.Fatalf("response next cursor = %q, want %q", decodedResponse.GetNextCursor(), response.NextCursor)
	}
}
