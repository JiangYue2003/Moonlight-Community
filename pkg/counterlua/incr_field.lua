-- INCR_FIELD_LUA：原子增减 SDS 中第 idx 个 32 位 BigEndian 字段
-- KEYS[1] = SDS key (cnt:v1:{etype}:{eid} 或 ucnt:{userId})
-- ARGV[1] = field index (0..SCHEMA_LEN-1)
-- ARGV[2] = delta (signed int)
-- ARGV[3] = SCHEMA_LEN（字段数）
-- ARGV[4] = FIELD_SIZE（每字段字节数，固定 4）
-- 返回更新后的字段值（int）
local idx       = tonumber(ARGV[1])
local delta     = tonumber(ARGV[2])
local schemaLen = tonumber(ARGV[3])
local fieldSize = tonumber(ARGV[4])
local total     = schemaLen * fieldSize

local raw = redis.call('GET', KEYS[1])
if not raw or #raw < total then
    -- 初始化为 SCHEMA_LEN*4 字节零
    raw = string.rep('\0', total)
end

local off = idx * fieldSize
local b1, b2, b3, b4 = string.byte(raw, off + 1, off + 4)
local val = ((b1 * 16777216) + (b2 * 65536) + (b3 * 256) + b4)
if val >= 2147483648 then val = val - 4294967296 end -- to signed

val = val + delta
if val < 0 then val = 0 end
if val > 4294967295 then val = 4294967295 end

local nb1 = math.floor(val / 16777216) % 256
local nb2 = math.floor(val / 65536)    % 256
local nb3 = math.floor(val / 256)      % 256
local nb4 =            val             % 256
local newBytes = string.char(nb1, nb2, nb3, nb4)

local prefix = (off > 0) and string.sub(raw, 1, off) or ''
local suffix = string.sub(raw, off + 5)
redis.call('SET', KEYS[1], prefix .. newBytes .. suffix)
return val
