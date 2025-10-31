package main

import "time"

type RegisterPayload struct {
	UserId    string `json:"userId"`
	Email     string `json:"email"`
	Timestamp string `json:"timestamp"`
}

func main() {
	data := RegisterPayload{
		userId:    "123",
		email:     "user@email.com",
		timestamp: time.Now().String(),
	}
	var payloadSend = map[string]interface{}{
		"userId":    data.userId,
		"email":     data.email,
		"timestamp": data.timestamp,
	}

	ProduceKeyService("email", payloadSend)
}
