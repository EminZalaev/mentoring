package internal

type currencyRequest struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
	Value        int    `json:"value"`
}

type currencyGetResponse struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
}
