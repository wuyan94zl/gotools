package cache

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisCache struct {
	Cli        *redis.Client
	Expire     time.Duration
	ExpireNull time.Duration
}

func NewRedisCacheTools(cli *redis.Client, expire ...time.Duration) *RedisCache {
	ins := &RedisCache{
		Cli:        cli,
		Expire:     time.Hour,
		ExpireNull: time.Second * 5,
	}
	if len(expire) > 0 {
		ins.Expire = expire[0]
	}
	if len(expire) > 1 {
		ins.ExpireNull = expire[1]
	}
	return ins
}

func (c *RedisCache) GetCacheByKeyOrFunc(ctx context.Context, key string, info interface{}, fn func() error) error {
	str, err := c.Cli.Get(ctx, key).Result()
	switch str {
	case "":
		err = fn()
		switch err {
		case nil:
			strByte, _ := json.Marshal(info)
			c.Cli.SetEx(ctx, key, strByte, c.Expire)
		default:
			c.Cli.SetEx(ctx, key, "null", c.ExpireNull)
		}
		return err
	case "null":
		return errors.New("数据不存在")
	default:
		return json.Unmarshal([]byte(str), info)
	}
}
