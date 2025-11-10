package services

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/eber404/bitcoin-notion/internal/models"
	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

type NotionService struct {
	notion notionapi.Client
}

func NewNotionService() *NotionService {
	godotenv.Load()
	notionApiKey := os.Getenv("NOTION_API_KEY")

	if len(notionApiKey) == 0 {
		panic("Missing NOTION_API_KEY env var")
	}

	return &NotionService{
		notion: *notionapi.NewClient(notionapi.Token(notionApiKey)),
	}
}

const pageId = "2a35403ae33981cfa7f4eb25bd713204"

func (n *NotionService) UpdateBitcoinPricePage(bitcoinPrice models.BitcoinPrice) error {
	now := notionapi.Date(time.Now().UTC())
	_, err := n.notion.Page.Update(context.TODO(), notionapi.PageID(pageId), &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			string(models.ARS): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.ARS)},
			string(models.AUD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.AUD)},
			string(models.BRL): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.BRL)},
			string(models.CAD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.CAD)},
			string(models.CHF): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.CHF)},
			string(models.CLP): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.CLP)},
			string(models.CNY): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.CNY)},
			string(models.CZK): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.CZK)},
			string(models.DKK): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.DKK)},
			string(models.EUR): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.EUR)},
			string(models.GBP): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.GBP)},
			string(models.HKD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.HKD)},
			string(models.HRK): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.HRK)},
			string(models.HUF): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.HUF)},
			string(models.INR): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.INR)},
			string(models.ISK): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.ISK)},
			string(models.JPY): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.JPY)},
			string(models.KRW): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.KRW)},
			string(models.NGN): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.NGN)},
			string(models.NZD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.NZD)},
			string(models.PLN): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.PLN)},
			string(models.RON): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.RON)},
			string(models.RUB): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.RUB)},
			string(models.SEK): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.SEK)},
			string(models.SGD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.SGD)},
			string(models.THB): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.THB)},
			string(models.TRY): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.TRY)},
			string(models.TWD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.TWD)},
			string(models.USD): notionapi.NumberProperty{Number: bitcoinPrice.GetPrice(models.USD)},
			"updated_at":       notionapi.DateProperty{Date: &notionapi.DateObject{Start: &now}},
		},
	})

	if err != nil {
		return fmt.Errorf("error updating Notion page: %w", err)
	}

	return nil
}
