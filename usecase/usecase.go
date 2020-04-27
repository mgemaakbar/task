package usecase

import (
	"fmt"
	"sync"
	"task/websocketer"
)

type Usecase interface {
	ReadMessageRealTime(ws websocketer.Websocketer) error
	WriteMessage(ws websocketer.Websocketer, message string) error
	ListSentMessage() ([]string, error)
}

type usecase struct {
	sentMessages []string
	ws           websocketer.Websocketer
}

func NewUsecase(sentMessages []string) *usecase {
	return &usecase{sentMessages: sentMessages}
}

func (u *usecase) ReadMessageRealTime(ws websocketer.Websocketer) error {
	var err error
	var message string
	for {
		message, err = ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("Read message: ", message)
	}
	return err
}

func (u *usecase) WriteMessage(ws websocketer.Websocketer, message string) error {
	err := ws.WriteMessage(message)
	if err != nil {
		return err
	}

	mu := &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()
	u.sentMessages = append(u.sentMessages, message)

	return nil
}

func (u *usecase) ListSentMessage() ([]string, error) {
	ret := []string{}

	mu := &sync.Mutex{}
	mu.Lock()
	defer mu.Unlock()

	for _, msg := range u.sentMessages {
		ret = append(ret, msg)
	}

	return ret, nil
}
