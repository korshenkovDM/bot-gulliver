package api

import (
	"strings"
	"sync"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var RequestMap = map[int64]*Request{
	404531178:  {Name: "Davalka", ChatID: 404531178, Group: 0},  //dima
	920344106:  {Name: "VaraVar", ChatID: 920344106, Group: 0},  //varya
	5649164769: {Name: "Romulus", ChatID: 5649164769, Group: 1}, //rome
	1410865161: {Name: "Gendalf", ChatID: 1410865161, Group: 1}, //gena
	1:          {Name: "test", ChatID: 1, Group: 2},
}

type Request struct {
	Name        string
	Group       uint8
	ChatID      int64
	MessageID   int
	RequestText string
	Callback    *tg.CallbackQuery
	Status      string
	Mu          sync.Mutex
}

func (r *Request) SetMsg(msg *tg.Message) {
	r.ChatID = msg.Chat.ID
	r.MessageID = msg.MessageID
	r.RequestText = strings.ToUpper(msg.Text)
	r.Status = "main"
}
