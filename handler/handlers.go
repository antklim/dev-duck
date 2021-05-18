package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/antklim/dev-duck/app/iface"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

type addHandler struct {
	srv iface.Service
}

func NewAddHandler(srv iface.Service) http.Handler {
	return &addHandler{srv: srv}
}

func (h *addHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	operand, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		return
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func AddHandler(srv iface.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		operand, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
			return
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}

type mulHandler struct {
	srv iface.Service
}

func NewMulHandler(srv iface.Service) http.Handler {
	return &mulHandler{srv: srv}
}

func (h *mulHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	operand, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		return
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func MulHandler(srv iface.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		operand, err := strconv.Atoi(s)
		if err != nil {
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
			return
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}
