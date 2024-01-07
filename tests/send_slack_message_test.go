package tests

import (
	"log"
	"testing"
	"u-go-03/utils"

	"github.com/joho/godotenv"
)

func TestSlackMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	err = utils.SendSlackMessage("Test Message")

	if err != nil {
		t.Errorf("Error sending Slack message: %s", err)
	}
}