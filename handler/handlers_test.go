package handler_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/antklim/dev-duck/app"
	"github.com/antklim/dev-duck/handler"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	srv := app.NewAdd(10)
	addHandler := handler.NewAddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=3", nil)
	rr := httptest.NewRecorder()

	addHandler.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "13", string(resBody))
}

func BenchmarkAdd10(b *testing.B) {
	srv := app.NewAdd(10)
	addHandler := handler.NewAddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=3", nil)
	rr := httptest.NewRecorder()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		addHandler.ServeHTTP(rr, req)
	}
}

func TestAddHandler(t *testing.T) {
	srv := app.NewAdd(14)
	addHandler := handler.AddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=4", nil)
	rr := httptest.NewRecorder()

	addHandler.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "18", string(resBody))
}

func BenchmarkAddHandler10(b *testing.B) {
	srv := app.NewAdd(10)
	addHandler := handler.AddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=4", nil)
	rr := httptest.NewRecorder()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		addHandler.ServeHTTP(rr, req)
	}
}
