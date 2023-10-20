package handler

import (
	"context"
	"testtask/internal/coingeko/model"
)

type Service interface {
	GetPrice(ctx context.Context, coinName string) (model.Coin, error)
}
