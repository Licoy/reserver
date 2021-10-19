package pkg

import (
	"github.com/gorilla/websocket"
)

const (
	RELOAD = "reload"
	CSS    = "css"
)

func WriteMsg(data string, conn *websocket.Conn) error {
	return conn.WriteMessage(websocket.TextMessage, []byte(data))
}

func WriteMsgToManyConn(data string, conn []*websocket.Conn) {
	msg := []byte(data)
	for _, w := range conn {
		_ = w.WriteMessage(websocket.TextMessage, msg)
	}
}
