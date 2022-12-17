package internal

import (
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
		return err
	}

	uri := buildReqUrl(apiKey, curr)

	resp, err := http.Get(uri)
	if err != nil {
		return err
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	defer func() {
		if ferr := resp.Body.Close(); err != nil && err == ferr {
			log.Println(err)
		}
	}()

	strResp := string(respBody)

	for _, el := range *curr {
		pair := make([]byte, 0)
		pair = append(pair, el.CurrencyFrom...)
		pair = append(pair, el.CurrencyTo...)

		index := strings.IndexAny(string(respBody), string(pair))

		currStr := strings.Split(strings.Split(strResp[index-1:], ",")[0], "\"")[3]

		currFloat, err := strconv.ParseFloat(currStr, 64)
		if err != nil {
			return err
		}

		if err := s.PostCurrency(&currency{
			CurrencyFrom: el.CurrencyFrom,
			CurrencyTo:   el.CurrencyTo,
			Value:        strconv.Itoa(int(currFloat)),
		}); err != nil {
			return err
		}
	}
	return nil
}

func buildReqUrl(apiKey string, curr *[]currency) string {
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
