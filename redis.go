package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6390",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	err := rdb.Set(ctx, "key1", "value1", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key1", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	zSetKey := "waitinglist:football"

	res, err := rdb.ZAdd(ctx, zSetKey, redis.Z{
		Score:  float64(time.Now().UnixMicro()),
		Member: "1",
	}).Result()

	if err != nil {
		fmt.Println("err ", err)
	} else {
		fmt.Println("res", res)
	}

	list, err := rdb.ZRangeWithScores(ctx, zSetKey, 0, time.Now().UnixMicro()).Result()

	if err != nil {
		fmt.Println("err ", err)
	}

	for _, item := range list {
		fmt.Println("member, score", item.Member, int64(item.Score))

		mStr, ok := item.Member.(string)

		if ok && mStr == "1" {
			res, err := rdb.ZRem(ctx, zSetKey, item.Member).Result()
			if err != nil {
				fmt.Println("err ", err)
			}
			fmt.Println("res", res)
		}
	}

}
