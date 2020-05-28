package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func pushLock() {
	pushURL := "http://127.0.0.1:12345/push"
	contentType := "application/json"
	b := []byte(`{"message":"lock"}`)

	http.DefaultClient.Post(pushURL, contentType, bytes.NewReader(b))
}

func main() {
	server := NewServer(":12345")

	// Define websocket connect url, default "/ws"
	server.WSPath = "/ws"
	// Define push message url, default "/push"
	server.PushPath = "/push"

	// Set AuthToken func to authorize websocket connection, token is sent by
	// client for registe.
	server.AuthToken = func(token string) (userID string, ok bool) {
		// TODO: check if token is valid and calculate userID
		if token == "aaa" {
			return "jack", true
		}

		return "", false
	}

	// Set PushAuth func to check push request. If the request is valid, returns
	// true. Otherwise return false and request will be ignored.
	server.PushAuth = func(r *http.Request) bool {
		return true
	}

	http.HandleFunc("/monitor", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("[INFO]: /monitor GET")
		pushLock()
	})

	// Run server
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
