package models

type CoinDeskResponse struct {
	Time       TimeResponse                   `json:"time"`
	Disclaimer string                         `json:"disclaimer"`
	chartName  string                         `json:"chartName"`
	Bpi        map[string]cryptoPriceCurrency `json:"bpi"`
}

type TimeResponse struct {
	Updated    string `json:"updated"`
	UpdatedISO string `json:"updatedISO"`
	UpdatedUK  string `json:"updatedUK"`
}

type cryptoPriceCurrency struct {
	Code        string  `json:"code"`
	Symbol      string  `json:"symbol"`
	Rate        string  `json:"rate"`
	Description string  `json:"description"`
	RateFloat   float64 `json:"rateFloat"`
}

func (c CoinDeskResponse) GetCryptoName() string {
	return c.chartName
}

func (c CoinDeskResponse) GetCryptoBpi() map[string]cryptoPriceCurrency {
	return c.Bpi
}

func (response CoinDeskResponse) GetPrice(currencyIdentifier string) string {
	return response.Bpi[currencyIdentifier].Rate
}
