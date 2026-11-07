package do

import (
	"bot-gulliver/day"
	"bot-gulliver/models"
	"bot-gulliver/restroom"
	"bytes"
	"fmt"
	"time"
)

func Act(user, client string, service string, price int, t time.Time) {
	act := models.CreateAct().
		SetupName(user).
		SetupDate(t).
		SetupData(client, service, price).
		ReturnAct()
	day.AddAct(act)
}

func ActiveList() string {
	var day = day.D
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

func ListAct()
