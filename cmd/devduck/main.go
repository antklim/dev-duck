package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/dev-duck/app"
	"github.com/antklim/dev-duck/handler"
	"github.com/oklog/run"
)

func Router() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", handler.HealthHandler)

	addHandler := handler.NewAddHandler(app.NewAdd(10))

	// Middlewares executed right to left
	addHandler = handler.WithMw(addHandler, handler.HZMw, handler.OtherMw("add10"), handler.LogMw("add10"))
	r.Handle("/add10", addHandler)

	// Middlewares executed left to right
	r.Handle("/add5", handler.ChainMw(handler.AddHandler(app.NewAdd(5)), handler.OtherMw("add5"), handler.LogMw("add5")))

	return r
}

func main() {
	fmt.Println("Welcome to devduck")

	address := ":8080"

	s := &http.Server{
		Addr:    address,
		Handler: Router(),
	}

	var g run.Group
	{
		g.Add(func() error {
			log.Printf("Starting server at: %s\n", address)
			return s.ListenAndServe()
		}, func(err error) {
			log.Printf("The server stopped: %+v\n", err)
		})
	}
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			osSignals := make(chan os.Signal, 1)
			signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-osSignals:
				err := fmt.Errorf("received signal %s", sig)
				s.Close()
				return err
			case <-ctx.Done():
				s.Close()
				return ctx.Err()
			}
		}, func(err error) {
			cancel()
		})
	}
	fmt.Printf("The group terminated: %v\n", g.Run())
}
