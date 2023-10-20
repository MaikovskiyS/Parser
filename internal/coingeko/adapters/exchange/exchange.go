package exchange

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"testtask/internal/coingeko/model"
	"time"
)

var gekoUrl = "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=250&page=1"
var timeout = time.Second * 3

type exchange struct {
}

func New() *exchange {
	return &exchange{}
}
func (e *exchange) GetPrices() (model.Result, error) {
	cl := http.DefaultClient
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	r, err := http.NewRequestWithContext(ctx, "GET", gekoUrl, nil)
	if err != nil {
		log.Println(err)
		return model.Result{}, err
	}
	res, err := cl.Do(r)
	if err != nil {
		log.Println(err)
		return model.Result{}, err
	}
	defer res.Body.Close()
	var result model.Result
	err = json.NewDecoder(res.Body).Decode(&result.Coins)
	if err != nil {
		return model.Result{}, err
	}
	return result, nil
}
