package main

import (
	"flag"
	"log"
	"net/http"
	"task/httphandler"
	"task/usecase"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{}

func main() {
	flag.Parse()
	log.SetFlags(0)
	uc := usecase.NewUsecase([]string{})
	hand := httphandler.NewHTTPHandler(uc, &upgrader)
	http.HandleFunc("/list", hand.ListSentMessage)
	http.HandleFunc("/read", hand.ReadMessageRealTime)
	http.HandleFunc("/send", hand.SendMessage)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
