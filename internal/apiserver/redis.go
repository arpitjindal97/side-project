package apiserver

import (
	"context"
	"github.com/go-redis/redis/v9"
	"time"
)

var ctx = context.Background()
var Rdb *redis.Client
var RedisKey string

func RedisCount(time string) (count int64) {
	count, _ = Rdb.ZCount(ctx, RedisKey, "0", time).Result()
	return
}

func RedisAdd(infohash string) int {
	value := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: infohash,
	}
	val, _ := Rdb.ZAddNX(ctx, RedisKey, value).Result()
	return int(val)
}

func RedisGet(time string, offset, count int64) (infohash []string) {
	member := &redis.ZRangeBy{
		Min:    "0",
		Max:    time,
		Offset: offset,
		Count:  count,
	}
	infohash, _ = Rdb.ZRangeByScore(ctx, RedisKey, member).Result()
	return
}

func RedisRemove(infohash string) {
	Rdb.ZRem(ctx, RedisKey, infohash)
}
