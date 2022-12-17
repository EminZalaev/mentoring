package internal

import (
	"net/http"
)

const cbApi = "https://www.cbr-xml-daily.ru/daily_json.js"

func UpdateCurrency(store Store) error {
	resp, err := http.Get(cbApi)
	if err != nil {
		return err
	}
	currency, err := store.GetCurrency()
	if err != nil {
		return err
	}
}
