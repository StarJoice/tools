//@Date 2024/12/6 22:30
//@Desc

package redis

import (
	"context"
	"github.com/StarJoice/tools/common"
	"github.com/StarJoice/tools/ginx/session"
	"github.com/redis/go-redis/v9"
	"time"
)

var _ session.Session = &Session{}

type Session struct {
	client     redis.Cmdable
	key        string
	claims     session.Claims
	expiration time.Duration
}

func newRedisSession(
	ssid string,
	expiration time.Duration,
	client redis.Cmdable, cl session.Claims) *Session {
	return &Session{
		client:     client,
		key:        "session:" + ssid,
		expiration: expiration,
		claims:     cl,
	}
}

func (sess *Session) Destroy(ctx context.Context) error {
	return sess.client.Del(ctx, sess.key).Err()
}

func (sess *Session) Claims() session.Claims {
	return sess.claims
}

func (sess *Session) init(ctx context.Context, kvs map[string]any) error {
	pip := sess.client.Pipeline()
	for k, v := range kvs {
		pip.HSet(ctx, k, v)
	}
	pip.Expire(ctx, sess.key, sess.expiration)
	_, err := pip.Exec(ctx)
	return err
}
func (sess *Session) Set(ctx context.Context, key string, val any) error {
	return sess.client.HSet(ctx, sess.key, key, val).Err()
}
func (sess *Session) Get(ctx context.Context, key string) common.AnyValue {
	res, err := sess.client.HGet(ctx, sess.key, key).Result()
	if err != nil {
		return common.AnyValue{Err: err}
	}
	return common.AnyValue{
		Val: res,
	}
}
func (sess *Session) Del(ctx context.Context, key string) error {
	return sess.client.HDel(ctx, sess.key, key).Err()
}
