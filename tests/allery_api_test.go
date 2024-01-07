package tests

import (
	"log"
	"testing"
	"u-go-03/allergy_api"

	"github.com/joho/godotenv"
)

func TestAlleryApi(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	message, err := allergy_api.GetHourlyLoadData()

	if err != nil {
		t.Errorf("Error getting hourly load data: %s", err)
	}

	if message == nil {
		t.Errorf("Error getting hourly load data: message is null")
	}

	if *message == "" {
		t.Errorf("Error getting hourly load data: message is empty")
	}

	
}