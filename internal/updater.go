package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	currencyApiHost   = "currate.ru"
	currencyApiScheme = "https"
	currencyApiPath   = "api/"
)

func UpdateCurrency(apiKey string, s Store) error {
	curr, err := s.GetCurrency()
	if err != nil {
		return fmt.Errorf("error get currencyRequest from storage: %w", err)
	}

	uri := buildReqUrl(apiKey, curr)

	resp, err := http.Get(uri)
	if err != nil {
		return fmt.Errorf("error get request from currencyRequest api: %w", err)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error read response body from currencyRequest api: %w", err)
	}

	defer func() {
		if ferr := resp.Body.Close(); err != nil && err == ferr {
			log.Println("error close resp body from currencyRequest api: %w", err)
		}
	}()

	strResp := string(respBody)

	for _, el := range *curr {
		pair := make([]byte, 0)
		pair = append(pair, el.CurrencyFrom...)
		pair = append(pair, el.CurrencyTo...)

		index := strings.Index(string(respBody), string(pair))
		cutCurr := strings.Split(strings.Split(strResp[index-1:], ",")[0], "\"")
		if len(cutCurr) < 2 {
			return fmt.Errorf("error value from currencyRequest api")
		}
		currStr := cutCurr[3]
		currFloat, err := strconv.ParseFloat(currStr, 64)
		if err != nil {
			return fmt.Errorf("error parse float value from currencyRequest api: %w", err)
		}

		if err := s.PutCurrency(&currency{
			CurrencyFrom: el.CurrencyFrom,
			CurrencyTo:   el.CurrencyTo,
			Well:         currFloat,
		}); err != nil {
			return fmt.Errorf("error put currencyRequest to storage: %w", err)
		}
	}
	return nil
}

func buildReqUrl(apiKey string, curr *[]currencyGetResponse) string {
	uri := url.URL{}

	currPairs := make([]byte, 0)
	for i, el := range *curr {
		currPairs = append(currPairs, el.CurrencyFrom...)
		currPairs = append(currPairs, el.CurrencyTo...)

		if i != len(*curr)-1 {
			currPairs = append(currPairs, ',')
		}
	}

	uri.Path = currencyApiPath
	uri.Scheme = currencyApiScheme
	uri.Host = currencyApiHost

	a := uri.Query()
	a.Set("get", "rates")
	a.Set("pairs", string(currPairs))
	a.Set("key", apiKey)

	uri.RawQuery = a.Encode()

	return uri.String()
}
