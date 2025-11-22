package models

import (
	"maps"
	"strconv"
)

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
