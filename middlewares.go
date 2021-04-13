package devduck

import (
	"log"
	"net/http"
	"time"
)

type MiddlewareFunc func(http.Handler) http.Handler

func LogMw(prefix string) MiddlewareFunc {
	log.Printf("call LogMw(%s)\n", prefix)

	return func(next http.Handler) http.Handler {
		log.Printf("call LogMw(%s)(next)\n", prefix)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func(start time.Time) {
				log.Println(prefix, "request time", time.Since(start))
			}(time.Now())
			log.Println(prefix, "executing log middleware")
			next.ServeHTTP(w, r)
		})
	}
}

func OtherMw(prefix string) MiddlewareFunc {
	log.Printf("call OtherMw(%s)\n", prefix)

	return func(next http.Handler) http.Handler {
		log.Printf("call OtherMw(%s)(next)\n", prefix)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println(prefix, "executing other middleware")
			next.ServeHTTP(w, r)
		})
	}
}

func HZMw(next http.Handler) http.Handler {
	log.Printf("call HZMw(next)\n")

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("executing HZ middleware")
		next.ServeHTTP(w, r)
	})
}

func WithMw(h http.Handler, mws ...MiddlewareFunc) http.Handler {
	log.Println("WithMw, loading mws...")
	for _, mw := range mws {
		h = mw(h)
	}
	log.Println("WithMw, loading mws DONE")
	return h
}
