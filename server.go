package devduck

import (
	"log"
	"net/http"
)

func Router() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/health", HealthHandler)

	addHandler := NewAddHandler(NewAdd(10))
	// Middlewares executed right to left
	addHandler = WithMw(addHandler, HZMw, OtherMw("add10"), LogMw("add10"))
	r.Handle("/add10", addHandler)

	r.Handle("/add5", OtherMw("add5")(LogMw("add5")(AddHandler(NewAdd(5)))))

	return r
}

func Serve(router http.Handler) {
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
