-- DECR_FIELD_LUA：从 Hash 聚合桶扣减一个字段
-- KEYS[1] = hash key (agg:v1:{etype}:{eid})
-- ARGV[1] = field
-- ARGV[2] = delta（要扣减的 absolute 值；脚本内做 HINCRBY -delta）
-- 返回扣减后的剩余值；若字段被消除，HDEL
local field = ARGV[1]
local delta = tonumber(ARGV[2])
local left  = redis.call('HINCRBY', KEYS[1], field, -delta)
if left == 0 then
    redis.call('HDEL', KEYS[1], field)
end
return left
