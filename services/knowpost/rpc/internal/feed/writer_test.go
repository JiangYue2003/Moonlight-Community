package feed

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zeromicro/go-zero/core/logx"
)

// MockRedisClient 用于测试的 Redis Mock
type MockRedisClient struct {
	zaddCalls    []ZAddCall
	expireCalls  []ExpireCall
	existsCalls  []ExistsCall
	setExCalls   []SetExCall
}

type ZAddCall struct {
	Key     string
	Members []interface{}
}

type ExpireCall struct {
	Key     string
	Seconds int
}

type ExistsCall struct {
	Key string
}

type SetExCall struct {
	Key     string
	Value   interface{}
	Seconds int
}

func NewMockRedisClient() *MockRedisClient {
	return &MockRedisClient{
		zaddCalls:   []ZAddCall{},
		expireCalls: []ExpireCall{},
		existsCalls: []ExistsCall{},
		setExCalls:  []SetExCall{},
	}
}

func (m *MockRedisClient) ZAdd(ctx context.Context, key string, members ...interface{}) error {
	m.zaddCalls = append(m.zaddCalls, ZAddCall{Key: key, Members: members})
	return nil
}

func (m *MockRedisClient) ZRemRangeByRank(ctx context.Context, key string, start, stop int64) error {
	return nil
}

func (m *MockRedisClient) Expire(ctx context.Context, key string, seconds int) error {
	m.expireCalls = append(m.expireCalls, ExpireCall{Key: key, Seconds: seconds})
	return nil
}

func (m *MockRedisClient) Exists(ctx context.Context, key string) (bool, error) {
	m.existsCalls = append(m.existsCalls, ExistsCall{Key: key})
	return false, nil
}

func (m *MockRedisClient) SetEx(ctx context.Context, key string, value interface{}, seconds int) error {
	m.setExCalls = append(m.setExCalls, SetExCall{Key: key, Value: value, Seconds: seconds})
	return nil
}

func (m *MockRedisClient) SCard(ctx context.Context, key string) (int64, error) {
	return 0, nil
}

// MockKafkaProducer 用于测试的 Kafka Mock
type MockKafkaProducer struct {
	messages []KafkaMessage
}

type KafkaMessage struct {
	Topic string
	Key   string
	Value []byte
}

func NewMockKafkaProducer() *MockKafkaProducer {
	return &MockKafkaProducer{messages: []KafkaMessage{}}
}

func (m *MockKafkaProducer) SendMessage(ctx context.Context, topic string, key string, value []byte) error {
	m.messages = append(m.messages, KafkaMessage{
		Topic: topic,
		Key:   key,
		Value: value,
	})
	return nil
}

// ===== 测试用例 =====

func TestFeedWriter_PushMode(t *testing.T) {
	// 场景：普通用户发帖（粉丝数 <= 1000），应该使用推模式

	redis := NewMockRedisClient()
	kafka := NewMockKafkaProducer()
	writer := NewFeedWriter(redis, kafka, logx.WithContext(context.Background()))

	ctx := context.Background()
	postID := int64(123)
	creatorID := int64(456)
	followerCount := int64(500) // 普通用户

	err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)

	// 验证
	assert.NoError(t, err)

	// 1. 应该写入自己的收件箱
	assert.Equal(t, 1, len(redis.zaddCalls))
	assert.Contains(t, redis.zaddCalls[0].Key, "feed:inbox:456")

	// 2. 应该发送 Kafka 消息（推模式）
	assert.Equal(t, 1, len(kafka.messages))
	assert.Equal(t, FEED_FANOUT_TOPIC, kafka.messages[0].Topic)
}

func TestFeedWriter_PullMode(t *testing.T) {
	// 场景：大V发帖（粉丝数 > 1000），应该使用拉模式

	redis := NewMockRedisClient()
	kafka := NewMockKafkaProducer()
	writer := NewFeedWriter(redis, kafka, logx.WithContext(context.Background()))

	ctx := context.Background()
	postID := int64(123)
	creatorID := int64(456)
	followerCount := int64(5000) // 大V

	err := writer.OnPostPublished(ctx, postID, creatorID, followerCount)

	// 验证
	assert.NoError(t, err)

	// 1. 应该写入自己的收件箱
	assert.GreaterOrEqual(t, len(redis.zaddCalls), 1)

	// 2. 应该写入大V发件箱
	foundOutbox := false
	for _, call := range redis.zaddCalls {
		if call.Key == "feed:bigv:456" {
			foundOutbox = true
			break
		}
	}
	assert.True(t, foundOutbox, "应该写入大V发件箱")

	// 3. 不应该发送 Kafka 消息（拉模式）
	assert.Equal(t, 0, len(kafka.messages))
}

func TestFeedWriter_Threshold(t *testing.T) {
	// 测试阈值边界

	redis := NewMockRedisClient()
	kafka := NewMockKafkaProducer()
	writer := NewFeedWriter(redis, kafka, logx.WithContext(context.Background()))

	ctx := context.Background()

	// 恰好等于阈值（1000），应该使用推模式
	err := writer.OnPostPublished(ctx, 1, 1, 1000)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(kafka.messages), "粉丝数=1000应该使用推模式")

	// 重置
	kafka.messages = []KafkaMessage{}

	// 超过阈值（1001），应该使用拉模式
	err = writer.OnPostPublished(ctx, 2, 2, 1001)
	assert.NoError(t, err)
	assert.Equal(t, 0, len(kafka.messages), "粉丝数=1001应该使用拉模式")
}
