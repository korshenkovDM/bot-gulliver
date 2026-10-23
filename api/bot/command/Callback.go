package command

import (
	models "bot-gulliver/api/bot/models"
	"bot-gulliver/do"
	back "bot-gulliver/models"
	"bot-gulliver/restroom"
	"time"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// "visit", "pay", "pause"
// "cash", "card", "credit"
// "cancel"
// func(chatID, messID, requestText string)
var Callback = map[string]func(int64, int, string) tg.EditMessageTextConfig{
	"visit": cmdVisit,
	"pay":   cmdPay,
	"pause": cmdPause,

	"cash":   cmdCash,
	"card":   cmdCard,
	"credit": cmdCredit,

	"cancel": cmdCancel,
}

// main command
func cmdVisit(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	var price int
	var method = back.Visit
	msg.ChatID, msg.MessageID = chatID, messID
	_, user := models.RequestMap[chatID].Status, models.RequestMap[chatID].Name
	if contract, ok := back.ListContract()[name]; ok {
		if contract.Status == back.PauseOn {
			method = back.PauseOff
			contract.SetupStatus(back.PauseOff).SaveContract()
		}
	}
	do.Act(user, name, method, price, time.Now())
	msg.Text = restroom.Answer(method, name)
	models.RequestMap[chatID].Status = ""
	return msg
}

func cmdPay(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	msg.ChatID, msg.MessageID = chatID, messID
	if contract, ok := back.ListContract()[name]; ok &&
		contract.Status == back.Credit {
		// close credit
		*msg.ReplyMarkup = models.MoneyMenu()
		models.RequestMap[chatID].Status = "credit"

	}
	*msg.ReplyMarkup = models.MoneyMenuCredit()
	models.RequestMap[chatID].Status = "month"
	return msg
}

func cmdPause(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	var method = back.PauseOn
	var price int

	msg.ChatID, msg.MessageID = chatID, messID
	_, user := models.RequestMap[chatID].Status, models.RequestMap[chatID].Name
	if contract, ok := back.ListContract()[name]; ok {
		if contract.Status == back.StatusActive {
			contract.SetupStatus(back.StatusPause).SaveContract()
		} else if contract.Status == back.StatusPause {
			contract.SetupStatus(back.StatusActive).SaveContract()
			method = back.PauseOff
		}
		do.Act(user, name, method, price, time.Now())
	}
	msg.Text = restroom.Answer(method, name)
	models.RequestMap[chatID].Status = ""
	return msg
}

// pay command
func cmdCash(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	var method = back.Cash
	var price int
	msg.ChatID, msg.MessageID = chatID, messID
	status, user := models.RequestMap[chatID].Status, models.RequestMap[chatID].Name
	if status == "one" {
		price = 200
	}
	if status == "month" {
		price = 1200
		do.Contract(name, back.StatusActive, time.Now())
	}
	if status == "credit" {
		if contract, ok := back.ListContract()[name]; ok &&
			contract.Status == back.Credit {
			contract.SetupStatus(back.StatusActive).SaveContract()
		}
		price = 1200
		status = "close" + back.Credit

	}
	do.Act(user, name, status+method, price, time.Now())
	msg.Text = restroom.Answer(method+status, name)
	models.RequestMap[chatID].Status = ""
	return msg
}

func cmdCard(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	var method = back.Card
	var price int
	msg.ChatID, msg.MessageID = chatID, messID
	status, user := models.RequestMap[chatID].Status, models.RequestMap[chatID].Name
	if status == "one" {
		price = 200
	}
	if status == "month" {
		price = 1200
		do.Contract(name, back.StatusActive, time.Now())
	}
	if status == "credit" {
		price = 1200
		status = "close" + back.Credit
	}
	do.Act(user, name, status+method, price, time.Now())
	msg.Text = restroom.Answer(method, name)
	models.RequestMap[chatID].Status = ""
	return msg
}

func cmdCredit(chatID int64, messID int, name string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	var method = back.Credit
	var price int
	msg.ChatID, msg.MessageID = chatID, messID
	status, user := models.RequestMap[chatID].Status, models.RequestMap[chatID].Name

	do.Act(user, name, status+method, price, time.Now())
	do.Contract(name, method, time.Now())
	msg.Text = restroom.Answer(method, name)
	models.RequestMap[chatID].Status = ""
	return msg
}

// soft command
func cmdCancel(chatID int64, messID int, cancelActive string) tg.EditMessageTextConfig {
	var msg tg.EditMessageTextConfig
	msg.ChatID, msg.MessageID = chatID, messID
	msg.Text = cancelActive + "\nотменил"
	return msg
}
