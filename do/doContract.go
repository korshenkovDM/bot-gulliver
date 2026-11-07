package do

import (
	"bot-gulliver/models"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func Contract(client string, status string, t time.Time) {
	// если контракт действующий
	if contract, ok := ListContract()[client]; ok {
		contract.DateEnd.AddDate(0, 0, 30)
		contract.SaveContract()
		return
	}

	// если контракт новый
	models.CreateContract().
		SetupName(client).
		SetupDate(t).
		SetupStatus(status).
		SaveContract()
	return
}

func ContractList() string {
	data, _ := os.ReadFile("contract.json")
	mp := make(map[string]models.Contract)
	json.Unmarshal(data, &mp)

	keys := make([]string, 0, len(mp))
	for key := range mp {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	str := bytes.Buffer{}
	for _, name := range keys {
		str.WriteString(fmt.Sprintf("%s %s\n(закончится)*%s*\n", mp[name].Status, mp[name].Name, mp[name].DateEnd.Format(time.DateOnly)))
	}
	return str.String()
}

func ListContract() map[string]models.Contract {
	var mp = make(map[string]models.Contract)
	data, err := os.ReadFile("contract.json")
	if err != nil {
		log.Println(err)
		return nil
	}
	json.Unmarshal(data, &mp)
	return mp
}

func ListContractEnd() string {
	mp := ListContract()
	names := bytes.Buffer{}
	for key, contract := range mp {
		if contract.EndContract() {
			// если это не должник и не на паузе, то его абонемент закончился
			if contract.Status == models.StatusActive {
				delete(mp, key)
				names.WriteString(key)
				names.WriteRune('\n')
			}
		}
	}
	ListSaveContract(mp)
	return names.String()
}

func SearchContract(name string) []string {
	ln, names := len(name), make([]string, 0, 4)
	contracts := ListContract()
	for key := range contracts {
		if ln < len(key) && key[:ln] == name {
			names = append(names, key)
		}
	}
	return names
}

func SearchContract_II(name_second string) []string {
	ln, names := len(name_second), make([]string, 0, 4)

	contract := ListContract()

	for key := range contract {
		tmp := len(key) - ln
		if tmp > 0 && key[tmp:] == name_second {
			names = append(names, key)
		}
	}

	return names
}

func ListSaveContract(contracts map[string]models.Contract) {
	data, err := json.MarshalIndent(contracts, "	", "	")
	if err != nil {
		log.Println(err)

	}
	err = os.WriteFile("contract.json", data, 0666)
	if err != nil {
		log.Println(err)

	}
}
