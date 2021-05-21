package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/dev-duck/handler"
	"github.com/oklog/run"
)

func Router() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", handler.HealthHandler)

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprint(rw, "Devduck Auth OK")
	})

	return r
}

func main() {
	fmt.Println("Welcome to devduckauth")

	address := ":8090"

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
