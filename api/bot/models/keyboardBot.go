package api

import (
	"strings"

	tg "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// first level menu
func MainMenu() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("🏋‍♂️ visit", "visit"),
			tg.NewInlineKeyboardButtonData("💸 pay", "pay"),
			tg.NewInlineKeyboardButtonData("❄️ pause", "pause"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel")))
}

// cash one menu
func MoneyMenu() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("💳 card", "card"),
			tg.NewInlineKeyboardButtonData("💰 cash", "cash"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"),
		))
}

// cash month menu
func MoneyMenuCredit() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("💳 card", "card"),
			tg.NewInlineKeyboardButtonData("💰 cash", "cash"),
			tg.NewInlineKeyboardButtonData("🧨 credit", "credit"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"),
		))
}

func BookMenuStart() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("⬅️ back", "back"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"),
		))
}

func BookMenu() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("⬅️ back", "back"),
			tg.NewInlineKeyboardButtonData("➡️ next", "next"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"),
		))
}

func BookMenuNext() tg.InlineKeyboardMarkup {
	return tg.NewInlineKeyboardMarkup(
		tg.NewInlineKeyboardRow(
			tg.NewInlineKeyboardButtonData("➡️ next", "next"),
			tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"),
		))
}

// duple names
func NamesMenu(names []string) (k tg.InlineKeyboardMarkup) {
	var b []tg.InlineKeyboardButton

	for i := range names {
		name := strings.Split(names[i], " ")[1]
		b = append(b, tg.NewInlineKeyboardButtonData(name, name))
	}
	b = append(b, tg.NewInlineKeyboardButtonData("🚫 cancel", "cancel"))
	k = tg.NewInlineKeyboardMarkup(b)
	return k
}
