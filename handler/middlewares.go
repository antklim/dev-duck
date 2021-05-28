package handler

import (
	"log"
	"net/http"
	"time"
)

// TODO: replace logger with router logger

type MiddlewareFunc func(http.Handler) http.Handler

func LogMw(prefix string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func(start time.Time) {
				log.Println(prefix, "request time", time.Since(start))
			}(time.Now())
			next.ServeHTTP(w, r)
		})
	}
}

func OtherMw(prefix string) MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(prefix, "executing OtherMw middleware")
			next.ServeHTTP(w, r)
		})
	}
}

func HZMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("executing HZ middleware")
		next.ServeHTTP(w, r)
	})
}

func WithMw(h http.Handler, mws ...MiddlewareFunc) http.Handler {
	for _, mw := range mws {
		h = mw(h)
	}
	return h
}

func ChainMw(h http.Handler, mws ...MiddlewareFunc) http.Handler {
	for i := len(mws) - 1; i >= 0; i-- {
		h = mws[i](h)
	}
	return h
}
