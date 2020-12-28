package main

// base64(hmac( method | url | timestamp ))

/*
curl -X GET "http://localhost:8080/api/test" \
-H "x_authorization: kOz2MR2AEASXyEvt0BaBTtCqxsw/bV8QkxS6jJiHuTA=" \
-H "x_timestamp: 2020-12-28T06:24:49Z"

curl -X POST "http://localhost:8080/api/test" \
-H "x_authorization: E36bN6iB3PB2y9S2QqDXfG2QX5wKFKK5HZ7PQPzY644=" \
-H "x_timestamp: 2020-12-28T06:29:21Z"


E36bN6iB3PB2y9S2QqDXfG2QX5wKFKK5HZ7PQPzY644= 2020-12-28T06:29:21Z POST|/api/test|2020-12-28T06:29:21Z|{"a":"12345","b":"12345","c":"12345","d":"12345","name":"mason"}
*/

func main() {
	server := NewServer()
	server.RegisterRoute()
	server.Start()
}
