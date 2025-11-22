package models

import (
	"maps"
	"strconv"
)

type DollarPrice struct {
	currencies ExchangeRate
}

func NewDollarPrice(rates ExchangeRate) DollarPrice {
	copied := make(ExchangeRate, len(rates))
	maps.Copy(copied, rates)
	return DollarPrice{currencies: copied}
}

func (u DollarPrice) GetPrice(currency Currency) float64 {
	price := u.currencies[currency]
	return float64(price) / multiplierFactor
}

func (u DollarPrice) GetAll() ExchangeRate {
	copied := make(ExchangeRate, len(u.currencies))
	maps.Copy(copied, u.currencies)
	return copied
}

func ToCents(amount float64) int64 {
	return int64(amount * multiplierFactor)
}

func (u DollarPrice) GetStringPrice(currency Currency) string {
	price := u.GetPrice(currency)
	stringPrice := strconv.FormatFloat(price, 'f', 2, 64)
	return stringPrice
}
