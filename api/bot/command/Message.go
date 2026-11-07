package command

import (
	//	models "bot-gulliver/api/bot/models"
	models "bot-gulliver/api/bot/models"
	"bot-gulliver/day"
	"bot-gulliver/do"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// "РАЗОВЫЙ" "ВСЕ" "АКТИВ" "STAT07"
// func(chatID)
var Message = map[string]func(int64) []tg.MessageConfig{
	"РАЗОВЫЙ": cmdOne,
	"ВСЕ":     cmdAll,
	"АКТИВ":   cmdActive,
	// soft
	"ENDDAY": cmdEndDay,
	"ОТМЕНА": cmdRecover,
	//
	// "SAVEDAY": cmdSaveDay
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
	msg.Text = do.ContractList()
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

// soft
func cmdEndDay(id int64) []tg.MessageConfig {
	group := models.RequestMap[id].Group

	if group != 0 {
		return nil
	}
	// идет всем
	endNames := do.ListContractEnd()
	// идет мне и Варе
	actsList := do.ActiveList()
	// сохраняем день в базу
	if day.SaveDay() != nil {
		return nil
	}
	// создаем новый день
	day.NewDay()
	return helperSpam(actsList, endNames)
}

func cmdRecover(id int64) []tg.MessageConfig {

	return []tg.MessageConfig{}
}

// text1 идет только Варе и Диме
// text 2 идет всем
func helperSpam(text1, text2 string) []tg.MessageConfig {
	var msgs = make([]tg.MessageConfig, 0, 4)
	for _, user := range models.RequestMap {
		var msg tg.MessageConfig
		msg.ChatID = user.ChatID
		msg.ParseMode = "Markdown"
		msg.Text = ""
		if user.Group == 2 {
			continue
		} else if user.Group == 0 {
			msg.Text += text1
		}
		if user.Group == 1 || user.Group == 0 {
			msg.Text += text2
		}
		msgs = append(msgs, msg)
	}
	return msgs
}
