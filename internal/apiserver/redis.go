package apiserver

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()
var Rdb *redis.Client

func RedisCount() (count int64) {
	count, _ = Rdb.DBSize(ctx).Result()
	return
}

func RedisGet(infohash string) error {
	_, err := Rdb.Get(ctx, infohash).Result()
	return err
}

func RedisSet(infohash string) {
	_, _ = Rdb.SetNX(ctx, infohash, "anything", time.Minute*10).Result()
}

func RedisRemove(infohash string) {
	_, _ = Rdb.GetDel(ctx, infohash).Result()
}
