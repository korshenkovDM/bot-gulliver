package do

import (
	"bot-gulliver/models"
	"bot-gulliver/restroom"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"time"
)

func ActiveList() string {
	var day = models.D
	date := day.Date.Format("2006-01-02 ")
	wd := restroom.Weekday(day.Date.Weekday().String())
	if day.Date.Format("2006-01-02") == "0001-01-01" {
		date = time.Now().Format("2006-01-02 ")
		wd = restroom.Weekday(time.Now().Weekday().String())
	}

	str := bytes.Buffer{}
	str.WriteString(date)
	str.WriteString(wd)
	for i := range day.Hour {
		if day.Hour[i] != nil {
			for a := range day.Hour[i] {
				inA := day.Hour[i][a]
				str.WriteString(fmt.Sprintf("🖌 %s *%s*\n   		%s %s\n", inA.Date.Format("15:04"), inA.NameClient, inA.Service, inA.NameUser))
			}
		}
	}
	str.WriteRune('\n')
	return str.String()
}

func ContractsList() string {
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

func SearchContract(first_name string) (names []string) {
	ln := len(first_name)

	contracts := models.ListContract()
	for key := range contracts {
		if len(key) > ln && key[:ln] == first_name {
			names = append(names, key)
		}
	}
	return names
}

func Act(user, client string, service string, price int, t time.Time) {
	models.CreateAct().
		SetupName(user).
		SetupDate(t).
		SetupData(client, service, price).
		SaveAct()
}

// CreateContract().SetupName(client).SetupDate(t).SetupStatus(status).SaveContract()
func Contract(client string, status string, t time.Time) {
	models.CreateContract().SetupName(client).SetupDate(t).SetupStatus(status).SaveContract()
}
