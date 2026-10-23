package command

import (
	models "bot-gulliver/api/bot/models"
	"bot-gulliver/do"
	back "bot-gulliver/models"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// "РАЗОВЫЙ" "ВСЕ" "АКТИВ" "STAT07"
// func(chatID)
var Message = map[string]func(int64) []tg.MessageConfig{
	"РАЗОВЫЙ": cmdOne,
	"ВСЕ":     cmdAll,
	"АКТИВ":   cmdActive,
	"ENDDAY":  cmdEndDay,
}

func cmdOne(id int64) []tg.MessageConfig {
	var msg tg.MessageConfig
	msg.ChatID = id
	msg.Text = "*Выбери способ оплаты*"
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = models.MoneyMenu()
	models.RequestMap[id].Status = "one"
	models.RequestMap[id].RequestText = "РАЗОВЫЙ"
	return []tg.MessageConfig{msg}
}

func cmdAll(id int64) []tg.MessageConfig {
	var msg tg.MessageConfig
	msg.ChatID = id
	msg.Text = do.ContractsList()
	msg.ParseMode = "Markdown"
	return []tg.MessageConfig{msg}
}

func cmdActive(id int64) []tg.MessageConfig {
	var msg tg.MessageConfig
	msg.ChatID = id
	msg.Text = do.ActiveList()
	msg.ParseMode = "Markdown"
	return []tg.MessageConfig{msg}
}

func cmdEndDay(id int64) []tg.MessageConfig {
	if models.RequestMap[id].Group != 0 {
		return []tg.MessageConfig{}
	}

	var t time.Time
	var text string
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	err := back.D.SaveDay()
	if err != nil {
		text = err.Error()
	}
	back.D = back.CreateDay().SetupDate(t)

	var msg tg.MessageConfig

	msg.ChatID = id
	msg.Text = text
	msg.ParseMode = "Markdown"
	return []tg.MessageConfig{msg}
}
