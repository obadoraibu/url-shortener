package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIserver_handleHello(t *testing.T) {
	s := New(NewConfig())
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	rec := httptest.NewRecorder()
	s.HandleHello().ServeHTTP(rec, req)

	assert.Equal(t, rec.Body.Bytes(), []byte("Hello"))
}
