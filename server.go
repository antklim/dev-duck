package main

import (
	"log"
	"net/http"

	"github.com/antklim/dev-duck/app"
	"github.com/antklim/dev-duck/handler"
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

func Serve(router http.Handler) {
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
