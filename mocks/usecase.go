package mocks

import (
	"errors"
	"task/usecase"
	"task/websocketer"
)

type UsecaseMock_ListSentMessageError struct {
	usecase.Usecase
}

func (u *UsecaseMock_ListSentMessageError) ListSentMessage() ([]string, error) {
	return nil, errors.New(`error`)
}

type UsecaseMock_ListSentMessageNoError struct {
	usecase.Usecase
}

func (u *UsecaseMock_ListSentMessageNoError) ListSentMessage() ([]string, error) {
	return []string{"wow", "wow2"}, nil
}

type UsecaseMock_WriteMessageError struct {
	usecase.Usecase
}

func (u *UsecaseMock_WriteMessageError) WriteMessage(ws websocketer.Websocketer, message string) error {
	return errors.New(`error`)
}

type UsecaseMock_WriteMessageNoError struct {
	usecase.Usecase
}

func (u *UsecaseMock_WriteMessageNoError) WriteMessage(ws websocketer.Websocketer, message string) error {
	return nil
}
