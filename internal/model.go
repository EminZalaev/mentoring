package internal

type currency struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
	Value        string `json:"value"`
}
