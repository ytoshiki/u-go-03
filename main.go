package main

import (
	"log"
	"os"
	"u-go-03/allergy_api"
	"u-go-03/utils"

	"github.com/joho/godotenv"
	"github.com/robfig/cron/v3"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cronJob := cron.New()

	cronJob.AddFunc(os.Getenv("CRON_SCHEDULE"), func() {
		dailyAverageMessage, err := allergy_api.GetHourlyLoadData()
		if err != nil {
			log.Println("Error")
			return
		}

		// historicalAverageMesage, err := allergy_api.GetCurrentChartData()
		// if err != nil {
		// 	log.Println("Error")
		// 	return
		// }

		// slackMessage := *dailyAverageMessage + "\n" + *historicalAverageMesage
		err = utils.SendSlackMessage(*dailyAverageMessage)
		if err != nil {
			log.Println("Error")
			return
		}

	})

	cronJob.Start()

	select {}
}

