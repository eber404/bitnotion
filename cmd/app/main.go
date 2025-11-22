package main

import (
	"log"
	"time"

	"github.com/eber404/bitcoin-notion/internal/services"
)

func main() {
	cryptoService := services.NewCryptoService()
	dollarService := services.NewDollarService()
	notionService := services.NewNotionService()

	updateBitcoinQuote := func() {
		log.Println("Fetching Bitcoin quote and updating Notion...")
		bitcoinQuote, err := cryptoService.GetBitcoinPrice()
		if err != nil {
			log.Printf("error on get bitcoin quote: %s", err)
			return
		}

		err = notionService.UpdateBitcoinQuotePage(*bitcoinQuote)
		if err != nil {
			log.Printf("error on update notion page: %s", err)
			return
		}
		log.Println("Notion page with Bitcoin quote updated successfully.")
	}

	updateDollarQuote := func() {
		log.Println("Fetching Dollar quote and updating Notion...")
		dollarQuote, err := dollarService.GetDollarPrice()
		if err != nil {
			log.Printf("error on get dollar quote: %s", err)
			return
		}

		err = notionService.UpdateDollarQuotePage(*dollarQuote)
		if err != nil {
			log.Printf("error on update notion page with dollar quote: %s", err)
			return
		}
		log.Println("Notion page with Dollar quote updated successfully.")
	}

	updateBitcoinQuote()
	updateDollarQuote()

	const REFRESH_TIME = 5 * time.Minute
	ticker := time.NewTicker(REFRESH_TIME)
	defer ticker.Stop()

	for range ticker.C {
		updateBitcoinQuote()
		updateDollarQuote()
	}
}
