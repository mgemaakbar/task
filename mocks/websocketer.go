package mocks

import (
	"errors"
	"task/websocketer"
)

type WebsocketerMock_ReadMessageError struct {
	websocketer.Websocketer
}

func (werr *WebsocketerMock_ReadMessageError) ReadMessage() (string, error) {
	return "msg", errors.New(`error`)
}

type WebsocketerMock_WriteMessageError struct {
	websocketer.Websocketer
}

func (werr *WebsocketerMock_WriteMessageError) WriteMessage(message string) error {
	return errors.New(`error`)
}


type WebsocketerMock_WriteMessageNoError struct {
	websocketer.Websocketer
}

func (werr *WebsocketerMock_WriteMessageNoError) WriteMessage(message string) error {
	return nil
}



