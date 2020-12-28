package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
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

	data := r.Method + DataSeparator + r.URL.Path + DataSeparator + timestamp
	fmt.Println(hash, timestamp, data)

	h, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		c.AbortWithStatus(401)
		return
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
