package statserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatServer_GetOrderBook(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "get-order-book", nil)
	s.GetOrderBook().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "get_order_book")
}

func TestStatServer_SaveOrderBook(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "save-order-book", nil)
	s.SaveOrderBook().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "save_order_book")
}

func TestStatServer_GetOrderHistory(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "get-order-history", nil)
	s.GetOrderHistory().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "get_order_history")
}

func TestStatServer_SaveOrder(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "save-order", nil)
	s.SaveOrder().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "save_order")
}
