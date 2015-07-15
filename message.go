package main

type Message struct {
	To   string `json:"to"`
	Text string `json:"text"`
}

type ChatResponse struct {
	Remote string `json:"remote"`
	Text   string `json:"text"`
}

type PresenceResponse struct {
	From string `json:"from"`
	Show string `json:"show"`
}
