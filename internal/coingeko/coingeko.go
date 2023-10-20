package coingeko

import (
	"context"
	"log"
	"net/http"
	"testtask/internal/coingeko/adapters/cache"
	"testtask/internal/coingeko/adapters/exchange"
	"testtask/internal/coingeko/controllers/http/handler"
	"testtask/internal/coingeko/service"
	"time"
)

func Run(ctx context.Context, router *http.ServeMux) error {
	ex := exchange.New()
	store := cache.New()
	svc := service.New(ex, store)
	go func(context.Context) {
		for {
			err := svc.CollectData()
			if err != nil {
				log.Println(err)
				ctx.Done()
				return
			}
			time.Sleep(time.Minute * 10)
		}

	}(ctx)
	h := handler.New(svc)
	h.RegisterRoutes(router)
	return nil
}
