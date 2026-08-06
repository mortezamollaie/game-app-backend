package redismatching

import (
	"context"
	"fmt"
	"game-app/entity"
	"game-app/pkg/richerror"
	"time"

	"github.com/redis/go-redis/v9"
)

const WaitingListPredix = "waitinglist"

func (d DB) AddToWaitingList(userID uint, category entity.Category) error {
	const op = richerror.Op("redismatching.AddToWaitingList")

	_, err := d.adapter.Client().
		ZAdd(context.Background(),
			fmt.Sprintf("%s,%s", WaitingListPredix, category),
			redis.Z{Score: float64(time.Now().UnixMicro()), Member: fmt.Sprintf("%s", userID)}).
		Result()

	if err != nil {
		return richerror.New(op).WithErr(err).WithKind(richerror.KindUnexpected)
	}

	return nil
}
