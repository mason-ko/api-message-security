package main

import (
	"api-message-security/src/hmac"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	t.Run("Success Test validate Hmac GET authorization", func(t *testing.T) {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		data := hmac.GetData(http.MethodGet, "/api/test?id=12345", timestamp)
		hash := hmac.MakeMac(data, __secret)

		fmt.Println(hash)

		req := httptest.NewRequest(http.MethodGet, "/api/test?id=12345", nil)
		req.Header.Set(hmac.HeaderHash, hash)
		req.Header.Set(hmac.HeaderTimestamp, timestamp)
		rec := httptest.NewRecorder()

		server.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, 200)
	})

	t.Run("Failure Test validate Hmac GET authorization", func(t *testing.T) {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		data := hmac.GetData(http.MethodGet, "/api/testXXXXXX", timestamp)
		hash := hmac.MakeMac(data, __secret)

		req := httptest.NewRequest(http.MethodGet, "/api/test", nil)
		req.Header.Set(hmac.HeaderHash, hash)
		req.Header.Set(hmac.HeaderTimestamp, timestamp)
		rec := httptest.NewRecorder()

		server.ServeHTTP(rec, req)

		assert.NotEqual(t, rec.Code, 200)
	})

	t.Run("Success Test validate Hmac POST authorization", func(t *testing.T) {
		timestamp := time.Now().UTC().Format(time.RFC3339)
		obj := map[string]string{
			"name": "mason",
			"a":    "12345",
			"b":    "12345",
			"c":    "12345",
			"d":    "12345",
		}
		bytes, _ := json.Marshal(obj)

		data := hmac.GetData(http.MethodPost, "/api/test", timestamp, string(bytes))
		hash := hmac.MakeMac(data, __secret)

		req := httptest.NewRequest(http.MethodPost, "/api/test", strings.NewReader(string(bytes)))
		req.Header.Set(hmac.HeaderHash, hash)
		req.Header.Set(hmac.HeaderTimestamp, timestamp)
		rec := httptest.NewRecorder()

		server.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, 200)
	})
}
