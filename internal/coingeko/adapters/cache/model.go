package cache

import (
	"testtask/internal/coingeko/model"
	"time"
)

type Item struct {
	Value      model.Coin
	Created    time.Time
	Expiration int64
}
