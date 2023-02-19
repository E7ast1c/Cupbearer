package telegram

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type TgBot struct {
	Token   string
	BotName string
}

func NewTgBot(token string) *TgBot {
	return &TgBot{Token: token}
}

type Message struct {
	Text                string `json:"text"`
	ChatId              string `json:"chat_id"`
	DisableNotification bool   `json:"disable_notification"`
}

func (t TgBot) Send(m Message) {
	requestUrl := "https://api.telegram.org/bot" + t.Token + "/sendMessage"

	client := &http.Client{}

	jsonParam, err := json.Marshal(m)
	if err != nil {
		logrus.Error(err)
		return
	}

	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonParam))
	if err != nil {
		logrus.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.Info(res.Status)
	defer res.Body.Close()
}
