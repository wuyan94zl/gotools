package model

import (
	"context"
	"encoding/json"
	"math/rand"
	"sync"
	"time"

	"github.com/go-redis/redis/v9"
	"gorm.io/gorm"
)

var (
	NotFoundExpire = time.Second * 5
	CacheExpire    = 86400
)

type BashModel struct {
	Conn  *gorm.DB
	Cache *redis.Client
	mu    sync.Mutex
}

func (b *BashModel) CacheFirst(ctx context.Context, info interface{}, fn func() error, key string) error {
	result, _ := b.getCache(ctx, key)
	switch result {
	case "":
		err := fn()
		switch err {
		case gorm.ErrRecordNotFound:
			b.setCache(ctx, key, "null", NotFoundExpire)
		case nil:
			strByte, _ := json.Marshal(info)
			b.setCache(ctx, key, strByte, b.setRandExpire(CacheExpire))
		}
		return err
	case "null":
		return gorm.ErrRecordNotFound
	default:
		return json.Unmarshal([]byte(result), info)
	}
}

func (b *BashModel) CacheUpdate(ctx context.Context, fn func() error, key string) error {
	b.delCache(ctx, key)
	err := fn()
	b.delCache(ctx, key)
	return err
}

func (b *BashModel) CacheDelete(ctx context.Context, fn func() error, key string) error {
	err := fn()
	if err == nil {
		b.delCache(ctx, key)
	}
	return err
}

func (b *BashModel) getCache(ctx context.Context, key string) (string, error) {
	return b.Cache.Get(ctx, key).Result()
}

func (b *BashModel) setCache(ctx context.Context, key string, value interface{}, expire time.Duration) {
	b.Cache.SetEx(ctx, key, value, expire)
}

func (b *BashModel) delCache(ctx context.Context, key string) {
	b.Cache.Del(ctx, key)
}

func (b *BashModel) setRandExpire(expire int) time.Duration {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(3600)
	return time.Duration(expire+n) * time.Second
}
