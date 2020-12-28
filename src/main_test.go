package main

import (
	"api-message-security/src/hmac"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"
)

var (
	server   Server
	__secret = []byte("secret")
)

func TestMain(m *testing.M) {
	server = NewServer()
	server.RegisterRoute()
	os.Exit(m.Run())
}

func TestHmac(t *testing.T) {
	t.Run("Success Test validate Hmac get authoriztion", func(t *testing.T) {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		data := hmac.GetData(http.MethodGet, "/api/test", timestamp)
		hash := hmac.MakeMac(data, __secret)

		req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
		req.Header.Set(hmac.HeaderHash, hash)
		req.Header.Set(hmac.HeaderTimestamp, timestamp)
		rec := httptest.NewRecorder()

		server.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, 200)
	})
}
