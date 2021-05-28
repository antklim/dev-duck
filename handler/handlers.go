package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/antklim/dev-duck/app/iface"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}

type addHandler struct {
	srv    iface.Service
	logger log.Logger
}

func NewAddHandler(srv iface.Service, logger log.Logger) http.Handler {
	return &addHandler{
		srv:    srv,
		logger: logger,
	}
}

func (h *addHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	level.Debug(h.logger).Log("operand", s)

	operand, err := strconv.Atoi(s)
	if err != nil {
		level.Info(h.logger).Log("error", errors.Wrap(err, "invalid operand value"))
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		return
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func AddHandler(srv iface.Service, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		level.Debug(logger).Log("operand", s)

		operand, err := strconv.Atoi(s)
		if err != nil {
			level.Info(logger).Log("error", errors.Wrap(err, "invalid operand value"))
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
			return
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}

type mulHandler struct {
	srv    iface.Service
	logger log.Logger
}

func NewMulHandler(srv iface.Service, logger log.Logger) http.Handler {
	return &mulHandler{
		srv:    srv,
		logger: logger,
	}
}

func (h *mulHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("operand")
	level.Debug(h.logger).Log("operand", s)

	operand, err := strconv.Atoi(s)
	if err != nil {
		level.Info(h.logger).Log("error", errors.Wrap(err, "invalid operand value"))
		http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
		return
	}

	result := h.srv.Do(operand)
	fmt.Fprint(w, result)
}

func MulHandler(srv iface.Service, logger log.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s := r.URL.Query().Get("operand")
		level.Debug(logger).Log("operand", s)

		operand, err := strconv.Atoi(s)
		if err != nil {
			level.Info(logger).Log("error", errors.Wrap(err, "invalid operand value"))
			http.Error(w, "invalid operand value "+s, http.StatusBadRequest)
			return
		}

		result := srv.Do(operand)
		fmt.Fprint(w, result)
	}
}
