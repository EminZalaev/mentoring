package internal

type currency struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
	Value        int    `json:"value"`
	Well         float64
}

type currencyGetResponse struct {
	CurrencyFrom string `json:"currencyFrom"`
	CurrencyTo   string `json:"currencyTo"`
}

type responseStatus struct {
	Status string `json:"status"`
}
