package models

import "time"

type Contract struct {
	Name      string
	DateBegin time.Time `json:"Begin"`
	DateEnd   time.Time `json:"End"`
	Status    string
}

const (
	StatusActive = "🟢"
	StatusPause  = "❄️"
	StatusCredit = "🧨"
)

func CreateContract() *Contract {
	return &Contract{}
}
func (contract *Contract) SetupName(name string) *Contract {
	contract.Name = name
	return contract
}
func (contract *Contract) SetupStatus(status string) *Contract {
	// closeCredit
	if contract.Status == StatusCredit && status == StatusActive {
		contract.Status = StatusActive
	}
	// PauseOff
	if contract.Status == StatusPause && status == StatusActive {
		t := time.Now()

		contract.Status = StatusActive
		diff := int(contract.DateBegin.Sub(contract.DateEnd).Hours() / 24)
		contract.DateBegin = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
		contract.DateEnd = contract.DateBegin.AddDate(0, 0, diff)
	}
	// PauseOn
	if contract.Status == StatusActive && status == StatusPause {
		t := time.Now()

		contract.Status = StatusPause
		contract.DateBegin = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	}
	contract.Status = status
	return contract
}
func (contract *Contract) SetupDate(t time.Time) *Contract {
	contract.DateBegin = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	contract.DateEnd = contract.DateBegin.AddDate(0, 0, 30)
	return contract
}

func (contract *Contract) SaveContract()

func ListContract() map[string]Contract {
	return make(map[string]Contract)
}
