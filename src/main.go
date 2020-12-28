package main

// base64(hmac( method | url | timestamp ))

/*
curl -X GET "http://localhost:8080/api/test" \
-H "x_authorization: kOz2MR2AEASXyEvt0BaBTtCqxsw/bV8QkxS6jJiHuTA=" \
-H "x_timestamp: 2020-12-28T06:24:49Z"
*/

func main() {
	server := NewServer()
	server.RegisterRoute()
	server.Start()
}
