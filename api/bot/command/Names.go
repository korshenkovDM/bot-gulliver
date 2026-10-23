package command

import (
	models "bot-gulliver/api/bot/models"
	"bot-gulliver/do"
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// if len(str) == 1(значит это поиск по фамилии)
func Names(name string) bool {
	return len(strings.Split(name, " ")) == 1
}

// поиск по first_name
func SearchName(name string, chatID int64) ([]tg.MessageConfig, bool) {
	names := do.SearchContract(name)
	if len(names) == 0 {
		// не нашел людей с таким first_name
		return nil, false
	}

	var msg tg.MessageConfig

	msg.ChatID = chatID
	msg.Text = "*нашел нескольких*"
	msg.ParseMode = "Markdown"
	msg.ReplyMarkup = models.NamesMenu(names)
	models.RequestMap[chatID].Status = "names"

	return []tg.MessageConfig{msg}, true
}

// выбрать second_name из списка
func PickName(chatID int64, msgID int, secondName string) (tg.EditMessageTextConfig, bool) {
	if models.RequestMap[chatID].Status != "names" {
		return tg.EditMessageTextConfig{}, false
	}
	firstName := models.RequestMap[chatID].RequestText
	fullName := firstName + " " + secondName
	models.RequestMap[chatID].RequestText = fullName
	models.RequestMap[chatID].Status = "main"

	var msg tg.EditMessageTextConfig

	msg.ChatID, msg.MessageID = chatID, msgID
	msg.Text = fullName
	msg.ParseMode = "Markdown"
	k := models.MainMenu()
	msg.ReplyMarkup = &k

	return msg, true
}
