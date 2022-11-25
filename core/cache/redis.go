package cache

import (
	"context"
	"encoding/json"
	"errors"
	"math/rand"
	"time"

	"github.com/go-redis/redis/v9"
)

const (
	expireTime     = 86400
	expireNullTime = 5
)

type RedisCache struct {
	cli        *redis.Client
	expire     int
	expireNull int
}

func NewRedisCache(cli *redis.Client, expire ...int) *RedisCache {
	ins := &RedisCache{
		cli:        cli,
		expire:     expireTime,
		expireNull: expireNullTime,
	}
	if len(expire) > 0 {
		ins.expire = expire[0]
	}
	if len(expire) > 1 {
		ins.expireNull = expire[1]
	}
	return ins
}

func (c *RedisCache) SetExpire(expire int) {
	c.expire = expire
}

func (c *RedisCache) SetExpireNull(expireNull int) {
	c.expireNull = expireNull
}

func (c *RedisCache) GetCacheByKeyOrFunc(ctx context.Context, key string, info interface{}, fn func() error) error {
	str, err := c.cli.Get(ctx, key).Result()
	switch str {
	case "":
		err = fn()
		switch err {
		case nil:
			strByte, _ := json.Marshal(info)
			c.cli.SetEx(ctx, key, strByte, c.setRandExpire())
		default:
			c.cli.SetEx(ctx, key, "null", time.Duration(c.expireNull)*time.Second)
		}
		return err
	case "null":
		return errors.New("数据不存在")
	default:
		return json.Unmarshal([]byte(str), info)
	}
}

func (c *RedisCache) setRandExpire() time.Duration {
	n := 0
	if c.expire > 3600 {
		rand.Seed(time.Now().UnixNano())
		n = rand.Intn(c.expire / 10)
	}
	return time.Duration(c.expire+n) * time.Second
}
