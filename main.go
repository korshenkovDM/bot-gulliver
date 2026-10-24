package main

import (
	api "bot-gulliver/api/bot"
	"bot-gulliver/restroom"
)

// 🟢🟡🔴💰💳🧨🏋‍♂️❄️🔥
func main() {

	token := restroom.Start()
	api.Bot(token)
}
