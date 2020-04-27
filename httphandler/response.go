package httphandler

type SendMessageResponse struct {
	Message string `json:"response"`
}

type ListSentMessagesResponse struct {
	Response           string   `json:"response"`
	ListOfSentMessages []string `json:"list_of_sent_messages"`
}

type ReadMessageRealTime struct {
	Response string `json:"response"`
}
