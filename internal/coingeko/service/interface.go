package service

import (
	"testtask/internal/coingeko/model"
	"time"
)

type Exchange interface {
	GetPrices() (model.Result, error)
}
type Cache interface {
	Set(key string, value model.Coin, duration time.Duration)
	Get(key string) (model.Coin, bool)
	Delete(key string) error
}
