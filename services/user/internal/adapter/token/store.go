// Package token 管理 refresh token 白名单（Redis），并暴露 jwt 签发能力。
package token

import (
	"context"
	"fmt"
	"time"

	goredis "github.com/redis/go-redis/v9"
)

// Store refresh token 白名单存取。Key 模板：auth:rt:{userId}:{jti}
type Store struct {
	rdb goredis.UniversalClient
}

func NewStore(rdb goredis.UniversalClient) *Store { return &Store{rdb: rdb} }

func keyOf(userId int64, jti string) string {
	return fmt.Sprintf("auth:rt:%d:%s", userId, jti)
}

// Save 把 jti 写入白名单，TTL 与 refresh 过期时长一致。
func (s *Store) Save(ctx context.Context, userId int64, jti string, ttl time.Duration) error {
	return s.rdb.Set(ctx, keyOf(userId, jti), "1", ttl).Err()
}

// Valid 检查 jti 是否存在白名单。
func (s *Store) Valid(ctx context.Context, userId int64, jti string) (bool, error) {
	n, err := s.rdb.Exists(ctx, keyOf(userId, jti)).Result()
	if err != nil {
		return false, err
	}
	return n > 0, nil
}

// Revoke 撤销单个 jti。
func (s *Store) Revoke(ctx context.Context, userId int64, jti string) error {
	return s.rdb.Del(ctx, keyOf(userId, jti)).Err()
}

// RevokeAll 撤销用户所有 refresh token；用 SCAN+DEL 避免 KEYS 阻塞。
func (s *Store) RevokeAll(ctx context.Context, userId int64) error {
	pattern := fmt.Sprintf("auth:rt:%d:*", userId)
	iter := s.rdb.Scan(ctx, 0, pattern, 200).Iterator()
	batch := make([]string, 0, 200)
	flush := func() error {
		if len(batch) == 0 {
			return nil
		}
		err := s.rdb.Del(ctx, batch...).Err()
		batch = batch[:0]
		return err
	}
	for iter.Next(ctx) {
		batch = append(batch, iter.Val())
		if len(batch) >= 200 {
			if err := flush(); err != nil {
				return err
			}
		}
	}
	if err := iter.Err(); err != nil {
		return err
	}
	return flush()
}
