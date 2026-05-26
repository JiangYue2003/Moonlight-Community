package schema

import "testing"

func TestUserIdxOf_Mapping(t *testing.T) {
	cases := map[string]int{
		UserMetricFollowings:    UserIdxFollowings,
		UserMetricFollowers:     UserIdxFollowers,
		UserMetricPosts:         UserIdxPosts,
		UserMetricLikesReceived: UserIdxLikesReceived,
		"":                      -1,
		"unknown":               -1,
		"read":                  -1, // 预留位不允许写入
	}
	for metric, want := range cases {
		if got := UserIdxOf(metric); got != want {
			t.Errorf("UserIdxOf(%q) = %d, want %d", metric, got, want)
		}
	}
}

func TestUserSchemaLayout_MatchesEntity(t *testing.T) {
	// 复用 cnt 同样的 SchemaLen / FieldSize，使 incr_field.lua 可被两个 SDS 共用。
	if UserSchemaLen != SchemaLen {
		t.Fatalf("UserSchemaLen drifted: %d vs %d", UserSchemaLen, SchemaLen)
	}
	if UserFieldSize != FieldSize {
		t.Fatalf("UserFieldSize drifted: %d vs %d", UserFieldSize, FieldSize)
	}
}

func TestUserSdsKey_AlreadyTestedInKeysTest(t *testing.T) {
	// 占位：UserSdsKey 的格式断言已在 keys_test.go 完成；此处仅做存活校验。
	if got := UserSdsKey(42); got == "" {
		t.Fatal("UserSdsKey returned empty")
	}
}

func TestUserIdxConstants_StableContract(t *testing.T) {
	// 这些索引值是 Java/Go 二进制布局契约。任何修改都会与已部署 ucnt:* 数据不兼容。
	cases := []struct {
		name string
		idx  int
		want int
	}{
		{"read", UserIdxRead, 0},
		{"followings", UserIdxFollowings, 1},
		{"followers", UserIdxFollowers, 2},
		{"posts", UserIdxPosts, 3},
		{"likes_received", UserIdxLikesReceived, 4},
	}
	for _, c := range cases {
		if c.idx != c.want {
			t.Errorf("%s: got idx=%d, want %d", c.name, c.idx, c.want)
		}
	}
}
