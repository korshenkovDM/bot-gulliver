package handler

import (
	command "bot-gulliver/api/bot/command"
	models "bot-gulliver/api/bot/models"
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Message(req *models.Request) []tg.MessageConfig {
	commandWord := req.RequestText
	if comm, ok := command.Message[commandWord]; ok {
		//ВСЕ
		//Актив
		//Разовый
		//EndDay
		return comm(req.ChatID)
	} else {
		// если это не команда, то это фамилия человека из ListContract
		firstName := commandWord
		log.Println("pickName input: ", firstName)
		if pickName, ok := command.SearchName(firstName, req.ChatID); ok {
			return pickName
		}
	}

	return []tg.MessageConfig{tg.NewMessage(req.ChatID, "запрос не получилось обработать")}
}

func Callback(req *models.Request) tg.EditMessageTextConfig {
	commandWord := req.Callback.Data
	if comm, ok := command.Callback[commandWord]; ok {
		//visit pay pause
		//cash card credit
		//cancel
		return comm(req.ChatID, req.Callback.Message.MessageID, req.RequestText)
	} else {
		// если это не команда, то это имя человека из listContract
		secondName := commandWord
		if pickName, ok := command.PickName(req.ChatID, req.Callback.Message.MessageID, secondName); ok {
			return pickName
		}
	}

	return tg.EditMessageTextConfig{}
}
