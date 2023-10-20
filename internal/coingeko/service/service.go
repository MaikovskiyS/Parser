package service

import (
	"context"
	"errors"
	"fmt"
	"testtask/internal/coingeko/model"
	"time"
)

type service struct {
	exchange Exchange
	cache    Cache
}

func New(ex Exchange, c Cache) *service {

	return &service{
		cache:    c,
		exchange: ex,
	}
}
func (s *service) CollectData() error {

	result, err := s.exchange.GetPrices()
	if err != nil {
		return err
	}
	//fmt.Println(result)
	for _, coin := range result.Coins {
		s.cache.Set(coin.Name, coin, time.Minute*10)
	}
	return nil
}
func (s *service) GetPrice(ctx context.Context, coinName string) (model.Coin, error) {
	fmt.Println("in get price")
	coin, ok := s.cache.Get(coinName)
	if !ok {
		return model.Coin{}, errors.New("coin not found")
	}
	fmt.Println(coin)
	return coin, nil
}
