package devduck_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	devduck "github.com/antklim/dev-duck"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	srv := devduck.NewAdd(10)
	addHandler := devduck.NewAddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=3", nil)
	rr := httptest.NewRecorder()

	addHandler.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "13", string(resBody))
}

func BenchmarkAdd10(b *testing.B) {
	srv := devduck.NewAdd(10)
	addHandler := devduck.NewAddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=3", nil)
	rr := httptest.NewRecorder()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		addHandler.ServeHTTP(rr, req)
	}
}

func TestAddHandler(t *testing.T) {
	srv := devduck.NewAdd(14)
	addHandler := devduck.AddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=4", nil)
	rr := httptest.NewRecorder()

	addHandler.ServeHTTP(rr, req)

	res := rr.Result()
	resBody, _ := ioutil.ReadAll(res.Body)

	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, "18", string(resBody))
}

func BenchmarkAddHandler10(b *testing.B) {
	srv := devduck.NewAdd(10)
	addHandler := devduck.AddHandler(srv)

	req := httptest.NewRequest("GET", "/add10?operand=4", nil)
	rr := httptest.NewRecorder()

	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		addHandler.ServeHTTP(rr, req)
	}
}
