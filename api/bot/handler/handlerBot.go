package handler

import (
	//command "bot-gulliver/api/bot/command"
	command "bot-gulliver/api/bot/command"
	models "bot-gulliver/api/bot/models"

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
	}

	if firstName := commandWord; command.If_Len_1(commandWord) {
		// если это не команда, то это фамилия человека из ListContract
		if pickName, ok := command.SearchName(firstName, req.ChatID); ok {
			return pickName
		}

		var msg tg.MessageConfig
		// если не нашел такую фамилию среди действующих абонементов
		msg.ChatID = req.ChatID
		msg.Text = "не нашел таких хомяков\n или не понял команды"
		msg.ParseMode = "Markdown"
		return []tg.MessageConfig{msg}
	}

	// если это не команда и не фамилия человека, то это новый абонемент(len == 2)
	// сразу открываем панель оплаты
	var msg tg.MessageConfig
	msg.ChatID = req.ChatID
	msg.Text, msg.ReplyMarkup, msg.ParseMode = req.RequestText, models.MoneyMenuCredit(), "Markdown"
	req.Status = "month"

	return []tg.MessageConfig{msg}
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
