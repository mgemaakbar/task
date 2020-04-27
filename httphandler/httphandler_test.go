package httphandler_test

import (
	"net/http"
	"net/http/httptest"
	"task/httphandler"
	"task/mocks"
	"testing"
)

func TestListSentMessage_Error(t *testing.T) {
	req, err := http.NewRequest("GET", "/wow", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := httphandler.NewHTTPHandler(&mocks.UsecaseMock_ListSentMessageError{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListSentMessage)

	handler.ServeHTTP(rr, req)

	expected := `{"response":"error","list_of_sent_messages":null}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestListSentMessage_NoError(t *testing.T) {
	req, err := http.NewRequest("GET", "/wow", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := httphandler.NewHTTPHandler(&mocks.UsecaseMock_ListSentMessageNoError{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.ListSentMessage)

	handler.ServeHTTP(rr, req)

	expected := `{"response":"success","list_of_sent_messages":["wow","wow2"]}` + "\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSendMessage_WriteMessageError(t *testing.T) {
	req, err := http.NewRequest("GET", "/wow?msg=xxx", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := httphandler.NewHTTPHandler(&mocks.UsecaseMock_WriteMessageError{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.SendMessage)

	handler.ServeHTTP(rr, req)

	expected := `{"response":"error"}` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSendMessage_NoError(t *testing.T) {
	req, err := http.NewRequest("GET", "/wow?msg=kek", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := httphandler.NewHTTPHandler(&mocks.UsecaseMock_WriteMessageNoError{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.SendMessage)

	handler.ServeHTTP(rr, req)

	expected := `{"response":"message successfully sent!"}` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSendMessage_BadRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/wow", nil)
	if err != nil {
		t.Fatal(err)
	}

	h := httphandler.NewHTTPHandler(&mocks.UsecaseMock_WriteMessageNoError{}, nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(h.SendMessage)

	handler.ServeHTTP(rr, req)

	expected := `{"response":"Url Param 'msg' is missing"}` + "\n"

	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
