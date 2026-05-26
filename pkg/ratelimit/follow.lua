-- 令牌桶限流（与原 Java RelationServiceImpl 内联脚本等价）。
--
-- KEYS[1] = bucket key，例如 "rl:follow:{userId}"
-- ARGV[1] = capacity         （桶容量）
-- ARGV[2] = refill_per_sec   （每秒补充令牌数）
-- ARGV[3] = now_ms           （客户端时间，毫秒）
--
-- 返回值：
--   1 → 允许通过（已扣 1 token）
--   0 → 限流命中（无 token 可扣）
--
-- HSET 字段：
--   tokens          剩余令牌数
--   last_refill_ms  上次补充令牌的时间戳（毫秒）
local key       = KEYS[1]
local capacity  = tonumber(ARGV[1])
local rate      = tonumber(ARGV[2])  -- per second
local now_ms    = tonumber(ARGV[3])

local data = redis.call('HMGET', key, 'tokens', 'last_refill_ms')
local tokens = tonumber(data[1])
local last   = tonumber(data[2])

if tokens == nil or last == nil then
    -- 首次访问：满桶
    tokens = capacity
    last = now_ms
else
    -- 按时间差补充令牌
    local delta_ms = now_ms - last
    if delta_ms < 0 then delta_ms = 0 end
    local refill = math.floor(delta_ms * rate / 1000)
    if refill > 0 then
        tokens = math.min(capacity, tokens + refill)
        last = now_ms
    end
end

local allowed = 0
if tokens >= 1 then
    tokens = tokens - 1
    allowed = 1
end

redis.call('HSET', key, 'tokens', tokens, 'last_refill_ms', last)
-- TTL 设为容量完全补满所需时间 + 60 秒缓冲；rate=0 视为冷桶，给固定 60 秒
local ttl_sec
if rate <= 0 then
    ttl_sec = 60
else
    ttl_sec = math.floor(capacity / rate) + 60
end
redis.call('EXPIRE', key, ttl_sec)

return allowed
