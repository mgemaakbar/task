package usecase_test

import (
	"task/mocks"
	"task/usecase"
	"testing"
)

func TestListMessage(t *testing.T) {
	uc := usecase.NewUsecase([]string{"asd", "wew"})

	ret, err := uc.ListSentMessage()

	if err != nil {
		t.Error("Error should be nil")
	}

	if len(ret) != 2 {
		t.Errorf("Expected 2 for the array length, got %d", len(ret))
	}

	if ret[0] != "asd" || ret[1] != "wew" {
		t.Error("Wrong array content")
	}

}

func TestReadMessage(t *testing.T) {
	arr := []string{}
	ws := &mocks.WebsocketerMock_ReadMessageError{}
	uc := usecase.NewUsecase(arr)

	err := uc.ReadMessageRealTime(ws)

	if err == nil {
		t.Error("Error should not be nil")
	}

	if err.Error() != `error` {
		t.Errorf("Error string should be 'error', got '%s'", err.Error())
	}
}

func TestWriteMessage_Error(t *testing.T) {
	arr := []string{}
	ws := &mocks.WebsocketerMock_WriteMessageError{}
	uc := usecase.NewUsecase(arr)

	err := uc.WriteMessage(ws, "msg")

	if err == nil {
		t.Error("Error should not be nil")
	}

	if err.Error() != `error` {
		t.Errorf("Error string should be 'error', got '%s'", err.Error())
	}
}

func TestWriteMessage_NoError(t *testing.T) {
	arr := []string{}
	ws := &mocks.WebsocketerMock_WriteMessageNoError{}
	uc := usecase.NewUsecase(arr)

	err := uc.WriteMessage(ws, "msg")

	if err != nil {
		t.Error("Error should be nil")
	}

}
