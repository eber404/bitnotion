package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/eber404/bitcoin-notion/internal/models"
)

type DollarService struct {
}

func NewDollarService() *DollarService {
	return &DollarService{}
}

func (d DollarService) GetDollarPrice() (*models.DollarPrice, error) {
	baseUrl := os.Getenv("DOLLAR_BASE_URL")
	dollarApiKey := os.Getenv("DOLLAR_API_KEY")

	if len(baseUrl) == 0 {
		panic("Missing DOLLAR_BASE_URL env var")
	}

	if len(dollarApiKey) == 0 {
		panic("Missing DOLLAR_API_KEY env var")
	}

	url := baseUrl + "/v1/latest" + "?apikey=" + dollarApiKey + "&currencies=EUR,USD,CAD,BRL"

	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error on fetch dollar price")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error on fetch dollar price: %s", err)
	}

	var dolarPriceResponse DollarPriceResponse

	err = json.Unmarshal(body, &dolarPriceResponse)

	if err != nil {
		return nil, fmt.Errorf("error on unmarshal dollar price response body: %s", err)
	}

	price := models.NewDollarPrice(models.ExchangeRate{
		models.BRL: models.ToCents(dolarPriceResponse.Data.BRL),
		models.CAD: models.ToCents(dolarPriceResponse.Data.CAD),
		models.EUR: models.ToCents(dolarPriceResponse.Data.EUR),
		models.USD: models.ToCents(dolarPriceResponse.Data.USD),
	})

	return &price, nil
}

type DollarPriceResponse struct {
	Data struct {
		BRL float64 `json:"brl"`
		CAD float64 `json:"cad"`
		EUR float64 `json:"eur"`
		USD float64 `json:"usd"`
	}
}
