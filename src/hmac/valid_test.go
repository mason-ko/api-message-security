package hmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"log"
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

func TestMakeMac(t *testing.T) {
	api := "/api/test"
	method := "GET"
	time := "2020-12-23T08:28:26.737Z"
	data := method + DataSeparator + api + DataSeparator + time

	ret := MakeMac(data, __secret)
	log.Println(ret)

	bytes,_ := base64.StdEncoding.DecodeString(ret)

	ok := ValidMAC([]byte(data), bytes, __secret)
	fmt.Println(ok)
}