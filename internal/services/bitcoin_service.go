package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/eber404/bitcoin-notion/internal/models"
)

type CryptoService struct {
}

func NewCryptoService() *CryptoService {
	return &CryptoService{}
}

func (c CryptoService) GetBitcoinPrice() (*models.BitcoinPrice, error) {
	res, err := http.Get("https://blockchain.info/ticker")

	if err != nil {
		return nil, fmt.Errorf("error on fetch bitcoin price: %s", err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, fmt.Errorf("error on read response body: %s", err)

	}

	var bitcoinPriceDTO BitcoinPriceResponse
	err = json.Unmarshal(body, &bitcoinPriceDTO)

	if err != nil {
		return nil, fmt.Errorf("error on unmarshal response body: %s", err)
	}

	exchangeRate := models.ExchangeRate{
		models.ARS: models.ToSatoshis(bitcoinPriceDTO.ARS.Sell),
		models.AUD: models.ToSatoshis(bitcoinPriceDTO.AUD.Sell),
		models.BRL: models.ToSatoshis(bitcoinPriceDTO.BRL.Sell),
		models.CAD: models.ToSatoshis(bitcoinPriceDTO.CAD.Sell),
		models.CHF: models.ToSatoshis(bitcoinPriceDTO.CHF.Sell),
		models.CLP: models.ToSatoshis(bitcoinPriceDTO.CLP.Sell),
		models.CNY: models.ToSatoshis(bitcoinPriceDTO.CNY.Sell),
		models.CZK: models.ToSatoshis(bitcoinPriceDTO.CZK.Sell),
		models.DKK: models.ToSatoshis(bitcoinPriceDTO.DKK.Sell),
		models.EUR: models.ToSatoshis(bitcoinPriceDTO.EUR.Sell),
		models.GBP: models.ToSatoshis(bitcoinPriceDTO.GBP.Sell),
		models.HKD: models.ToSatoshis(bitcoinPriceDTO.HKD.Sell),
		models.HRK: models.ToSatoshis(bitcoinPriceDTO.HRK.Sell),
		models.HUF: models.ToSatoshis(bitcoinPriceDTO.HUF.Sell),
		models.INR: models.ToSatoshis(bitcoinPriceDTO.INR.Sell),
		models.ISK: models.ToSatoshis(bitcoinPriceDTO.ISK.Sell),
		models.JPY: models.ToSatoshis(bitcoinPriceDTO.JPY.Sell),
		models.KRW: models.ToSatoshis(bitcoinPriceDTO.KRW.Sell),
		models.NGN: models.ToSatoshis(bitcoinPriceDTO.NGN.Sell),
		models.NZD: models.ToSatoshis(bitcoinPriceDTO.NZD.Sell),
		models.PLN: models.ToSatoshis(bitcoinPriceDTO.PLN.Sell),
		models.RON: models.ToSatoshis(bitcoinPriceDTO.RON.Sell),
		models.RUB: models.ToSatoshis(bitcoinPriceDTO.RUB.Sell),
		models.SEK: models.ToSatoshis(bitcoinPriceDTO.SEK.Sell),
		models.SGD: models.ToSatoshis(bitcoinPriceDTO.SGD.Sell),
		models.THB: models.ToSatoshis(bitcoinPriceDTO.THB.Sell),
		models.TRY: models.ToSatoshis(bitcoinPriceDTO.TRY.Sell),
		models.TWD: models.ToSatoshis(bitcoinPriceDTO.TWD.Sell),
		models.USD: models.ToSatoshis(bitcoinPriceDTO.USD.Sell),
	}
	bitcoinPrice := models.NewBitcoinPrice(exchangeRate)

	return &bitcoinPrice, nil
}

type BitcoinPriceResponse struct {
	ARS struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ARS"`
	AUD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"AUD"`
	BRL struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"BRL"`
	CAD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CAD"`
	CHF struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CHF"`
	CLP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CLP"`
	CNY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CNY"`
	CZK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"CZK"`
	DKK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"DKK"`
	EUR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"EUR"`
	GBP struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"GBP"`
	HKD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HKD"`
	HRK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HRK"`
	HUF struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"HUF"`
	INR struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"INR"`
	ISK struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"ISK"`
	JPY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"JPY"`
	KRW struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"KRW"`
	NGN struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NGN"`
	NZD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"NZD"`
	PLN struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"PLN"`
	RON struct {
		One5M  float64 `json:"1m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RON"`
	RUB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"RUB"`
	SEK struct {
		One5M  float64 `json:"1m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SEK"`
	SGD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"SGD"`
	THB struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"THB"`
	TRY struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TRY"`
	TWD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"TWD"`
	USD struct {
		One5M  float64 `json:"15m"`
		Last   float64 `json:"last"`
		Buy    float64 `json:"buy"`
		Sell   float64 `json:"sell"`
		Symbol string  `json:"symbol"`
	} `json:"USD"`
}
