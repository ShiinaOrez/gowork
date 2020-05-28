package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

// websocketHandler defines to handle websocket upgrade request.
type websocketHandler struct {
	// upgrader is used to upgrade request.
	upgrader *websocket.Upgrader

	// calcUserIDFunc defines to calculate userID by token. The userID will
	// be equal to token if this function is nil.
	calcUserIDFunc func(token string) (userID string, ok bool)
}

// RegisterMessage defines message struct client send after connect
// to the server.
type RegisterMessage struct {
	Token string
	Event string
}

// First try to upgrade connection to websocket. If success, connection will
// be kept until client send close message or server drop them.
func (wh *websocketHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	wsConn, err := wh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer wsConn.Close()

	// handle Websocket request
	conn := NewConn(wsConn)
	SetConn(conn)
	conn.AfterReadFunc = func(messageType int, r io.Reader) {
		var rm RegisterMessage
		decoder := json.NewDecoder(r)
		if err := decoder.Decode(&rm); err != nil {
			return
		}

		// calculate userID by token
		if wh.calcUserIDFunc != nil {
			_, ok := wh.calcUserIDFunc(rm.Token)
			if !ok {
				return
			}
		}
	}

	conn.Listen()
}

// closeConns unbind conns filtered by userID and event and close them.
// The userID can't be empty, but event can be empty. The event will be ignored
// if empty.
func (wh *websocketHandler) closeConns(userID, event string) (int, error) {
	conn := GetConn()
	if err := conn.Close(); err != nil {
		log.Printf("conn close fail: %v", err)
	}
	return 0, nil
}

// ErrRequestIllegal describes error when data of the request is unaccepted.
var ErrRequestIllegal = errors.New("request data illegal")

// pushHandler defines to handle push message request.
type pushHandler struct {
	// authFunc defines to authorize request. The request will proceed only
	// when it returns true.
	authFunc func(r *http.Request) bool
}

// Authorize if needed. Then decode the request and push message to each
// realted websocket connection.
func (s *pushHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// authorize
	if s.authFunc != nil {
		if ok := s.authFunc(r); !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	// read request
	var pm PushMessage
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&pm); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(ErrRequestIllegal.Error()))
		return
	}

	cnt, err := s.push(pm.Message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	result := strings.NewReader(fmt.Sprintf("message sent to %d clients", cnt))
	io.Copy(w, result)
}

func (s *pushHandler) push(message string) (int, error) {
	if message == "" {
		return 0, errors.New("parameters(userId, event, message) can't be empty")
	}
	conn := GetConn()
	_, _ = conn.Write([]byte(message))
	return 0, nil
}

// PushMessage defines message struct send by client to push to each connected
// websocket client.
type PushMessage struct {
	Message string `json:"message"`
}
