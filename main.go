package main

import (
	"log"
	"net/http"
)

var client = CreateXmppClient()
var config = GetConfig()

func main() {
	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
