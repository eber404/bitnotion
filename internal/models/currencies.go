package models

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
