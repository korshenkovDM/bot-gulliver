package api

import (
	"log"

	handler "bot-gulliver/api/bot/handler"
	models "bot-gulliver/api/bot/models"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Bot(token string) {
	bot, err := tg.NewBotAPI(token)

	if err != nil {
		log.Println("func Bot 10 error")
		panic(err)
	}

	bot.Debug = true
	updateConf := tg.NewUpdate(0)
	updateConf.Timeout = 60
	updates := bot.GetUpdatesChan(updateConf)
	//
	//
	//
	// каналы на вход
	InMessage := make(chan *tg.Message, 9)
	InCallback := make(chan *tg.CallbackQuery, 9)
	// каналы на выход
	OutMessage := make(chan tg.MessageConfig, 9)
	OutCallback := make(chan tg.EditMessageTextConfig, 9)

	// горутины для взаимодействия
	go GoMessage(InMessage, OutMessage)
	go GoCallback(InCallback, OutCallback)
	go GoSender(bot, OutMessage, OutCallback)

	log.Println("BOT START!!!")

	for update := range updates {

		if update.Message != nil {
			// сделай что-нибудь с message
			InMessage <- update.Message
		}

		if update.CallbackQuery != nil {
			// сделай что-нибудь с callback
			bot.Send(tg.NewCallback(update.CallbackQuery.ID, ""))
			InCallback <- update.CallbackQuery
		}
	}
}

func GoMessage(in <-chan *tg.Message, out chan tg.MessageConfig) {
	for message := range in {
		req, ok := models.RequestMap[message.From.ID]

		if !ok {
			out <- tg.NewMessage(message.Chat.ID, "А мы знакомы?")
			continue

		}

		req.Mu.Lock()
		req.SetMsg(message)
		outMessage := handler.Message(req)
		req.Mu.Unlock()

		//
		for i := range outMessage {
			out <- outMessage[i]
		}
	}
}

func GoCallback(in <-chan *tg.CallbackQuery, out chan tg.EditMessageTextConfig) {
	for callback := range in {
		// сделай что-нибудь
		// обрабатывает request по ID, который либо закрывается, либо идет дальше
		req := models.RequestMap[callback.From.ID]

		req.Mu.Lock()
		req.Callback = callback
		outMessage := handler.Callback(req)
		req.Mu.Unlock()
		//
		out <- outMessage
	}
}

func GoSender(bot *tg.BotAPI, message <-chan tg.MessageConfig, callback <-chan tg.EditMessageTextConfig) {

	for {
		select {
		case msg := <-message:
			if _, err := bot.Send(msg); err != nil {
				log.Println(err, err.Error())
			}
		case msg := <-callback:
			if _, err := bot.Send(msg); err != nil {
				log.Println(err, err.Error())
			}
		}
	}
}
