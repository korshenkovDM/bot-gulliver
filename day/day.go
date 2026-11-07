package day

import (
	"bot-gulliver/models"
	"encoding/json"
	"log"
	"os"
	"time"
)

var D Day

type Day struct {
	Date time.Time
	Hour [13][]models.Act
}

// устанавливает время t = t.AddDate(0,0,1)
func SetupDate(t time.Time) {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t.AddDate(0, 0, 1)
	D.Date = t
	return
}

func AddAct(act models.Act) {
	timeLine, t := act.Date, act.Date
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	if !D.Date.Equal(t) {
		D.Date = t
	}

	index := timeLine.Hour() - 10
	if index == 13 {
		index--
	}
	if D.Hour[index] == nil {
		D.Hour[index] = make([]models.Act, 0, 4)
	}
	D.Hour[index] = append(D.Hour[index], act)
}

func SaveDay() error {
	var d Day
	d = D
	t := time.Now()
	data, err := os.ReadFile("activity.json")
	if err != nil {
		log.Println(err)
		return err
	}

	mp := make(map[string][]Day)
	err = json.Unmarshal(data, &mp)
	if err != nil {
		log.Println(err)
		return err
	}
	if _, month := mp[t.Format("2006-01")]; !month {
		mp[t.Format("2006-01")] = make([]Day, 1, 2)
	}
	mp[t.Format("2006-01")] = append(mp[t.Format("2006-01")], d)

	data, err = json.MarshalIndent(mp, "", "	")
	if err != nil {
		log.Println(err)
		return err
	}

	err = os.WriteFile("activity.json", data, 0666)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func NewDay() {
	D = Day{}
	SetupDate(time.Now())
}
