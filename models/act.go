package models

import "time"

type Act struct {
	Date       time.Time
	NameClient string
	Service    string
	Price      int
	NameUser   string
}

// 🟢🟡🔴💰💳🧨🏋‍♂️❄️🔥
const (
	Visit    = "🏋‍♂️"
	PauseOn  = "❄️"
	PauseOff = "🔥"
	Cash     = "💰"
	Card     = "💳"
	Credit   = "🧨"
)

func CreateAct() *Act {
	return &Act{}
}

func (act *Act) SetupName(name string) *Act {
	act.NameUser = name
	return act
}

func (act *Act) SetupDate(t time.Time) *Act {
	act.Date = t
	return act
}

func (act *Act) SetupData(nameClient string, service string, price int) *Act {
	act.NameClient, act.Service, act.Price = nameClient, service, price
	return act
}

func (act *Act) SaveAct() {
	D.AddAct(*act)
}
