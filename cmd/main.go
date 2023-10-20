package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"testtask/internal/coingeko"
	"testtask/internal/hypeauditor"
)

func main() {
	//cuncurrency
	wg := &sync.WaitGroup{}
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	//init router
	mux := http.NewServeMux()

	//init coingeko pricer
	err := coingeko.Run(ctx, mux)
	if err != nil {
		cancel()
		log.Fatal(err)
	}
	//init hypeauditor parser
	err = hypeauditor.Run(ctx)
	if err != nil {
		cancel()
		log.Fatal(err)
	}

	//server
	go func(context.Context, *http.ServeMux) {
		log.Println("starting app")
		err := http.ListenAndServe(":8080", mux)
		if err != nil {
			cancel()
			return
		}
	}(ctx, mux)

	//exit
	wg.Add(1)
	go func() {
		defer wg.Done()

		select {
		case <-exit:
			log.Println("closing app by system call")
			cancel()
			os.Exit(1)
		case <-ctx.Done():
			log.Println("closing app by context")
			return
		}
	}()
	wg.Wait()

}
