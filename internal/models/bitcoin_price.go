package models

import (
	"maps"
	"strconv"
)

const multiplierFactor = 100

type Currency string

const (
	ARS Currency = "ARS"
	AUD Currency = "AUD"
	BRL Currency = "BRL"
	CAD Currency = "CAD"
	CHF Currency = "CHF"
	CLP Currency = "CLP"
	CNY Currency = "CNY"
	CZK Currency = "CZK"
	DKK Currency = "DKK"
	EUR Currency = "EUR"
	GBP Currency = "GBP"
	HKD Currency = "HKD"
	HRK Currency = "HRK"
	HUF Currency = "HUF"
	INR Currency = "INR"
	ISK Currency = "ISK"
	JPY Currency = "JPY"
	KRW Currency = "KRW"
	NGN Currency = "NGN"
	NZD Currency = "NZD"
	PLN Currency = "PLN"
	RON Currency = "RON"
	RUB Currency = "RUB"
	SEK Currency = "SEK"
	SGD Currency = "SGD"
	THB Currency = "THB"
	TRY Currency = "TRY"
	TWD Currency = "TWD"
	USD Currency = "USD"
)

type ExchangeRate map[Currency]int64

type BitcoinPrice struct {
	currencies ExchangeRate
}

func NewBitcoinPrice(rates ExchangeRate) BitcoinPrice {
	copied := make(ExchangeRate, len(rates))
	maps.Copy(copied, rates)

	return BitcoinPrice{currencies: copied}
}

func (b BitcoinPrice) GetPrice(currency Currency) float64 {
	price := b.currencies[currency]

	return float64(price) / multiplierFactor
}

func (b BitcoinPrice) GetAll() ExchangeRate {
	copied := make(ExchangeRate, len(b.currencies))
	maps.Copy(copied, b.currencies)
	return copied
}

func ToSatoshis(amount float64) int64 {
	return int64(amount * multiplierFactor)
}

func ToBitcoin(amount int64) float64 {
	return float64(amount / multiplierFactor)
}

func (b BitcoinPrice) GetStringPrice(currency Currency) string {
	price := b.GetPrice(currency)

	stringPrice := strconv.FormatFloat(price, 'f', 2, 64)

	return stringPrice
}
