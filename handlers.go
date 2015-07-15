package main

import (
	"encoding/json"
	"github.com/mattn/go-xmpp"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var message Message

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1000))
	if err != nil {
		log.Fatal(err)
	}
	if err := r.Body.Close(); err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(body, &message); err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	client.Send(xmpp.Chat{
		Remote: message.To,
		Text:   message.Text,
		Type:   "chat",
	})

	encoder := json.NewEncoder(w)
	chat, err := client.Recv()
	if err != nil {
		log.Fatal(err)
	}

	switch v := chat.(type) {

	case xmpp.Chat:
		respone := ChatResponse{
			Remote: v.Remote,
			Text:   v.Text,
		}
		err = encoder.Encode(respone)

	case xmpp.Presence:
		respone := PresenceResponse{
			From: v.From,
			Show: v.Show,
		}
		err = encoder.Encode(respone)
	}

	if err != nil {
		log.Fatal(err)
	}

}
