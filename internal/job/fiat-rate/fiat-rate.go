package fiat_rate

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"

	Config "github.com/E7ast1c/Cupbearer/config"
	"github.com/E7ast1c/Cupbearer/internal/provider/telegram"
	"github.com/sirupsen/logrus"
	"golang.org/x/text/encoding/charmap"
)

var requiredCharCode = map[string]any{
	"EUR": nil, "USD": nil,
}

type CbrRate struct {
	TgConfig Config.TgConfig
}

func (c CbrRate) Do() {
	rate, err := GetCBRRate()
	if err != nil {
		logrus.Error(err)
	}

	var resultValutes = make([]Valute, 0, len(requiredCharCode))
	for _, v := range rate.Valutes {
		if _, ok := requiredCharCode[v.CharCode]; ok {
			resultValutes = append(resultValutes, v)
		}
	}

	telegram.NewTgBot(c.TgConfig.Token).
		Send(telegram.Message{
			Text:                fmt.Sprint(resultValutes),
			ChatId:              c.TgConfig.ChatId,
			DisableNotification: true,
		})
}

func (c CbrRate) Cancel() {
	fmt.Println("Cancel CBR rate job")
}

func GetCBRRate() (*ValCurs, error) {
	resp, err := http.Get("https://cbr.ru/scripts/XML_daily.asp")
	if err != nil {
		return nil, err
	}

	vc := ValCurs{}
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = makeCharsetReader
	err = decoder.Decode(&vc)
	if err != nil {
		return nil, err
	}

	return &vc, nil
}

func makeCharsetReader(charset string, input io.Reader) (io.Reader, error) {
	if charset == "windows-1251" {
		return charmap.Windows1251.NewDecoder().Reader(input), nil
	}
	return nil, fmt.Errorf("Unknown charset: %s", charset)
}
