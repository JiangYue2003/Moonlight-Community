package feed

// 推拉混合架构的核心常量配置

const (
	// BIGV_THRESHOLD 大V阈值：粉丝数超过此值使用拉模式，否则使用推模式
	BIGV_THRESHOLD = 1000

	// INBOX_MAX_SIZE 收件箱最大容量（保留最新的N条）
	INBOX_MAX_SIZE = 1000

	// INBOX_TTL 收件箱过期时间（7天）
	INBOX_TTL_SECONDS = 7 * 24 * 3600

	// BIGV_OUTBOX_MAX_SIZE 大V发件箱最大容量
	BIGV_OUTBOX_MAX_SIZE = 100

	// BIGV_OUTBOX_TTL 大V发件箱过期时间（24小时）
	BIGV_OUTBOX_TTL_SECONDS = 24 * 3600
)

// Redis Key 模板
const (
	// FEED_INBOX_KEY 用户收件箱 feed:inbox:{user_id}
	FEED_INBOX_KEY = "feed:inbox:%d"

	// FEED_BIGV_OUTBOX_KEY 大V发件箱 feed:bigv:{creator_id}
	FEED_BIGV_OUTBOX_KEY = "feed:bigv:%d"

	// FANOUT_PROCESSING_KEY 扇出处理标记（幂等性）
	FANOUT_PROCESSING_KEY = "feed:fanout:processing:%d"

	// FANOUT_PROCESSING_TTL 扇出处理标记过期时间（7天）
	FANOUT_PROCESSING_TTL_SECONDS = 7 * 24 * 3600
)

// Kafka Topic
const (
	// FEED_FANOUT_TOPIC Feed扇出事件Topic
	FEED_FANOUT_TOPIC = "feed-fanout"
)
