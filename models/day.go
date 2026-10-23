package models

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

var D Day

type Day struct {
	Date time.Time
	Hour [13][]Act
}

func CreateDay() *Day {
	return &Day{}
}

func (day *Day) SetupDate(t time.Time) Day {
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	t.AddDate(0, 0, 1)
	day.Date = t
	return *day
}

func (day *Day) AddAct(act Act) {
	timeLine, t := act.Date, act.Date
	t = time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	if !day.Date.Equal(t) {
		day.Date = t
	}

	index := timeLine.Hour() - 10
	if index == 13 {
		index--
	}
	if day.Hour[index] == nil {
		day.Hour[index] = make([]Act, 0)
	}
	day.Hour[index] = append(day.Hour[index], act)
}

func (day *Day) DelAct() {}

func (day *Day) SaveDay() error {
	var d Day
	d = *day
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
