package devduck

import (
	"log"
	"net/http"
)

func Router() http.Handler {
	r := http.NewServeMux()

	addHandler := NewAddHandler(NewAdd(10))
	r.Handle("/add10", addHandler)
	r.Handle("/add5", AddHandler(NewAdd(5)))

	return r
}

func Serve(router http.Handler) {
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	log.Fatal(s.ListenAndServe())
}
