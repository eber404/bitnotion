package main

import (
	"log"
	"time"

	"github.com/eber404/bitcoin-notion/internal/services"
)

func main() {
	cryptoService := services.NewCryptoService()
	notionService := services.NewNotionService()

	updatePrice := func() {
		log.Println("Fetching Bitcoin price and updating Notion...")
		bitcoinPrice, err := cryptoService.GetBitcoinPrice()
		if err != nil {
			log.Printf("error on get bitcoin price: %s", err)
			return
		}

		err = notionService.UpdateBitcoinPricePage(*bitcoinPrice)
		if err != nil {
			log.Printf("error on update notion page: %s", err)
			return
		}
		log.Println("Notion page updated successfully.")
	}

	updatePrice()

	const REFRESH_TIME = 5
	ticker := time.NewTicker(REFRESH_TIME * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		updatePrice()
	}
}
