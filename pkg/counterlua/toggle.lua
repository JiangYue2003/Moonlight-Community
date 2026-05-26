-- TOGGLE_LUA：原子切换位图位
-- KEYS[1] = bitmap key (bm:{metric}:{etype}:{eid}:{chunk})
-- ARGV[1] = bit offset (0..32767)
-- ARGV[2] = "add" / "remove"
-- 返回 1 表示状态发生变化（产生事件），0 表示幂等无操作
local off = tonumber(ARGV[1])
local op  = ARGV[2]
local cur = redis.call('GETBIT', KEYS[1], off)
if op == 'add' then
    if cur == 1 then return 0 end
    redis.call('SETBIT', KEYS[1], off, 1)
    return 1
else
    if cur == 0 then return 0 end
    redis.call('SETBIT', KEYS[1], off, 0)
    return 1
end
