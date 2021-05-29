package main

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/antklim/dev-duck/handler"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/statsd"
	grun "github.com/oklog/run"
	"github.com/pkg/errors"
)

// TODO: add https://github.com/spf13/viper configuration manager

const (
	defaultPort        = "8080"
	defaultProxyTarget = "http://devduck:8080"
	defaultStatsdHost  = "localhost"
	defaultStatsdPort  = "8125"
)

var (
	logger log.Logger
	m      *statsd.Statsd
)

func reverseProxy(target *url.URL, rw http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(target)

	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(rw, r)
}

func Router(proxyTarget *url.URL) http.Handler {
	r := http.NewServeMux()

	reqCounter := m.NewCounter("req_count", 1)
	reqRejCounter := m.NewCounter("req_rejected_count", 1)

	r.HandleFunc("/health", handler.HealthHandler)

	r.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		reqCounter.Add(1)

		auth := r.Header.Get("Authorization")
		if auth != "secret word" {
			reqRejCounter.Add(1)
			http.Error(rw, "Unauthorized", http.StatusUnauthorized)
			return
		}

		reverseProxy(proxyTarget, rw, r)
	})

	return r
}

func run() error {
	// configure logger
	logger = log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "container", "devduckauth", "time", log.DefaultTimestampUTC)
	logger = level.Info(logger)

	// configure statsd client
	statsdHost := os.Getenv("STATSD_HOST")
	if statsdHost == "" {
		statsdHost = defaultStatsdHost
	}

	statsdPort := os.Getenv("STATSD_PORT")
	if statsdPort == "" {
		statsdPort = defaultStatsdPort
	}

	statsdAddress := fmt.Sprintf("%s:%s", statsdHost, statsdPort)
	logger.Log("statsd_address", statsdAddress)

	m = statsd.New("devduckauth.", log.NewNopLogger())
	report := time.NewTicker(5 * time.Second)
	defer report.Stop()

	go m.SendLoop(context.Background(), report.C, "udp", statsdAddress)

	goroutines := m.NewGauge("goroutine_count")
	go exportGoroutines(goroutines)

	// configure proxy
	proxyTargetVal := os.Getenv("DEV_DUCK_URL")
	if proxyTargetVal == "" {
		proxyTargetVal = defaultProxyTarget
	}
	logger.Log("proxy_target", proxyTargetVal)

	proxyTarget, err := url.Parse(proxyTargetVal)
	if err != nil {
		return errors.Wrap(err, "failed to parse proxy target url")
	}

	// configure server
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

func exportGoroutines(g metrics.Gauge) {
	for range time.Tick(time.Second) {
		g.Set(float64(runtime.NumGoroutine()))
	}
}

func main() {
	if err := run(); err != nil {
		logger.Log("error", err, "msg", "server terminated")
	}
}
