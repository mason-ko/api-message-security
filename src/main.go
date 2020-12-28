package main

// base64(hmac( method | url | timestamp ))

/*
curl -X GET "http://localhost:8080/api/test" \
-H "x_authorization: 91kKQixU8OZU1XcIBSSwo+G7l43M603De2LtjF/Khoo=" \
-H "x_timestamp: 2020-12-23T08:28:26.737Z"
*/

func main() {
	server := NewServer()
	server.RegisterRoute()
	server.Start()
}
