// Package verification 实现验证码的发送限频、存储、校验。
//
// Redis 数据结构（与 Java 端兼容）：
//
//	HASH  auth:code:{scene}:{identifier}    {code, attempts, maxAttempts}
//	STR   auth:code:cooldown:{identifier}   "1"            （TTL=ResendInterval）
//	STR   auth:code:daily:{date}:{identifier} 计数         （TTL=至当日结束）
package verification

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	goredis "github.com/redis/go-redis/v9"
	"github.com/zhiguang/zhiguang-go/pkg/errorx"
)

const (
	SceneRegister      = "REGISTER"
	SceneLogin         = "LOGIN"
	SceneResetPassword = "RESET_PASSWORD"
)

// Config 验证码参数。
type Config struct {
	CodeTtl          time.Duration
	ResendInterval   time.Duration
	DailyLimit       int
	MaxAttempts      int
	CodeLength       int
	TooManyExtendTtl time.Duration
}

// Service 提供验证码发送/校验/失效。
type Service struct {
	rdb goredis.UniversalClient
	cfg Config
}

// New 构造 Service。
func New(rdb goredis.UniversalClient, cfg Config) *Service {
	if cfg.CodeLength <= 0 {
		cfg.CodeLength = 6
	}
	if cfg.CodeTtl <= 0 {
		cfg.CodeTtl = 5 * time.Minute
	}
	if cfg.ResendInterval <= 0 {
		cfg.ResendInterval = 60 * time.Second
	}
	if cfg.DailyLimit <= 0 {
		cfg.DailyLimit = 10
	}
	if cfg.MaxAttempts <= 0 {
		cfg.MaxAttempts = 5
	}
	if cfg.TooManyExtendTtl <= 0 {
		cfg.TooManyExtendTtl = 30 * time.Minute
	}
	return &Service{rdb: rdb, cfg: cfg}
}

// Send 发送验证码：检查冷却、日上限，写 Hash，返回明文码（demo 阶段直接返回；
// 生产版应改为通过 SMS/邮件渠道发送，调用方仅得 cooldown/expire）。
func (s *Service) Send(ctx context.Context, scene, identifier string) (string, error) {
	cdKey := fmt.Sprintf("auth:code:cooldown:%s", identifier)
	if ok, _ := s.rdb.Exists(ctx, cdKey).Result(); ok > 0 {
		return "", errorx.New(errorx.CodeVerificationCooldown, "code requested too frequently")
	}
	dayKey := fmt.Sprintf("auth:code:daily:%s:%s", time.Now().Format("20060102"), identifier)
	cnt, err := s.rdb.Incr(ctx, dayKey).Result()
	if err != nil {
		return "", err
	}
	if cnt == 1 {
		_ = s.rdb.Expire(ctx, dayKey, 24*time.Hour).Err()
	}
	if int(cnt) > s.cfg.DailyLimit {
		return "", errorx.New(errorx.CodeVerificationDailyLimit, "daily limit exceeded")
	}
	code, err := genCode(s.cfg.CodeLength)
	if err != nil {
		return "", err
	}
	hk := hashKey(scene, identifier)
	if err := s.rdb.HSet(ctx, hk, map[string]any{
		"code":        code,
		"attempts":    "0",
		"maxAttempts": strconv.Itoa(s.cfg.MaxAttempts),
	}).Err(); err != nil {
		return "", err
	}
	if err := s.rdb.Expire(ctx, hk, s.cfg.CodeTtl).Err(); err != nil {
		return "", err
	}
	if err := s.rdb.Set(ctx, cdKey, "1", s.cfg.ResendInterval).Err(); err != nil {
		return "", err
	}
	return code, nil
}

// Verify 校验验证码；成功后立即删除（防重放）。
func (s *Service) Verify(ctx context.Context, scene, identifier, code string) error {
	hk := hashKey(scene, identifier)
	vals, err := s.rdb.HGetAll(ctx, hk).Result()
	if err != nil {
		return err
	}
	if len(vals) == 0 {
		return errorx.New(errorx.CodeVerificationNotFound, "code expired or not sent")
	}
	max, _ := strconv.Atoi(vals["maxAttempts"])
	if max == 0 {
		max = s.cfg.MaxAttempts
	}
	att, _ := strconv.Atoi(vals["attempts"])
	if att >= max {
		_ = s.rdb.Expire(ctx, hk, s.cfg.TooManyExtendTtl).Err()
		return errorx.New(errorx.CodeVerificationTooMany, "too many attempts")
	}
	if vals["code"] != code {
		s.rdb.HIncrBy(ctx, hk, "attempts", 1)
		return errorx.New(errorx.CodeVerificationMismatch, "code mismatch")
	}
	s.rdb.Del(ctx, hk)
	return nil
}

// Invalidate 强制失效一个验证码。
func (s *Service) Invalidate(ctx context.Context, scene, identifier string) error {
	return s.rdb.Del(ctx, hashKey(scene, identifier)).Err()
}

// CooldownSeconds 用于响应给前端的剩余冷却秒数。
func (s *Service) CooldownSeconds() int { return int(s.cfg.ResendInterval.Seconds()) }

// ExpireSeconds 用于响应给前端的过期秒数。
func (s *Service) ExpireSeconds() int { return int(s.cfg.CodeTtl.Seconds()) }

func hashKey(scene, identifier string) string {
	return fmt.Sprintf("auth:code:%s:%s", scene, identifier)
}

func genCode(n int) (string, error) {
	if n <= 0 {
		return "", errors.New("verification: invalid code length")
	}
	max := big.NewInt(10)
	out := make([]byte, n)
	for i := 0; i < n; i++ {
		v, err := rand.Int(rand.Reader, max)
		if err != nil {
			return "", err
		}
		out[i] = byte('0') + byte(v.Int64())
	}
	return string(out), nil
}
