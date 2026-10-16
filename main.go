package main

import (
	"log"
	"os"
	"time"
)

var API_TOKEN string

func main() {
	Start()

	log.Println("end")
}

func Start() {

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
	if Secret() {
		log.Println("Start SUCCESS")
	}
}

func Secret() bool {
	secret := os.Getenv("API_TOKEN")

	if secret == "" {
		log.Println("secret == nil \n API_TOKEN not found")
		return false
	}

	API_TOKEN = secret
	return true
}
