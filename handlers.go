package devduck

import (
	"fmt"
	"net/http"
	"strconv"
)

type addHandler struct {
	srv Service
}

func NewAddHandler(srv Service) http.Handler {
	return &addHandler{srv: srv}
}

func (h *addHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	operand, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func AddHandler(srv Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		operand, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}

type mulHandler struct {
	srv Service
}

func NewMulHandler(srv Service) http.Handler {
	return &mulHandler{srv: srv}
}

func (h *mulHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	operand, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func MulHandler(srv Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		operand, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}
