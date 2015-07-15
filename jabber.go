package main

import (
	"crypto/tls"
	"github.com/mattn/go-xmpp"
	"log"
)

func CreateXmppClient() *xmpp.Client {
	var client *xmpp.Client
	var err error

	xmpp.DefaultConfig = tls.Config{
		ServerName:         "jabber.ru",
		InsecureSkipVerify: false,
	}

	options := xmpp.Options{
		Host:     config.Host,
		User:     config.User,
		Password: config.Password,
		Session:  true,
		NoTLS:    false,
		StartTLS: true,
		Debug:    true,
	}

	client, err = options.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	return client
}
