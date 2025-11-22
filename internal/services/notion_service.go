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

const bitcoinPageId = "2a35403ae33981cfa7f4eb25bd713204"
const dollarPageId = "2b35403ae33980e0b40ed08f7fb87b3d"

func (n *NotionService) UpdateBitcoinQuotePage(bitcoinQuote models.BitcoinPrice) error {
	now := notionapi.Date(time.Now().UTC())
	_, err := n.notion.Page.Update(context.TODO(), notionapi.PageID(bitcoinPageId), &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"updated_at":       notionapi.DateProperty{Date: &notionapi.DateObject{Start: &now}},
			string(models.ARS): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.ARS)},
			string(models.AUD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.AUD)},
			string(models.BRL): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.BRL)},
			string(models.CAD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.CAD)},
			string(models.CHF): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.CHF)},
			string(models.CLP): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.CLP)},
			string(models.CNY): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.CNY)},
			string(models.CZK): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.CZK)},
			string(models.DKK): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.DKK)},
			string(models.EUR): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.EUR)},
			string(models.GBP): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.GBP)},
			string(models.HKD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.HKD)},
			string(models.HRK): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.HRK)},
			string(models.HUF): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.HUF)},
			string(models.INR): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.INR)},
			string(models.ISK): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.ISK)},
			string(models.JPY): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.JPY)},
			string(models.KRW): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.KRW)},
			string(models.NGN): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.NGN)},
			string(models.NZD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.NZD)},
			string(models.PLN): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.PLN)},
			string(models.RON): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.RON)},
			string(models.RUB): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.RUB)},
			string(models.SEK): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.SEK)},
			string(models.SGD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.SGD)},
			string(models.THB): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.THB)},
			string(models.TRY): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.TRY)},
			string(models.TWD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.TWD)},
			string(models.USD): notionapi.NumberProperty{Number: bitcoinQuote.GetPrice(models.USD)},
		},
	})

	if err != nil {
		return fmt.Errorf("error updating Notion page: %w", err)
	}

	return nil
}

func (n *NotionService) UpdateDollarQuotePage(dollarQuote models.DollarPrice) error {
	now := notionapi.Date(time.Now().UTC())
	_, err := n.notion.Page.Update(context.TODO(), notionapi.PageID(dollarPageId), &notionapi.PageUpdateRequest{
		Properties: notionapi.Properties{
			"updated_at":       notionapi.DateProperty{Date: &notionapi.DateObject{Start: &now}},
			string(models.BRL): notionapi.NumberProperty{Number: dollarQuote.GetPrice(models.BRL)},
			string(models.CAD): notionapi.NumberProperty{Number: dollarQuote.GetPrice(models.CAD)},
			string(models.EUR): notionapi.NumberProperty{Number: dollarQuote.GetPrice(models.EUR)},
			string(models.USD): notionapi.NumberProperty{Number: dollarQuote.GetPrice(models.USD)},
		},
	})

	if err != nil {
		return fmt.Errorf("error updating Notion page: %w", err)
	}

	return nil
}
