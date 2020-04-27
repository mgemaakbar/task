package websocketer

import (
	"github.com/gorilla/websocket"
)

type Websocketer interface {
	WriteMessage(message string) error
	ReadMessage() (string, error)
	Close() error
}

type websocketer struct {
	readCon  *websocket.Conn
	writeCon *websocket.Conn
}

func NewWebsocketer(readCon *websocket.Conn) *websocketer {
	return &websocketer{readCon: readCon}
}

func (w *websocketer) WriteMessage(message string) error {
	if w.writeCon == nil {
		var err error
		w.writeCon, _, err = websocket.DefaultDialer.Dial("ws://localhost:8080/read", nil)
		if err != nil {
			return err
		}
	}

	return w.writeCon.WriteMessage(websocket.TextMessage, []byte(message))
}

func (w *websocketer) ReadMessage() (string, error) {
	_, msg, err := w.readCon.ReadMessage()
	return string(msg), err
}

func (w *websocketer) Close() error {
	return w.readCon.Close()
}
