package models

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

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

func (contract *Contract) SaveContract() {
	var mp = make(map[string]Contract)
	data, err := os.ReadFile("contract.json")
	if err != nil {
		log.Println(err)
		return
	}
	json.Unmarshal(data, &mp)
	if tmp, ok := mp[contract.Name]; ok {
		contract.DateBegin = tmp.DateBegin
	}
	mp[contract.Name] = *contract

	data, err = json.MarshalIndent(mp, "	", "	")
	if err != nil {
		log.Println(err)
		return
	}
	err = os.WriteFile("contract.json", data, 0666)
	if err != nil {
		log.Println(err)
		return
	}
}

func (contract *Contract) EndContract() bool {
	t := time.Now()
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	if contract.DateEnd.Equal(t) || t.After(contract.DateEnd) {
		return true
	}
	return false
}

func ListContract() map[string]Contract {
	var mp = make(map[string]Contract)
	data, err := os.ReadFile("contract.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(data, &mp)
	return mp
}

func ListSaveContract(contracts map[string]Contract) {
	data, err := json.MarshalIndent(contracts, "	", "	")
	if err != nil {
		log.Println(err)

	}
	err = os.WriteFile("contract.json", data, 0666)
	if err != nil {
		log.Println(err)

	}
}
