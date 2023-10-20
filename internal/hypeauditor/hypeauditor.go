package hypeauditor

import (
	"context"
	"testtask/internal/hypeauditor/adapters/parser"
	"testtask/internal/hypeauditor/service"
)

func Run(ctx context.Context) error {
	p := parser.New()
	svc := service.New(p)
	err := svc.Parse()
	if err != nil {
		return err
	}
	return nil
}
