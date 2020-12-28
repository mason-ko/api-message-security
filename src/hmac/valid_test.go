package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestValidMAC_Direct(t *testing.T) {
	data := "/user/info"

	fmt.Printf("Secret: %s Data: %s\n", __secret, data)

	h := hmac.New(sha256.New, __secret)
	h.Write([]byte(data))

	ret := ValidMAC([]byte(data), h.Sum(nil), __secret)
	fmt.Println(ret)
}

func TestMakeMacGet(t *testing.T) {
	method := "GET"
	api := "/api/test"
	time := "2020-12-23T08:28:26.737Z"

	data := GetData(method, api, time)

	ret := MakeMac(data, __secret)
	bytes, _ := base64.StdEncoding.DecodeString(ret)
	ok := ValidMAC([]byte(data), bytes, __secret)

	assert.Equal(t, ok, true)
}

func TestMakeMacPost(t *testing.T) {
	method := "POST"
	api := "/api/test"
	time := "2020-12-23T08:28:26.737Z"
	obj := map[string]string{
		"name": "mason",
		"a":    "12345",
		"b":    "12345",
		"c":    "12345",
		"d":    "12345",
	}
	b, _ := json.Marshal(obj)

	data := GetData(method, api, time, string(b))

	ret := MakeMac(data, __secret)
	bytes, _ := base64.StdEncoding.DecodeString(ret)
	ok := ValidMAC([]byte(data), bytes, __secret)

	assert.Equal(t, ok, true)
}
