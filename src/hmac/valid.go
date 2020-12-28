package hmac

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	HeaderHash      = "x_authorization"
	HeaderTimestamp = "x_timestamp"
	DataSeparator   = "|"
)

var (
	__secret = []byte("secret")
)

func ValidGinMiddleware(c *gin.Context) {
	r := c.Request
	hash := r.Header.Get(HeaderHash)
	timestamp := r.Header.Get(HeaderTimestamp)

	if hash == "" || timestamp == "" {
		c.AbortWithStatus(401)
		return
	}

	h, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	var data string
	if c.Request.Method == http.MethodGet {
		data = GetData(r.Method, r.URL.RequestURI(), timestamp)
	} else {
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.AbortWithStatusJSON(401, err)
			return
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body)) // Write body back

		data = GetData(r.Method, r.URL.RequestURI(), timestamp, string(body))
	}

	result := ValidMAC([]byte(data), h, __secret)
	if !result {
		c.AbortWithStatus(401)
		return
	}

	c.Next()
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func MakeMac(data string, key []byte) string {
	h := hmac.New(sha256.New, key)
	h.Write([]byte(data))
	bytes := h.Sum(nil)
	str := base64.StdEncoding.EncodeToString(bytes)
	return str
}

func GetData(params ...string) string {
	return strings.Join(params, DataSeparator)
}
