package tracker

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"github.com/xgfone/bt/metainfo"
	"net"
	"strconv"
	"time"
)

var ctx = context.Background()
var Rdb *redis.Client

func RedisCount(key string) (count int64) {
	count, _ = Rdb.ZCount(ctx, key, "-inf", "+inf").Result()
	return
}

func RedisAdd(infohash string, left int64, ip net.IP, port uint16) {
	if left == 0 {
		infohash += ":complete"
	} else {
		infohash += ":incomplete"
	}
	value := redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: marshalJSON(ip, port),
	}
	Rdb.ZAdd(ctx, infohash, value)

	// expire key after 30 mins
	Rdb.Expire(ctx, infohash, time.Second*1800)
}

func RedisGet(infohash string) (result []metainfo.Address) {

	// remove all peers which are 15 minutes old
	max := strconv.Itoa(int(time.Now().Unix() - 900))
	Rdb.ZRemRangeByScore(ctx, infohash+":complete", "0", max)
	Rdb.ZRemRangeByScore(ctx, infohash+":incomplete", "0", max)

	//get 20 random seeders
	items, err := Rdb.ZRandMember(ctx, infohash+":complete", 20).Result()
	if err != nil {
		return
	}
	result = make([]metainfo.Address, len(items))
	for i, _ := range items {
		result[i] = unmarshalJSON(items[i])
	}

	// if there are less than 20 seeders
	// fetch 20 random leechers
	if len(result) < 20 {
		items, _ = Rdb.ZRandMember(ctx, infohash+":incomplete", 20).Result()
		for _, item := range items {
			result = append(result, unmarshalJSON(item))
		}
	}
	return
}

func marshalJSON(ip net.IP, port uint16) string {
	peer := metainfo.Address{
		IP:   ip,
		Port: port,
	}
	bytes, _ := json.Marshal(peer)
	return string(bytes)
}

func unmarshalJSON(str string) metainfo.Address {
	var peer metainfo.Address
	_ = json.Unmarshal([]byte(str), &peer)
	return peer
}
