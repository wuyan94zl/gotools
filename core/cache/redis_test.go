package cache

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/magiconair/properties/assert"
	"github.com/redis/go-redis/v9"
)

var redisConn = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "123456"})

func TestGetCacheByKeyOrFunc(t *testing.T) {
	cli := NewRedisCache(redisConn, 2, 1)
	str := ""
	err := cli.GetCacheByKeyOrFunc(context.Background(), "testCacheFunc", &str, func() error {
		str = "wuyan"
		return nil
	})
	assert.Equal(t, "wuyan", str)
	err = cli.GetCacheByKeyOrFunc(context.Background(), "testCacheFunc", &str, func() error {
		str = "wuyan94zl"
		return nil
	})
	assert.Equal(t, "wuyan", str)
	time.Sleep(time.Second * 3)
	err = cli.GetCacheByKeyOrFunc(context.Background(), "testCacheFunc", &str, func() error {
		str = "wuyan94zl"
		return nil
	})
	assert.Equal(t, "wuyan94zl", str)

	data := make(map[string]string)
	err = cli.GetCacheByKeyOrFunc(context.Background(), "testCacheFuncMap", &data, func() error {
		data["name"] = "wuyan94zl"
		return nil
	})
	oneMap := make(map[string]string)
	err = cli.GetCacheByKeyOrFunc(context.Background(), "testCacheFuncMap", &oneMap, func() error {
		data["name"] = "wuyan"
		return nil
	})
	one, _ := json.Marshal(oneMap)
	twoMap := make(map[string]string)
	twoMap["name"] = "wuyan94zl"
	two, _ := json.Marshal(twoMap)
	assert.Equal(t, string(one), string(two))
	assert.Equal(t, nil, err)
}
