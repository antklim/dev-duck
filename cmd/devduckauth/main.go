package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/dev-duck/handler"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	grun "github.com/oklog/run"
	"github.com/pkg/errors"
)

// TODO: add https://github.com/spf13/viper configuration manager
// TODO: hook metrics to statsd

const (
	defaultPort        = "8080"
	defaultProxyTarget = "http://devduck:8080"
)

var logger log.Logger

func reverseProxy(target *url.URL, rw http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(target)

	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(rw, r)
}

func Router(proxyTarget *url.URL) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", handler.HealthHandler)

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth != "secret word" {
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		reverseProxy(proxyTarget, rw, r)
	})

	return r
}

func run() error {
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "container", "devduckauth", "time", log.DefaultTimestampUTC)
	logger = level.Info(logger)

	proxyTargetVal := os.Getenv("DEV_DUCK_URL")
	if proxyTargetVal == "" {
		proxyTargetVal = defaultProxyTarget
	}
	logger.Log("proxy_target", proxyTargetVal)

	proxyTarget, err := url.Parse(proxyTargetVal)
	if err != nil {
		return errors.Wrap(err, "failed to parse proxy target url")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	address := fmt.Sprintf(":%s", port)

	s := &http.Server{
		Addr:    address,
		Handler: Router(proxyTarget),
	}

	var g grun.Group
	{
		g.Add(func() error {
			logger.Log("msg", "starting server", "server_address", address)
			return s.ListenAndServe()
		}, func(err error) {
			logger.Log("error", err, "msg", "stopping server")
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
