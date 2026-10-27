package restroom

import (
	"fmt"
	"log"
	"os"
	"time"
)

func Start() string {
	// change work dir, time.Local and create API_TOKEN
	err := os.Chdir("/data/DataBase")
	if err != nil {
		log.Println(err)
		err = os.Mkdir("data", 0755)
		if err != nil {
			log.Println(err, "func Start 13 error")
			//return
		}
		err = os.Chdir("/data")
		if err != nil {
			log.Println(err, "func Start 18 error")
			//	return
		}
	}
	err = os.Chdir("/data/DataBase")
	if err != nil {
		log.Println(err)
		err = os.Mkdir("DataBase", 0755)
		if err != nil {
			log.Println(err)
			//return
		}
		err = os.Chdir("/data/DataBase")
		if err != nil {
			log.Println(err, "func Start 32 error")
			//	return
		}
	}

	zone, err := time.LoadLocation("Europe/Moscow")

	if err != nil {
		log.Println(err)
	}

	time.Local = zone
	return secret()

}

func secret() string {
	secret := os.Getenv("API_TOKEN")

	if secret == "" {
		log.Println("secret == nil \n API_TOKEN not found")
	}
	return secret
}

func Answer(status string, object string) string {
	return fmt.Sprintf("%s %s - success\n %s", status, object, kek())
}
func kek() string {
	return "ты знал, что пчелы мастера бжж?"
}

func Weekday(day string) string {
	var answer string
	switch day {
	case "Monday":
		answer = "Понедельник\n"
	case "Tuesday":
		answer = "Вторник\n"
	case "Wednesday":
		answer = "Среда\n"
	case "Thursday":
		answer = "Четверг\n"
	case "Friday":
		answer = "Пятница🥂🍻\n"
	case "Saturday":
		answer = "Суббота\n"
	case "Sunday":
		answer = "Воксресенье\n"
	}
	return answer
}

func SearchTmpFile() bool {
	filepath := "tmp.json"

	_, err := os.Stat(filepath)
	return err == nil
}
