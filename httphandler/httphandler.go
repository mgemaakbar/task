package httphandler

import (
	"encoding/json"
	"net/http"
	"task/usecase"
	"task/websocketer"

	"github.com/gorilla/websocket"
)

type httphandler struct {
	uc     usecase.Usecase
	upgrdr *websocket.Upgrader
}

func NewHTTPHandler(uc usecase.Usecase, upgrdr *websocket.Upgrader) *httphandler {
	return &httphandler{uc: uc, upgrdr: upgrdr}
}

func (h *httphandler) ReadMessageRealTime(w http.ResponseWriter, r *http.Request) {
	readCon, err := h.upgrdr.Upgrade(w, r, nil)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		resp := ReadMessageRealTime{Response: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}
	h.uc.ReadMessageRealTime(websocketer.NewWebsocketer(readCon))
}

func (h *httphandler) ListSentMessage(w http.ResponseWriter, r *http.Request) {
	msgs, err := h.uc.ListSentMessage()
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		resp := ListSentMessagesResponse{Response: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := ListSentMessagesResponse{Response: "success", ListOfSentMessages: msgs}
	json.NewEncoder(w).Encode(resp)
	return
}

func (h *httphandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	msg, ok := r.URL.Query()["msg"]
	if !ok || len(msg[0]) < 1 {
		w.Header().Set("Content-Type", "application/json")
		resp := SendMessageResponse{Message: "Url Param 'msg' is missing"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	err := h.uc.WriteMessage(websocketer.NewWebsocketer(nil), msg[0])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		resp := SendMessageResponse{Message: err.Error()}
		json.NewEncoder(w).Encode(resp)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	resp := SendMessageResponse{Message: "message successfully sent!"}
	json.NewEncoder(w).Encode(resp)
}
