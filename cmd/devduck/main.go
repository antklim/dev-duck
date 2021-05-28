package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/dev-duck/app"
	"github.com/antklim/dev-duck/handler"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	grun "github.com/oklog/run"
)

// TODO: add https://github.com/spf13/viper configuration manager
// TODO: hook metrics to statsd

const defaultPort = "8080"

var logger log.Logger

func Router(logger log.Logger) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", handler.HealthHandler)

	addHandler := handler.NewAddHandler(app.NewAdd(10), logger)

	// Middlewares executed right to left
	addHandler = handler.WithMw(addHandler, handler.HZMw, handler.OtherMw("add10"), handler.LogMw("add10"))
	r.Handle("/add10", addHandler)

	// Middlewares executed left to right
	r.Handle("/add5", handler.ChainMw(handler.AddHandler(app.NewAdd(5), logger), handler.OtherMw("add5"), handler.LogMw("add5")))

	return r
}

func run() error {
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "container", "devduckauth", "time", log.DefaultTimestampUTC)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	address := fmt.Sprintf(":%s", port)

	s := &http.Server{
		Addr:    address,
		Handler: Router(logger),
	}

	var g grun.Group
	{
		g.Add(func() error {
			level.Info(logger).Log("msg", "starting server", "server_address", address)
			return s.ListenAndServe()
		}, func(err error) {
			level.Info(logger).Log("error", err, "msg", "stopping server")
		})
	}
	{
		ctx, cancel := context.WithCancel(context.Background())
		g.Add(func() error {
			osSignals := make(chan os.Signal, 1)
			signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)
			select {
			case sig := <-osSignals:
				err := fmt.Errorf("received signal: %s", sig)
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
	return g.Run()
}

func main() {
	if err := run(); err != nil {
		logger.Log("error", err, "msg", "server terminated")
	}
}
