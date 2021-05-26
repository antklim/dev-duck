package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"

	"github.com/antklim/dev-duck/handler"
	"github.com/oklog/run"
)

func reverseProxy(target *url.URL, rw http.ResponseWriter, r *http.Request) {
	proxy := httputil.NewSingleHostReverseProxy(target)

	r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))

	proxy.ServeHTTP(rw, r)
}

func Router(proxyTarget *url.URL) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", handler.HealthHandler) // TODO: reverse proxy to main app

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

func main() {
	fmt.Println("Welcome to devduckauth")

	proxyTarget, err := url.Parse("http://devduck:8080")
	if err != nil {
		fmt.Printf("failed to parse proxy target url: %+v", err)
		return
	}

	address := ":8080"
	s := &http.Server{
		Addr:    address,
		Handler: Router(proxyTarget),
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
